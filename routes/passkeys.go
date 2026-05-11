package routes

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aarondl/sqlboiler/v4/queries/qm"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/j45k4/talktocow/config"
	"github.com/j45k4/talktocow/models"
)

const (
	passkeySessionKindLogin        = "login"
	passkeySessionKindRegistration = "registration"
	passkeySessionTTL              = 5 * time.Minute
)

var errPasskeyCredentialAlreadyExists = errors.New("passkey credential already exists")

type PasskeyBeginResponse struct {
	CeremonyID string      `json:"ceremonyId"`
	Options    interface{} `json:"options"`
}

type PasskeyFinishRequest struct {
	CeremonyID string          `json:"ceremonyId"`
	Response   json.RawMessage `json:"response"`
}

type PasskeyMetadata struct {
	ID                      int        `json:"id"`
	Name                    string     `json:"name"`
	AuthenticatorAttachment string     `json:"authenticatorAttachment"`
	CloneWarning            bool       `json:"cloneWarning"`
	CreatedAt               time.Time  `json:"createdAt"`
	LastUsedAt              *time.Time `json:"lastUsedAt"`
}

type webAuthnSessionRecord struct {
	CeremonyKind string
	UserID       sql.NullInt64
	RPID         string
	SessionData  webauthn.SessionData
	ExpiresAt    time.Time
}

type webAuthnUser struct {
	user        *models.User
	handle      []byte
	credentials []webauthn.Credential
}

func (u webAuthnUser) WebAuthnID() []byte {
	return u.handle
}

func (u webAuthnUser) WebAuthnName() string {
	if u.user.Username.Valid && u.user.Username.String != "" {
		return u.user.Username.String
	}

	return fmt.Sprint(u.user.ID)
}

func (u webAuthnUser) WebAuthnDisplayName() string {
	if u.user.Name.Valid && u.user.Name.String != "" {
		return u.user.Name.String
	}

	return u.WebAuthnName()
}

func (u webAuthnUser) WebAuthnCredentials() []webauthn.Credential {
	return u.credentials
}

func newWebAuthn() (*webauthn.WebAuthn, error) {
	return webauthn.New(&webauthn.Config{
		RPID:          config.WebAuthnRPID,
		RPDisplayName: config.WebAuthnRPDisplayName,
		RPOrigins:     config.WebAuthnRPOrigins,
	})
}

func HandlePasskeyLoginBegin(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	wa, err := newWebAuthn()

	if err != nil {
		log.Printf("creating webauthn relying party failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Passkey login is not configured"))
		return
	}

	assertion, session, err := wa.BeginDiscoverableLogin(
		webauthn.WithUserVerification(protocol.VerificationRequired),
	)

	if err != nil {
		log.Printf("beginning passkey login failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Could not start passkey login"))
		return
	}

	ceremonyID, err := saveWebAuthnSession(ctx.Request.Context(), db, passkeySessionKindLogin, config.WebAuthnRPID, nil, session)

	if err != nil {
		log.Printf("saving passkey login session failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Could not start passkey login"))
		return
	}

	ctx.JSON(http.StatusOK, PasskeyBeginResponse{
		CeremonyID: ceremonyID,
		Options:    assertion.Response,
	})
}

func HandlePasskeyLoginFinish(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	var body PasskeyFinishRequest

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Invalid passkey login response"))
		return
	}

	sessionRecord, err := loadAndDeleteWebAuthnSession(ctx.Request.Context(), db, body.CeremonyID)

	if err != nil {
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(InvalidCredentials, "Passkey login expired or is invalid"))
		return
	}

	if sessionRecord.CeremonyKind != passkeySessionKindLogin || sessionRecord.ExpiresAt.Before(time.Now()) {
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(InvalidCredentials, "Passkey login expired or is invalid"))
		return
	}

	parsedResponse, err := protocol.ParseCredentialRequestResponseBytes(body.Response)

	if err != nil {
		log.Printf("parsing passkey login response failed %v", err)
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(InvalidCredentials, "Passkey login response is invalid"))
		return
	}

	wa, err := newWebAuthn()

	if err != nil {
		log.Printf("creating webauthn relying party failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Passkey login is not configured"))
		return
	}

	var resolvedUser *webAuthnUser
	handler := func(rawID, userHandle []byte) (webauthn.User, error) {
		user, err := findWebAuthnUserForPasskey(ctx.Request.Context(), db, sessionRecord.RPID, rawID, userHandle)

		if err != nil {
			return nil, err
		}

		resolvedUser = user

		return user, nil
	}

	_, credential, err := wa.ValidatePasskeyLogin(handler, sessionRecord.SessionData, parsedResponse)

	if err != nil {
		log.Printf("validating passkey login failed %v", err)
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(InvalidCredentials, "Passkey login response is invalid"))
		return
	}

	if resolvedUser == nil {
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(InvalidCredentials, "Passkey login response is invalid"))
		return
	}

	if err := updateWebAuthnCredential(ctx.Request.Context(), db, resolvedUser.user.ID, sessionRecord.RPID, credential); err != nil {
		log.Printf("updating passkey credential failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Could not finish passkey login"))
		return
	}

	resp, err := CreateLoginResponseForUser(resolvedUser.user, authMethodPasskey)

	if err != nil {
		log.Printf("creating passkey login response failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Could not finish passkey login"))
		return
	}

	SetAuthCookie(ctx, resp.Token)
	resp.Token = ""

	ctx.JSON(http.StatusOK, resp)
}

func GetPasskeys(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	userSession := GetUserSessionFromContext(ctx)

	passkeys, err := getPasskeysForUser(ctx.Request.Context(), db, int(userSession.UserID), config.WebAuthnRPID)

	if err != nil {
		log.Printf("fetching passkeys failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Could not fetch passkeys"))
		return
	}

	ctx.JSON(http.StatusOK, passkeys)
}

func HandlePasskeyRegistrationBegin(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	userSession := GetUserSessionFromContext(ctx)

	if userSession.AuthMethod != "" && userSession.AuthMethod != authMethodPassword {
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(AuthorizationError, "Log in with your password before adding a passkey"))
		return
	}

	user, err := models.FindUser(ctx.Request.Context(), db, int(userSession.UserID))

	if err != nil {
		log.Printf("fetching passkey registration user failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Could not start passkey registration"))
		return
	}

	webAuthnUser, err := loadWebAuthnUser(ctx.Request.Context(), db, user, config.WebAuthnRPID, true)

	if err != nil {
		log.Printf("loading passkey registration user failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Could not start passkey registration"))
		return
	}

	wa, err := newWebAuthn()

	if err != nil {
		log.Printf("creating webauthn relying party failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Passkey registration is not configured"))
		return
	}

	creation, session, err := wa.BeginRegistration(
		webAuthnUser,
		webauthn.WithAuthenticatorSelection(protocol.AuthenticatorSelection{
			ResidentKey:      protocol.ResidentKeyRequirementRequired,
			UserVerification: protocol.VerificationRequired,
		}),
		webauthn.WithResidentKeyRequirement(protocol.ResidentKeyRequirementRequired),
		webauthn.WithExclusions(webauthn.Credentials(webAuthnUser.WebAuthnCredentials()).CredentialDescriptors()),
	)

	if err != nil {
		log.Printf("beginning passkey registration failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Could not start passkey registration"))
		return
	}

	userID := user.ID
	ceremonyID, err := saveWebAuthnSession(ctx.Request.Context(), db, passkeySessionKindRegistration, config.WebAuthnRPID, &userID, session)

	if err != nil {
		log.Printf("saving passkey registration session failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Could not start passkey registration"))
		return
	}

	ctx.JSON(http.StatusOK, PasskeyBeginResponse{
		CeremonyID: ceremonyID,
		Options:    creation.Response,
	})
}

func HandlePasskeyRegistrationFinish(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	userSession := GetUserSessionFromContext(ctx)

	if userSession.AuthMethod != "" && userSession.AuthMethod != authMethodPassword {
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(AuthorizationError, "Log in with your password before adding a passkey"))
		return
	}

	var body PasskeyFinishRequest

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Invalid passkey registration response"))
		return
	}

	sessionRecord, err := loadAndDeleteWebAuthnSession(ctx.Request.Context(), db, body.CeremonyID)

	if err != nil {
		log.Printf("passkey registration session not found or invalid ceremony_id=%q err=%v", body.CeremonyID, err)
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(InvalidInput, "Passkey registration expired or is invalid"))
		return
	}

	if sessionRecord.CeremonyKind != passkeySessionKindRegistration || !sessionRecord.UserID.Valid || sessionRecord.ExpiresAt.Before(time.Now()) {
		log.Printf(
			"passkey registration session rejected ceremony_id=%q kind=%q user_id_valid=%v user_id=%d expires_at=%s now=%s session_expires=%s",
			body.CeremonyID,
			sessionRecord.CeremonyKind,
			sessionRecord.UserID.Valid,
			sessionRecord.UserID.Int64,
			sessionRecord.ExpiresAt.Format(time.RFC3339Nano),
			time.Now().Format(time.RFC3339Nano),
			sessionRecord.SessionData.Expires.Format(time.RFC3339Nano),
		)
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(InvalidInput, "Passkey registration expired or is invalid"))
		return
	}

	if int32(sessionRecord.UserID.Int64) != userSession.UserID {
		log.Printf("passkey registration user mismatch ceremony_id=%q session_user_id=%d token_user_id=%d", body.CeremonyID, sessionRecord.UserID.Int64, userSession.UserID)
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(AuthorizationError, "Passkey registration does not belong to the current user"))
		return
	}

	parsedResponse, err := protocol.ParseCredentialCreationResponseBytes(body.Response)

	if err != nil {
		log.Printf("parsing passkey registration response failed %v", err)
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(InvalidInput, "Passkey registration response is invalid"))
		return
	}

	user, err := models.FindUser(ctx.Request.Context(), db, int(userSession.UserID))

	if err != nil {
		log.Printf("fetching passkey registration user failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Could not finish passkey registration"))
		return
	}

	webAuthnUser, err := loadWebAuthnUser(ctx.Request.Context(), db, user, sessionRecord.RPID, true)

	if err != nil {
		log.Printf("loading passkey registration user failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Could not finish passkey registration"))
		return
	}

	wa, err := newWebAuthn()

	if err != nil {
		log.Printf("creating webauthn relying party failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Passkey registration is not configured"))
		return
	}

	credential, err := wa.CreateCredential(webAuthnUser, sessionRecord.SessionData, parsedResponse)

	if err != nil {
		log.Printf("validating passkey registration failed %v", err)
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(InvalidInput, "Passkey registration response is invalid"))
		return
	}

	if err := insertWebAuthnCredential(ctx.Request.Context(), db, user.ID, sessionRecord.RPID, credential); err != nil {
		log.Printf("storing passkey credential failed %v", err)

		if errors.Is(err, errPasskeyCredentialAlreadyExists) {
			ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "This passkey is already registered"))
			return
		}

		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Could not store passkey"))
		return
	}

	ctx.JSON(http.StatusOK, PasskeyMetadata{
		AuthenticatorAttachment: string(credential.Authenticator.Attachment),
		CloneWarning:            credential.Authenticator.CloneWarning,
		CreatedAt:               time.Now(),
	})
}

func loadWebAuthnUser(ctx context.Context, db *sql.DB, user *models.User, rpid string, createHandle bool) (webAuthnUser, error) {
	var handle []byte
	var err error

	if createHandle {
		handle, err = ensureWebAuthnUserHandle(ctx, db, user.ID, rpid)
	} else {
		handle, err = getWebAuthnUserHandle(ctx, db, user.ID, rpid)
	}

	if err != nil {
		return webAuthnUser{}, err
	}

	credentials, err := loadWebAuthnCredentials(ctx, db, user.ID, rpid)

	if err != nil {
		return webAuthnUser{}, err
	}

	return webAuthnUser{
		user:        user,
		handle:      handle,
		credentials: credentials,
	}, nil
}

func ensureWebAuthnUserHandle(ctx context.Context, db *sql.DB, userID int, rpid string) ([]byte, error) {
	handle := make([]byte, 64)

	if _, err := rand.Read(handle); err != nil {
		return nil, err
	}

	_, err := db.ExecContext(ctx, `
		insert into webauthn_users (user_id, rpid, handle)
		values ($1, $2, $3)
		on conflict (rpid, user_id) do nothing
	`, userID, rpid, handle)

	if err != nil {
		return nil, err
	}

	return getWebAuthnUserHandle(ctx, db, userID, rpid)
}

func getWebAuthnUserHandle(ctx context.Context, db *sql.DB, userID int, rpid string) ([]byte, error) {
	var handle []byte

	err := db.QueryRowContext(ctx, `
		select handle
		from webauthn_users
		where user_id = $1 and rpid = $2
	`, userID, rpid).Scan(&handle)

	return handle, err
}

func loadWebAuthnCredentials(ctx context.Context, db *sql.DB, userID int, rpid string) ([]webauthn.Credential, error) {
	rows, err := db.QueryContext(ctx, `
		select credential_json::text
		from webauthn_credentials
		where user_id = $1 and rpid = $2
		order by id
	`, userID, rpid)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	credentials := []webauthn.Credential{}

	for rows.Next() {
		var credentialJSON []byte

		if err := rows.Scan(&credentialJSON); err != nil {
			return nil, err
		}

		var credential webauthn.Credential

		if err := json.Unmarshal(credentialJSON, &credential); err != nil {
			return nil, err
		}

		credentials = append(credentials, credential)
	}

	return credentials, rows.Err()
}

func insertWebAuthnCredential(ctx context.Context, db *sql.DB, userID int, rpid string, credential *webauthn.Credential) error {
	credentialJSON, err := json.Marshal(credential)

	if err != nil {
		return err
	}

	result, err := db.ExecContext(ctx, `
		insert into webauthn_credentials (user_id, rpid, credential_id, credential_json, name)
		values ($1, $2, $3, $4::jsonb, $5)
		on conflict (rpid, credential_id) do nothing
	`, userID, rpid, credential.ID, string(credentialJSON), string(credential.Authenticator.Attachment))

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errPasskeyCredentialAlreadyExists
	}

	return nil
}

func updateWebAuthnCredential(ctx context.Context, db *sql.DB, userID int, rpid string, credential *webauthn.Credential) error {
	credentialJSON, err := json.Marshal(credential)

	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, `
		update webauthn_credentials
		set credential_json = $1::jsonb, last_used_at = now()
		where user_id = $2 and rpid = $3 and credential_id = $4
	`, string(credentialJSON), userID, rpid, credential.ID)

	return err
}

func findWebAuthnUserForPasskey(ctx context.Context, db *sql.DB, rpid string, credentialID []byte, userHandle []byte) (*webAuthnUser, error) {
	var userID int

	err := db.QueryRowContext(ctx, `
		select wu.user_id
		from webauthn_users wu
		inner join webauthn_credentials wc on wc.user_id = wu.user_id and wc.rpid = wu.rpid
		where wu.rpid = $1 and wu.handle = $2 and wc.credential_id = $3
	`, rpid, userHandle, credentialID).Scan(&userID)

	if err != nil {
		return nil, err
	}

	user, err := models.Users(
		qm.Where("id = ?", userID),
	).One(ctx, db)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, sql.ErrNoRows
	}

	webAuthnUser, err := loadWebAuthnUser(ctx, db, user, rpid, false)

	if err != nil {
		return nil, err
	}

	return &webAuthnUser, nil
}

func getPasskeysForUser(ctx context.Context, db *sql.DB, userID int, rpid string) ([]PasskeyMetadata, error) {
	rows, err := db.QueryContext(ctx, `
		select id, coalesce(name, ''), credential_json::text, created_at, last_used_at
		from webauthn_credentials
		where user_id = $1 and rpid = $2
		order by created_at desc, id desc
	`, userID, rpid)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	passkeys := []PasskeyMetadata{}

	for rows.Next() {
		var passkey PasskeyMetadata
		var credentialJSON []byte
		var lastUsedAt sql.NullTime

		if err := rows.Scan(&passkey.ID, &passkey.Name, &credentialJSON, &passkey.CreatedAt, &lastUsedAt); err != nil {
			return nil, err
		}

		if lastUsedAt.Valid {
			passkey.LastUsedAt = &lastUsedAt.Time
		}

		var credential webauthn.Credential

		if err := json.Unmarshal(credentialJSON, &credential); err != nil {
			return nil, err
		}

		passkey.AuthenticatorAttachment = string(credential.Authenticator.Attachment)
		passkey.CloneWarning = credential.Authenticator.CloneWarning

		passkeys = append(passkeys, passkey)
	}

	return passkeys, rows.Err()
}

func saveWebAuthnSession(ctx context.Context, db *sql.DB, kind string, rpid string, userID *int, session *webauthn.SessionData) (string, error) {
	ceremonyIDBytes := make([]byte, 32)

	if _, err := rand.Read(ceremonyIDBytes); err != nil {
		return "", err
	}

	ceremonyID := base64.RawURLEncoding.EncodeToString(ceremonyIDBytes)

	if session.Expires.IsZero() {
		session.Expires = time.Now().Add(passkeySessionTTL)
	}

	sessionData, err := json.Marshal(session)

	if err != nil {
		return "", err
	}

	var nullableUserID interface{}

	if userID != nil {
		nullableUserID = *userID
	}

	_, _ = db.ExecContext(ctx, `delete from webauthn_sessions where expires_at < now()`)

	_, err = db.ExecContext(ctx, `
		insert into webauthn_sessions (ceremony_id, ceremony_kind, user_id, rpid, session_data, expires_at)
		values ($1, $2, $3, $4, $5, $6)
	`, ceremonyID, kind, nullableUserID, rpid, sessionData, session.Expires)

	return ceremonyID, err
}

func loadAndDeleteWebAuthnSession(ctx context.Context, db *sql.DB, ceremonyID string) (webAuthnSessionRecord, error) {
	var record webAuthnSessionRecord
	var sessionData []byte

	err := db.QueryRowContext(ctx, `
		delete from webauthn_sessions
		where ceremony_id = $1
		returning ceremony_kind, user_id, rpid, session_data, expires_at
	`, ceremonyID).Scan(&record.CeremonyKind, &record.UserID, &record.RPID, &sessionData, &record.ExpiresAt)

	if err != nil {
		return webAuthnSessionRecord{}, err
	}

	if err := json.Unmarshal(sessionData, &record.SessionData); err != nil {
		return webAuthnSessionRecord{}, err
	}

	if !record.SessionData.Expires.IsZero() {
		record.ExpiresAt = record.SessionData.Expires
	}

	return record, nil
}
