package routes

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/auth"
)

type APIKeyMetadata struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Prefix     string     `json:"prefix"`
	CreatedAt  time.Time  `json:"createdAt"`
	LastUsedAt *time.Time `json:"lastUsedAt"`
	RevokedAt  *time.Time `json:"revokedAt"`
}

type CreateAPIKeyRequest struct {
	Name string `json:"name"`
}

type CreateAPIKeyResponse struct {
	APIKeyMetadata
	Token string `json:"token"`
}

func GetAPIKeys(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	userSession := GetUserSessionFromContext(ctx)

	keys, err := getAPIKeysForUser(ctx.Request.Context(), db, int(userSession.UserID))

	if err != nil {
		log.Printf("fetching api keys failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Could not fetch API keys"))
		return
	}

	ctx.JSON(http.StatusOK, keys)
}

func CreateAPIKey(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	userSession := GetUserSessionFromContext(ctx)

	var body CreateAPIKeyRequest
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Invalid payload"))
		return
	}

	name := strings.TrimSpace(body.Name)
	if err := auth.ValidateAPIKeyName(name); err != nil {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, err.Error()))
		return
	}

	token, prefix, tokenHash, err := auth.GenerateAPIKey()
	if err != nil {
		log.Printf("generating api key failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Could not generate API key"))
		return
	}

	var key APIKeyMetadata
	var lastUsedAt sql.NullTime
	var revokedAt sql.NullTime
	err = db.QueryRowContext(ctx.Request.Context(), `
		insert into api_keys (user_id, name, prefix, token_hash)
		values ($1, $2, $3, $4)
		returning id, name, prefix, created_at, last_used_at, revoked_at
	`, int(userSession.UserID), name, prefix, tokenHash).Scan(
		&key.ID,
		&key.Name,
		&key.Prefix,
		&key.CreatedAt,
		&lastUsedAt,
		&revokedAt,
	)
	if err != nil {
		log.Printf("storing api key failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Could not store API key"))
		return
	}

	key.LastUsedAt = nullTimePtr(lastUsedAt)
	key.RevokedAt = nullTimePtr(revokedAt)

	ctx.JSON(http.StatusOK, CreateAPIKeyResponse{
		APIKeyMetadata: key,
		Token:          token,
	})
}

func RevokeAPIKey(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	userSession := GetUserSessionFromContext(ctx)

	apiKeyID, err := strconv.Atoi(ctx.Param("apiKeyId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Invalid API key id"))
		return
	}

	result, err := db.ExecContext(ctx.Request.Context(), `
		update api_keys
		set revoked_at = now()
		where id = $1 and user_id = $2 and revoked_at is null
	`, apiKeyID, int(userSession.UserID))

	if err != nil {
		log.Printf("revoking api key failed %v", err)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Could not revoke API key"))
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("checking revoked api key rows failed %v", err)
	}

	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, CreateErrorResponse(InvalidInput, "API key not found"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ok": true})
}

func AuthenticateAPIKey(ctx *gin.Context, db *sql.DB, token string) (UserSession, bool) {
	if !auth.LooksLikeAPIKey(token) {
		return UserSession{}, false
	}

	tokenHash := auth.HashAPIKey(token)

	var userID int
	var name sql.NullString
	var username sql.NullString

	err := db.QueryRowContext(ctx.Request.Context(), `
		select u.id, u.name, u.username
		from api_keys ak
		inner join users u on u.id = ak.user_id
		where ak.token_hash = $1 and ak.revoked_at is null
	`, tokenHash).Scan(&userID, &name, &username)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("api key lookup failed %v", err)
		}
		return UserSession{}, false
	}

	if _, err := db.ExecContext(ctx.Request.Context(), `
		update api_keys set last_used_at = now() where token_hash = $1
	`, tokenHash); err != nil {
		log.Printf("updating api key last_used_at failed %v", err)
	}

	userName := name.String
	if userName == "" {
		userName = username.String
	}

	return UserSession{
		UserID:     int32(userID),
		UserName:   userName,
		AuthMethod: authMethodAPIKey,
	}, true
}

func getAPIKeysForUser(ctx context.Context, db *sql.DB, userID int) ([]APIKeyMetadata, error) {
	rows, err := db.QueryContext(ctx, `
		select id, name, prefix, created_at, last_used_at, revoked_at
		from api_keys
		where user_id = $1
		order by created_at desc
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	keys := []APIKeyMetadata{}
	for rows.Next() {
		var key APIKeyMetadata
		var lastUsedAt sql.NullTime
		var revokedAt sql.NullTime
		if err := rows.Scan(&key.ID, &key.Name, &key.Prefix, &key.CreatedAt, &lastUsedAt, &revokedAt); err != nil {
			return nil, err
		}
		key.LastUsedAt = nullTimePtr(lastUsedAt)
		key.RevokedAt = nullTimePtr(revokedAt)
		keys = append(keys, key)
	}

	return keys, rows.Err()
}

func nullTimePtr(value sql.NullTime) *time.Time {
	if !value.Valid {
		return nil
	}

	return &value.Time
}
