package routes

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
	"io"
	"math"
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
const imageVariantJPEGQuality = 86

var imageVariantMaxDimension = map[string]int{
	"thumb":  360,
	"medium": 1280,
	"large":  2048,
}

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

func storedFileVariantPath(contentHash string, size string) string {
	return filepath.Join(config.FileStoragePath, "variants", contentHash+"_"+size+".jpg")
}

func fileVariantFromRequest(ctx *gin.Context) (string, *fileRouteError) {
	size := ctx.Query("size")

	if size == "" {
		size = ctx.Query("variant")
	}

	if size == "" || size == "original" {
		return "original", nil
	}

	if _, ok := imageVariantMaxDimension[size]; ok {
		return size, nil
	}

	return "", &fileRouteError{status: http.StatusBadRequest, code: InvalidInput, message: "Unsupported file size"}
}

func isResizableRasterImage(contentType string) bool {
	return contentType == "image/jpeg" || contentType == "image/png"
}

func sanitizeContentDispositionFileName(fileName string) string {
	replacer := strings.NewReplacer(
		"\"", "",
		"\r", "",
		"\n", "",
		"\\", "",
		"/", "",
	)

	return replacer.Replace(fileName)
}

func variantFileName(fileName string) string {
	extension := filepath.Ext(fileName)

	if extension == "" {
		return fileName + ".jpg"
	}

	return strings.TrimSuffix(fileName, extension) + ".jpg"
}

func imageFitsMaxDimension(bounds image.Rectangle, maxDimension int) bool {
	return bounds.Dx() <= maxDimension && bounds.Dy() <= maxDimension
}

func scaledImageSize(bounds image.Rectangle, maxDimension int) (int, int) {
	width := bounds.Dx()
	height := bounds.Dy()

	if width <= 0 || height <= 0 {
		return 0, 0
	}

	scale := math.Min(1, float64(maxDimension)/float64(max(width, height)))

	return max(1, int(math.Round(float64(width)*scale))), max(1, int(math.Round(float64(height)*scale)))
}

func rgbaOverWhite(img image.Image, x int, y int) color.RGBA {
	source := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
	alpha := uint32(source.A)
	inverseAlpha := uint32(255 - source.A)

	return color.RGBA{
		R: uint8((uint32(source.R)*alpha + 255*inverseAlpha) / 255),
		G: uint8((uint32(source.G)*alpha + 255*inverseAlpha) / 255),
		B: uint8((uint32(source.B)*alpha + 255*inverseAlpha) / 255),
		A: 255,
	}
}

func resizeImageToJPEG(img image.Image, maxDimension int, output io.Writer) error {
	bounds := img.Bounds()
	targetWidth, targetHeight := scaledImageSize(bounds, maxDimension)
	if targetWidth == 0 || targetHeight == 0 {
		return image.ErrFormat
	}

	resized := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))
	sourceWidth := bounds.Dx()
	sourceHeight := bounds.Dy()

	for y := 0; y < targetHeight; y++ {
		sourceY := bounds.Min.Y + int(float64(y)*float64(sourceHeight)/float64(targetHeight))
		if sourceY >= bounds.Max.Y {
			sourceY = bounds.Max.Y - 1
		}

		for x := 0; x < targetWidth; x++ {
			sourceX := bounds.Min.X + int(float64(x)*float64(sourceWidth)/float64(targetWidth))
			if sourceX >= bounds.Max.X {
				sourceX = bounds.Max.X - 1
			}

			resized.SetRGBA(x, y, rgbaOverWhite(img, sourceX, sourceY))
		}
	}

	return jpeg.Encode(output, resized, &jpeg.Options{Quality: imageVariantJPEGQuality})
}

func ensureImageVariant(storedFile StoredFile, size string) (string, bool, error) {
	maxDimension, ok := imageVariantMaxDimension[size]
	if !ok || !isResizableRasterImage(storedFile.ContentType) {
		return storedFilePath(storedFile.ContentHash), false, nil
	}

	variantPath := storedFileVariantPath(storedFile.ContentHash, size)
	if _, err := os.Stat(variantPath); err == nil {
		return variantPath, true, nil
	} else if !os.IsNotExist(err) {
		return "", false, err
	}

	source, err := os.Open(storedFilePath(storedFile.ContentHash))
	if err != nil {
		return "", false, err
	}
	defer source.Close()

	img, _, err := image.Decode(source)
	if err != nil {
		return storedFilePath(storedFile.ContentHash), false, nil
	}

	if imageFitsMaxDimension(img.Bounds(), maxDimension) {
		return storedFilePath(storedFile.ContentHash), false, nil
	}

	variantDir := filepath.Dir(variantPath)
	if err := os.MkdirAll(variantDir, 0755); err != nil {
		return "", false, err
	}

	tempFile, err := os.CreateTemp(variantDir, storedFile.ContentHash+"_"+size+"_*.jpg")
	if err != nil {
		return "", false, err
	}

	tempPath := tempFile.Name()
	cleanupTemp := true
	defer func() {
		if cleanupTemp {
			_ = os.Remove(tempPath)
		}
	}()

	if err := resizeImageToJPEG(img, maxDimension, tempFile); err != nil {
		_ = tempFile.Close()
		return "", false, err
	}

	if err := tempFile.Close(); err != nil {
		return "", false, err
	}

	if err := os.Rename(tempPath, variantPath); err != nil {
		return "", false, err
	}

	cleanupTemp = false
	return variantPath, true, nil
}

func removeStoredFileVariants(contentHash string) {
	matches, err := filepath.Glob(storedFileVariantPath(contentHash, "*"))
	if err != nil {
		return
	}

	for _, match := range matches {
		_ = os.Remove(match)
	}
}

func serveStoredFile(ctx *gin.Context, storedFile StoredFile) {
	size, routeErr := fileVariantFromRequest(ctx)
	if routeErr != nil {
		ctx.JSON(routeErr.status, CreateErrorResponse(routeErr.code, routeErr.message))
		return
	}

	filePath := storedFilePath(storedFile.ContentHash)
	contentType := storedFile.ContentType
	fileName := storedFile.FileName

	if size != "original" {
		variantPath, isVariant, err := ensureImageVariant(storedFile, size)
		if err != nil {
			ctx.JSON(http.StatusNotFound, CreateErrorResponse(NotFound, "File not found"))
			return
		}

		filePath = variantPath

		if isVariant {
			contentType = "image/jpeg"
			fileName = variantFileName(fileName)
		}
	}

	file, err := os.Open(filePath)
	if err != nil {
		ctx.JSON(http.StatusNotFound, CreateErrorResponse(NotFound, "File not found"))
		return
	}
	defer file.Close()

	ctx.Header("Content-Type", contentType)
	ctx.Header("Content-Disposition", "inline; filename=\""+sanitizeContentDispositionFileName(fileName)+"\"")
	ctx.Header("X-Content-Type-Options", "nosniff")
	http.ServeContent(ctx.Writer, ctx.Request, fileName, storedFile.CreatedAt, file)
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
			removeStoredFileVariants(contentHash)
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

	serveStoredFile(ctx, storedFile)
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
