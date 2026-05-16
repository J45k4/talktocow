package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/auth"
	"github.com/j45k4/talktocow/chatroom"
	"github.com/j45k4/talktocow/config"
	"github.com/j45k4/talktocow/eventbus"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

type e2eHarness struct {
	t       *testing.T
	db      *sql.DB
	server  http.Handler
	userIDs []int
}

func TestBackendE2EAuthFilesAndDiary(t *testing.T) {
	if os.Getenv("TALKTOCOW_E2E") != "1" {
		t.Skip("set TALKTOCOW_E2E=1 to run backend e2e tests against the configured Postgres database")
	}

	h := newE2EHarness(t)
	defer h.cleanup()

	username := fmt.Sprintf("e2e_%d", time.Now().UnixNano())
	password := "test-password"
	h.createUser(username, password)

	tokenLogin := h.jsonRequest("POST", "/api/token", map[string]string{
		"username": username,
		"password": password,
	}, nil)
	assertStatus(t, tokenLogin, http.StatusOK)

	var tokenBody struct {
		Token    string `json:"token"`
		UserID   string `json:"userId"`
		Username string `json:"username"`
	}
	decodeJSON(t, tokenLogin.Body, &tokenBody)
	if tokenBody.Token == "" {
		t.Fatalf("/api/token did not return a token: %#v", tokenBody)
	}
	if tokenBody.Username != username {
		t.Fatalf("/api/token returned username %q, want %q", tokenBody.Username, username)
	}

	cookieRefresh := h.request("POST", "/api/session/cookie", nil, map[string]string{
		"Authorization": "Bearer " + tokenBody.Token,
	})
	assertStatus(t, cookieRefresh, http.StatusNoContent)
	refreshedCookies := cookieRefresh.Result().Cookies()
	if len(refreshedCookies) == 0 {
		t.Fatalf("/api/session/cookie did not set an auth cookie")
	}

	fileID := h.uploadJPEGWithBearer(tokenBody.Token)
	medium := h.request("GET", fmt.Sprintf("/api/files/%d?size=medium", fileID), nil, map[string]string{
		"Authorization": "Bearer " + tokenBody.Token,
	})
	assertStatus(t, medium, http.StatusOK)
	if contentType := medium.Header().Get("Content-Type"); contentType != "image/jpeg" {
		t.Fatalf("medium content type = %q, want image/jpeg", contentType)
	}

	img, err := jpeg.Decode(medium.Body)
	if err != nil {
		t.Fatalf("decode medium jpeg: %v", err)
	}
	if got, want := img.Bounds().Dx(), 960; got != want {
		t.Fatalf("medium width = %d, want %d", got, want)
	}
	if got, want := img.Bounds().Dy(), 1280; got != want {
		t.Fatalf("medium height = %d, want %d", got, want)
	}

	mediumWithRefreshedCookie := h.request("GET", fmt.Sprintf("/api/files/%d?size=medium", fileID), nil, nil, refreshedCookies...)
	assertStatus(t, mediumWithRefreshedCookie, http.StatusOK)
	if contentType := mediumWithRefreshedCookie.Header().Get("Content-Type"); contentType != "image/jpeg" {
		t.Fatalf("medium with refreshed cookie content type = %q, want image/jpeg", contentType)
	}
	if mediumWithRefreshedCookie.Body.Len() == 0 {
		t.Fatalf("medium with refreshed cookie response was empty")
	}

	cookieLogin := h.jsonRequest("POST", "/api/login", map[string]string{
		"username": username,
		"password": password,
	}, nil)
	assertStatus(t, cookieLogin, http.StatusOK)
	if token := jsonField(t, cookieLogin.Body.Bytes(), "token"); token != "" {
		t.Fatalf("/api/login returned token %q; browser login should rely on HttpOnly cookie", token)
	}

	cookies := cookieLogin.Result().Cookies()
	if len(cookies) == 0 {
		t.Fatalf("/api/login did not set an auth cookie")
	}

	diary := h.jsonRequest("POST", "/api/diary/entry", map[string]any{
		"title":          "E2E diary entry",
		"body":           "Uploaded through backend e2e test",
		"createdAt":      "2026-05-11T12:00:00Z",
		"pictureFileIds": []int{fileID},
	}, cookies)
	assertStatus(t, diary, http.StatusOK)
	if id := jsonField(t, diary.Body.Bytes(), "id"); id == "" {
		t.Fatalf("diary create response did not include id: %s", diary.Body.String())
	}
}

func TestBackendE2EDiaryImageVisibleToOtherAuthenticatedUser(t *testing.T) {
	if os.Getenv("TALKTOCOW_E2E") != "1" {
		t.Skip("set TALKTOCOW_E2E=1 to run backend e2e tests against the configured Postgres database")
	}

	h := newE2EHarness(t)
	defer h.cleanup()

	stamp := time.Now().UnixNano()
	ownerUsername := fmt.Sprintf("e2e_owner_%d", stamp)
	viewerUsername := fmt.Sprintf("e2e_viewer_%d", stamp)
	password := "test-password"
	h.createUser(ownerUsername, password)
	h.createUser(viewerUsername, password)

	ownerCookies := h.loginCookies(ownerUsername, password)
	viewerCookies := h.loginCookies(viewerUsername, password)

	fileID := h.uploadJPEGWithCookies(ownerCookies)
	diary := h.jsonRequest("POST", "/api/diary/entry", map[string]any{
		"title":          "E2E shared diary image",
		"body":           "Uploaded by one user and viewed by another",
		"createdAt":      "2026-05-13T05:00:00Z",
		"pictureFileIds": []int{fileID},
	}, ownerCookies)
	assertStatus(t, diary, http.StatusOK)

	entryID, err := strconv.Atoi(jsonField(t, diary.Body.Bytes(), "id"))
	if err != nil || entryID == 0 {
		t.Fatalf("parse created diary entry id: %v; body=%s", err, diary.Body.String())
	}

	entryAsViewer := h.request("GET", fmt.Sprintf("/api/diary/entry/%d", entryID), nil, nil, viewerCookies...)
	assertStatus(t, entryAsViewer, http.StatusOK)
	if got := jsonField(t, entryAsViewer.Body.Bytes(), "pictureCount"); got != "1" {
		t.Fatalf("viewer diary pictureCount = %q, want 1; body=%s", got, entryAsViewer.Body.String())
	}

	picturesAsViewer := h.request("GET", fmt.Sprintf("/api/diary/entry/%d/pictures", entryID), nil, nil, viewerCookies...)
	assertStatus(t, picturesAsViewer, http.StatusOK)

	imageAsViewer := h.request("GET", fmt.Sprintf("/api/files/%d?size=medium", fileID), nil, nil, viewerCookies...)
	assertStatus(t, imageAsViewer, http.StatusOK)
	if contentType := imageAsViewer.Header().Get("Content-Type"); contentType != "image/jpeg" {
		t.Fatalf("viewer image content type = %q, want image/jpeg", contentType)
	}
	if imageAsViewer.Body.Len() == 0 {
		t.Fatalf("viewer image response was empty")
	}
}

func newE2EHarness(t *testing.T) *e2eHarness {
	t.Helper()

	oldFileStoragePath := config.FileStoragePath
	config.FileStoragePath = t.TempDir()
	t.Cleanup(func() {
		config.FileStoragePath = oldFileStoragePath
	})

	connectionString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", config.DBName, config.DBUser, config.DBPassword, config.DbHost, config.DBPort)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	migrations := &migrate.FileMigrationSource{Dir: "./migrations"}
	if _, err := migrate.Exec(db, "postgres", migrations, migrate.Up); err != nil {
		t.Fatalf("run migrations: %v", err)
	}

	gin.SetMode(gin.TestMode)
	return &e2eHarness{
		t:      t,
		db:     db,
		server: setupRouter(db, eventbus.New(), chatroom.NewChatroomEventbus()),
	}
}

func (h *e2eHarness) createUser(username, password string) {
	h.t.Helper()

	hash, err := auth.HashPassword(password)
	if err != nil {
		h.t.Fatalf("hash password: %v", err)
	}

	var userID int
	err = h.db.QueryRowContext(context.Background(), `
		insert into users (name, username, password_hash, created_at)
		values ($1, $1, $2, now())
		returning id
	`, username, hash).Scan(&userID)
	if err != nil {
		h.t.Fatalf("insert e2e user: %v", err)
	}
	h.userIDs = append(h.userIDs, userID)
}

func (h *e2eHarness) cleanup() {
	if len(h.userIDs) == 0 {
		return
	}

	ctx := context.Background()
	for _, userID := range h.userIDs {
		_, _ = h.db.ExecContext(ctx, `delete from diary_entry_comments where user_id = $1`, userID)
		_, _ = h.db.ExecContext(ctx, `delete from diary_entry_pictures where diary_entry_id in (select id from diary_entries where who_posted_user_id = $1)`, userID)
		_, _ = h.db.ExecContext(ctx, `delete from shared_diary_entries where diary_entry_id in (select id from diary_entries where who_posted_user_id = $1)`, userID)
		_, _ = h.db.ExecContext(ctx, `delete from diary_entries where who_posted_user_id = $1`, userID)
		_, _ = h.db.ExecContext(ctx, `delete from files where uploaded_by_user_id = $1`, userID)
		_, _ = h.db.ExecContext(ctx, `delete from users where id = $1`, userID)
	}
}

func (h *e2eHarness) loginCookies(username, password string) []*http.Cookie {
	h.t.Helper()

	response := h.jsonRequest("POST", "/api/login", map[string]string{
		"username": username,
		"password": password,
	}, nil)
	assertStatus(h.t, response, http.StatusOK)

	cookies := response.Result().Cookies()
	if len(cookies) == 0 {
		h.t.Fatalf("/api/login did not set an auth cookie for %s", username)
	}
	return cookies
}

func (h *e2eHarness) jsonRequest(method, path string, body any, cookies []*http.Cookie) *httptest.ResponseRecorder {
	h.t.Helper()

	payload, err := json.Marshal(body)
	if err != nil {
		h.t.Fatalf("marshal request: %v", err)
	}

	return h.request(method, path, bytes.NewReader(payload), map[string]string{
		"Content-Type": "application/json",
	}, cookies...)
}

func (h *e2eHarness) request(method, path string, body io.Reader, headers map[string]string, cookies ...*http.Cookie) *httptest.ResponseRecorder {
	h.t.Helper()

	req := httptest.NewRequest(method, path, body)
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	w := httptest.NewRecorder()
	h.server.ServeHTTP(w, req)
	return w
}

func (h *e2eHarness) uploadJPEGWithBearer(token string) int {
	h.t.Helper()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("file", "oriented.jpg")
	if err != nil {
		h.t.Fatalf("create multipart file: %v", err)
	}
	if _, err := part.Write(orientedJPEGFixture(h.t)); err != nil {
		h.t.Fatalf("write multipart file: %v", err)
	}
	if err := writer.Close(); err != nil {
		h.t.Fatalf("close multipart writer: %v", err)
	}

	response := h.request("POST", "/api/files", &body, map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  writer.FormDataContentType(),
	})
	assertStatus(h.t, response, http.StatusOK)

	fileID, err := strconv.Atoi(jsonField(h.t, response.Body.Bytes(), "id"))
	if err != nil {
		h.t.Fatalf("parse uploaded file id: %v; body=%s", err, response.Body.String())
	}
	return fileID
}

func (h *e2eHarness) uploadJPEGWithCookies(cookies []*http.Cookie) int {
	h.t.Helper()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("file", "diary.jpg")
	if err != nil {
		h.t.Fatalf("create multipart file: %v", err)
	}
	if _, err := part.Write(orientedJPEGFixture(h.t)); err != nil {
		h.t.Fatalf("write multipart file: %v", err)
	}
	if err := writer.Close(); err != nil {
		h.t.Fatalf("close multipart writer: %v", err)
	}

	response := h.request("POST", "/api/files", &body, map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}, cookies...)
	assertStatus(h.t, response, http.StatusOK)

	fileID, err := strconv.Atoi(jsonField(h.t, response.Body.Bytes(), "id"))
	if err != nil {
		h.t.Fatalf("parse uploaded file id: %v; body=%s", err, response.Body.String())
	}
	return fileID
}

func orientedJPEGFixture(t *testing.T) []byte {
	t.Helper()

	img := image.NewRGBA(image.Rect(0, 0, 1600, 1200))
	for y := 0; y < 1200; y++ {
		for x := 0; x < 1600; x++ {
			img.Set(x, y, color.RGBA{R: uint8(x % 256), G: uint8(y % 256), B: 180, A: 255})
		}
	}

	var encoded bytes.Buffer
	if err := jpeg.Encode(&encoded, img, &jpeg.Options{Quality: 90}); err != nil {
		t.Fatalf("encode fixture jpeg: %v", err)
	}

	return injectEXIFOrientation(t, encoded.Bytes(), 6)
}

func injectEXIFOrientation(t *testing.T, jpegData []byte, orientation uint16) []byte {
	t.Helper()

	if len(jpegData) < 2 || jpegData[0] != 0xff || jpegData[1] != 0xd8 {
		t.Fatalf("fixture is not a jpeg")
	}

	var tiff bytes.Buffer
	tiff.WriteString("MM")
	_ = binary.Write(&tiff, binary.BigEndian, uint16(42))
	_ = binary.Write(&tiff, binary.BigEndian, uint32(8))
	_ = binary.Write(&tiff, binary.BigEndian, uint16(1))
	_ = binary.Write(&tiff, binary.BigEndian, uint16(0x0112))
	_ = binary.Write(&tiff, binary.BigEndian, uint16(3))
	_ = binary.Write(&tiff, binary.BigEndian, uint32(1))
	_ = binary.Write(&tiff, binary.BigEndian, orientation)
	_ = binary.Write(&tiff, binary.BigEndian, uint16(0))
	_ = binary.Write(&tiff, binary.BigEndian, uint32(0))

	exifPayload := append([]byte("Exif\x00\x00"), tiff.Bytes()...)
	segmentLength := len(exifPayload) + 2
	if segmentLength > 0xffff {
		t.Fatalf("exif segment too large")
	}

	var out bytes.Buffer
	out.Write(jpegData[:2])
	out.Write([]byte{0xff, 0xe1, byte(segmentLength >> 8), byte(segmentLength)})
	out.Write(exifPayload)
	out.Write(jpegData[2:])
	return out.Bytes()
}

func assertStatus(t *testing.T, response *httptest.ResponseRecorder, want int) {
	t.Helper()
	if response.Code != want {
		t.Fatalf("status = %d, want %d; body=%s", response.Code, want, response.Body.String())
	}
}

func decodeJSON(t *testing.T, body *bytes.Buffer, target any) {
	t.Helper()
	if err := json.Unmarshal(body.Bytes(), target); err != nil {
		t.Fatalf("decode json %q: %v", body.String(), err)
	}
}

func jsonField(t *testing.T, data []byte, field string) string {
	t.Helper()

	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("decode json %q: %v", string(data), err)
	}

	value, ok := payload[field]
	if !ok || value == nil {
		return ""
	}

	switch typed := value.(type) {
	case string:
		return typed
	case float64:
		return strconv.Itoa(int(typed))
	default:
		return fmt.Sprint(typed)
	}
}

func TestMain(m *testing.M) {
	if os.Getenv("TALKTOCOW_E2E") == "1" {
		if cwd, err := os.Getwd(); err == nil {
			_ = os.Setenv("PATH", filepath.Join(cwd, "node_modules", ".bin")+string(os.PathListSeparator)+os.Getenv("PATH"))
		}
	}
	os.Exit(m.Run())
}
