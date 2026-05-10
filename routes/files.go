package routes

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/config"
)

const maxUploadedFileSize = 8 * 1024 * 1024

type StoredFile struct {
	ID               int
	UploadedByUserID int
	FileName         string
	ContentType      string
	ContentHash      string
	SizeBytes        int
	CreatedAt        time.Time
}

type FileResponse struct {
	ID          int    `json:"id"`
	FileName    string `json:"fileName"`
	ContentType string `json:"contentType"`
	ContentHash string `json:"contentHash"`
	SizeBytes   int    `json:"sizeBytes"`
	CreatedAt   string `json:"createdAt"`
	URL         string `json:"url"`
}

type fileRouteError struct {
	status  int
	code    ErrorCode
	message string
}

type storeUploadedFileOptions struct {
	AllowedContentTypePrefixes []string
}

func fileURL(fileID int) string {
	return "/api/files/" + strconv.Itoa(fileID)
}

func hashFileContent(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func storedFilePath(contentHash string) string {
	return filepath.Join(config.FileStoragePath, contentHash)
}

func fileResponse(file StoredFile) FileResponse {
	return FileResponse{
		ID:          file.ID,
		FileName:    file.FileName,
		ContentType: file.ContentType,
		ContentHash: file.ContentHash,
		SizeBytes:   file.SizeBytes,
		CreatedAt:   file.CreatedAt.Format(time.RFC3339),
		URL:         fileURL(file.ID),
	}
}

func scanStoredFile(scanner interface {
	Scan(dest ...any) error
}) (StoredFile, error) {
	var file StoredFile

	err := scanner.Scan(
		&file.ID,
		&file.UploadedByUserID,
		&file.FileName,
		&file.ContentType,
		&file.ContentHash,
		&file.SizeBytes,
		&file.CreatedAt,
	)

	return file, err
}

func getStoredFileByID(ctx context.Context, db *sql.DB, fileID int) (StoredFile, error) {
	row := db.QueryRowContext(ctx, `
		select id, uploaded_by_user_id, file_name, content_type, content_hash, size_bytes, created_at
		from files
		where id = $1
	`, fileID)

	return scanStoredFile(row)
}

func readUploadedFile(fileHeader *multipart.FileHeader, options storeUploadedFileOptions) ([]byte, string, *fileRouteError) {
	if fileHeader.Size > maxUploadedFileSize {
		return nil, "", &fileRouteError{status: http.StatusBadRequest, code: InvalidInput, message: "File must be 8 MB or smaller"}
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, "", &fileRouteError{status: http.StatusBadRequest, code: InvalidInput, message: "Could not read file"}
	}
	defer file.Close()

	data, err := io.ReadAll(io.LimitReader(file, maxUploadedFileSize+1))
	if err != nil {
		return nil, "", &fileRouteError{status: http.StatusBadRequest, code: InvalidInput, message: "Could not read file"}
	}

	if len(data) > maxUploadedFileSize {
		return nil, "", &fileRouteError{status: http.StatusBadRequest, code: InvalidInput, message: "File must be 8 MB or smaller"}
	}

	contentType := http.DetectContentType(data)

	if len(options.AllowedContentTypePrefixes) > 0 {
		allowed := false

		for _, prefix := range options.AllowedContentTypePrefixes {
			if strings.HasPrefix(contentType, prefix) {
				allowed = true
				break
			}
		}

		if !allowed {
			return nil, "", &fileRouteError{status: http.StatusBadRequest, code: InvalidInput, message: "Unsupported file type"}
		}
	}

	return data, contentType, nil
}

func storeUploadedFile(ctx context.Context, db *sql.DB, uploadedByUserID int32, fileHeader *multipart.FileHeader, options storeUploadedFileOptions) (StoredFile, *fileRouteError) {
	data, contentType, routeErr := readUploadedFile(fileHeader, options)
	if routeErr != nil {
		return StoredFile{}, routeErr
	}

	if err := os.MkdirAll(config.FileStoragePath, 0755); err != nil {
		return StoredFile{}, &fileRouteError{status: http.StatusInternalServerError, code: InternalServerError, message: "Failed to prepare file storage"}
	}

	contentHash := hashFileContent(data)
	diskPath := storedFilePath(contentHash)

	fileAlreadyExists := false
	if _, err := os.Stat(diskPath); err == nil {
		fileAlreadyExists = true
	} else if !os.IsNotExist(err) {
		return StoredFile{}, &fileRouteError{status: http.StatusInternalServerError, code: InternalServerError, message: "Failed to check file storage"}
	}

	if !fileAlreadyExists {
		if err := os.WriteFile(diskPath, data, 0644); err != nil {
			return StoredFile{}, &fileRouteError{status: http.StatusInternalServerError, code: InternalServerError, message: "Failed to write file"}
		}
	}

	row := db.QueryRowContext(ctx, `
		insert into files (uploaded_by_user_id, file_name, content_type, size_bytes, content_hash)
		values ($1, $2, $3, $4, $5)
		returning id, uploaded_by_user_id, file_name, content_type, content_hash, size_bytes, created_at
	`, uploadedByUserID, fileHeader.Filename, contentType, len(data), contentHash)

	storedFile, err := scanStoredFile(row)
	if err != nil {
		if !fileAlreadyExists {
			_ = os.Remove(diskPath)
		}

		return StoredFile{}, &fileRouteError{status: http.StatusInternalServerError, code: InternalServerError, message: "Failed to save file metadata"}
	}

	return storedFile, nil
}

func removeUnreferencedFiles(ctx context.Context, db *sql.DB, contentHashes []string) {
	seen := map[string]bool{}

	for _, contentHash := range contentHashes {
		if contentHash == "" || seen[contentHash] {
			continue
		}

		seen[contentHash] = true

		var referenceCount int
		if err := db.QueryRowContext(ctx, "select count(*) from files where content_hash = $1", contentHash).Scan(&referenceCount); err != nil {
			continue
		}

		if referenceCount == 0 {
			_ = os.Remove(storedFilePath(contentHash))
		}
	}
}

func deleteStoredFileRecord(ctx context.Context, db *sql.DB, fileID int) {
	var contentHash string
	if err := db.QueryRowContext(ctx, "delete from files where id = $1 returning content_hash", fileID).Scan(&contentHash); err != nil {
		return
	}

	removeUnreferencedFiles(ctx, db, []string{contentHash})
}

func UploadFile(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	userSession := GetUserSessionFromContext(ctx)

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "No file provided"))
		return
	}

	storedFile, routeErr := storeUploadedFile(ctx.Request.Context(), db, userSession.UserID, fileHeader, storeUploadedFileOptions{})
	if routeErr != nil {
		ctx.JSON(routeErr.status, CreateErrorResponse(routeErr.code, routeErr.message))
		return
	}

	ctx.JSON(http.StatusOK, fileResponse(storedFile))
}

func GetFile(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	fileID, err := strconv.Atoi(ctx.Params.ByName("fileId"))

	if err != nil || fileID == 0 {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "No file id provided"))
		return
	}

	storedFile, err := getStoredFileByID(ctx.Request.Context(), db, fileID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, CreateErrorResponse(NotFound, "File not found"))
		return
	}

	file, err := os.Open(storedFilePath(storedFile.ContentHash))
	if err != nil {
		ctx.JSON(http.StatusNotFound, CreateErrorResponse(NotFound, "File not found"))
		return
	}
	defer file.Close()

	ctx.Header("Content-Type", storedFile.ContentType)
	ctx.Header("Content-Disposition", "inline; filename=\""+strings.ReplaceAll(storedFile.FileName, "\"", "")+"\"")
	http.ServeContent(ctx.Writer, ctx.Request, storedFile.FileName, storedFile.CreatedAt, file)
}

func DeleteFile(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	userSession := GetUserSessionFromContext(ctx)
	fileID, err := strconv.Atoi(ctx.Params.ByName("fileId"))

	if err != nil || fileID == 0 {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "No file id provided"))
		return
	}

	var uploadedByUserID int
	var diaryReferenceCount int
	var contentHash string

	err = db.QueryRowContext(ctx.Request.Context(), `
		select files.uploaded_by_user_id, files.content_hash, count(diary_entry_pictures.id)
		from files
		left join diary_entry_pictures on diary_entry_pictures.file_id = files.id
		where files.id = $1
		group by files.id
	`, fileID).Scan(&uploadedByUserID, &contentHash, &diaryReferenceCount)

	if err != nil {
		ctx.JSON(http.StatusNotFound, CreateErrorResponse(NotFound, "File not found"))
		return
	}

	if uploadedByUserID != int(userSession.UserID) {
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(AuthorizationError, "Only the uploader can delete this file"))
		return
	}

	if diaryReferenceCount > 0 {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "File is still attached"))
		return
	}

	if _, err := db.ExecContext(ctx.Request.Context(), "delete from files where id = $1", fileID); err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to delete file"))
		return
	}

	removeUnreferencedFiles(ctx.Request.Context(), db, []string{contentHash})

	ctx.JSON(http.StatusOK, struct{}{})
}
