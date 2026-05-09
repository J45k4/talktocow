package routes

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/config"
)

type DiaryEntry struct {
	ID             int     `json:"id"`
	Title          string  `json:"title"`
	Body           string  `json:"body"`
	PostedByUserID string  `json:"postedByUserId"`
	CreateAt       string  `json:"createdAt"`
	Label          *string `json:"label"`
	PictureCount   int     `json:"pictureCount"`
}

type DiaryEntryPicture struct {
	ID           int    `json:"id"`
	DiaryEntryID int    `json:"diaryEntryId"`
	FileName     string `json:"fileName"`
	ContentType  string `json:"contentType"`
	CreatedAt    string `json:"createdAt"`
	URL          string `json:"url"`
}

type CreateDiaryEntryRequest struct {
	Title     string  `json:"title"`
	Body      string  `json:"body"`
	CreatedAt string  `json:"createdAt"`
	Label     *string `json:"label"`
}

type UpdateDiaryEntryRequest struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Offset int      `json:"offset"`
	Mask   []string `json:"mask"`
	Label  *string  `json:"label"`
}

func parseDiaryTime(value string) (time.Time, error) {
	if parsed, err := time.Parse(time.RFC3339, value); err == nil {
		return parsed, nil
	}

	if parsed, err := time.Parse("2006-01-02T15:04", value); err == nil {
		return parsed, nil
	}

	return time.Parse("2006-01-02", value)
}

func formatDiaryTime(value time.Time) string {
	return value.Format(time.RFC3339)
}

func scanDiaryEntry(scanner interface {
	Scan(dest ...any) error
}) (DiaryEntry, error) {
	var entry DiaryEntry
	var title sql.NullString
	var body sql.NullString
	var label sql.NullString
	var createdAt time.Time
	var postedByUserID int

	err := scanner.Scan(
		&entry.ID,
		&title,
		&body,
		&postedByUserID,
		&createdAt,
		&label,
		&entry.PictureCount,
	)

	if err != nil {
		return entry, err
	}

	entry.Title = title.String
	entry.Body = body.String
	entry.PostedByUserID = strconv.Itoa(postedByUserID)
	entry.CreateAt = formatDiaryTime(createdAt)

	if label.Valid {
		entry.Label = &label.String
	}

	return entry, nil
}

func CreateDiaryEntry(ctx *gin.Context) {
	userSession := GetUserSessionFromContext(ctx)
	db := GetDBFromContext(ctx)

	var createRequest CreateDiaryEntryRequest

	if err := ctx.BindJSON(&createRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Invalid diary entry payload"))
		return
	}

	createdAt := time.Now()

	if createRequest.CreatedAt != "" {
		parsed, err := parseDiaryTime(createRequest.CreatedAt)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Invalid diary entry date"))
			return
		}

		createdAt = parsed
	}

	title := createRequest.Title
	if title == "" && createRequest.Label != nil {
		title = *createRequest.Label
	}

	row := db.QueryRowContext(ctx.Request.Context(), `
		insert into diary_entries (title, body, who_posted_user_id, created_at, label)
		values ($1, $2, $3, $4, $5)
		returning id, title, body, who_posted_user_id, created_at, label, 0 as picture_count
	`, title, createRequest.Body, userSession.UserID, createdAt, createRequest.Label)

	entry, err := scanDiaryEntry(row)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to create diary entry"))
		return
	}

	ctx.JSON(http.StatusOK, entry)
}

func GetDiaryEntry(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	entryId, err := strconv.Atoi(ctx.Params.ByName("diaryEntryId"))

	if err != nil || entryId == 0 {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "No entry id provided"))
		return
	}

	row := db.QueryRowContext(ctx.Request.Context(), `
		select diary_entries.id, title, body, who_posted_user_id, diary_entries.created_at, label, count(diary_entry_pictures.id) as picture_count
		from diary_entries
		left join diary_entry_pictures on diary_entry_pictures.diary_entry_id = diary_entries.id
		where diary_entries.id = $1
		group by diary_entries.id
	`, entryId)

	diaryEntry, err := scanDiaryEntry(row)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, ""))
		return
	}

	ctx.JSON(http.StatusOK, diaryEntry)
}

func UpdateDiaryEntry(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	entryId, err := strconv.Atoi(ctx.Params.ByName("diaryEntryId"))

	if err != nil || entryId == 0 {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "No entry id provided"))
		return
	}

	var updateDiaryEntryRequest UpdateDiaryEntryRequest

	if err := ctx.BindJSON(&updateDiaryEntryRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Invalid payload"))
		return
	}

	_, err = db.ExecContext(ctx.Request.Context(), `
		update diary_entries
		set title = case when $2 then $3 else title end,
		    body = case when $4 then $5 else body end,
		    label = case when $6 then $7 else label end
		where id = $1
	`,
		entryId,
		DoesMaskHaveField(updateDiaryEntryRequest.Mask, "title"), updateDiaryEntryRequest.Title,
		DoesMaskHaveField(updateDiaryEntryRequest.Mask, "body"), updateDiaryEntryRequest.Body,
		DoesMaskHaveField(updateDiaryEntryRequest.Mask, "label"), updateDiaryEntryRequest.Label,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to update diary entry"))
		return
	}

	GetDiaryEntry(ctx)
}

func DeleteDiaryEntry(ctx *gin.Context) {
	userSession := GetUserSessionFromContext(ctx)
	db := GetDBFromContext(ctx)

	entryId, _ := strconv.Atoi(ctx.Params.ByName("diaryEntryId"))

	if entryId == 0 {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "No entry id provided"))
		return
	}

	var postedByUserID int

	err := db.QueryRowContext(ctx.Request.Context(), "select who_posted_user_id from diary_entries where id = $1", entryId).Scan(&postedByUserID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, ""))
		return
	}

	if postedByUserID != int(userSession.UserID) {
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(AuthorizationError, "Only the author can delete this diary entry"))
		return
	}

	tx, err := db.BeginTx(ctx.Request.Context(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to start diary entry delete"))
		return
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx.Request.Context(), "delete from diary_entry_comments where diary_entry_id = $1", entryId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to delete diary entry comments"))
		return
	}

	_, err = tx.ExecContext(ctx.Request.Context(), "delete from shared_diary_entries where diary_entry_id = $1", entryId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to delete shared diary entries"))
		return
	}

	fileRows, err := tx.QueryContext(ctx.Request.Context(), `
		select files.id, files.disk_path
		from files
		join diary_entry_pictures on diary_entry_pictures.file_id = files.id
		where diary_entry_pictures.diary_entry_id = $1
	`, entryId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to find diary entry pictures"))
		return
	}

	fileIds := []int{}
	filePaths := []string{}

	for fileRows.Next() {
		var fileId int
		var filePath string

		if err := fileRows.Scan(&fileId, &filePath); err != nil {
			fileRows.Close()
			ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to read diary entry pictures"))
			return
		}

		fileIds = append(fileIds, fileId)
		filePaths = append(filePaths, filePath)
	}

	fileRows.Close()

	_, err = tx.ExecContext(ctx.Request.Context(), "delete from diary_entry_pictures where diary_entry_id = $1", entryId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to delete diary entry pictures"))
		return
	}

	for _, fileId := range fileIds {
		if _, err := tx.ExecContext(ctx.Request.Context(), "delete from files where id = $1", fileId); err != nil {
			ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to delete diary entry files"))
			return
		}
	}

	_, err = tx.ExecContext(ctx.Request.Context(), "delete from diary_entries where id = $1", entryId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to delete diary entry"))
		return
	}

	if err := tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to finish diary entry delete"))
		return
	}

	removeFiles(filePaths)

	ctx.JSON(http.StatusOK, struct{}{})
}

func userCanEditDiaryEntry(ctx *gin.Context, entryId int) (bool, error) {
	userSession := GetUserSessionFromContext(ctx)
	db := GetDBFromContext(ctx)

	var postedByUserID int
	err := db.QueryRowContext(ctx.Request.Context(), "select who_posted_user_id from diary_entries where id = $1", entryId).Scan(&postedByUserID)

	if err != nil {
		return false, err
	}

	return postedByUserID == int(userSession.UserID), nil
}

func pictureURL(entryId int, pictureId int) string {
	return "/api/diary/entry/" + strconv.Itoa(entryId) + "/picture/" + strconv.Itoa(pictureId)
}

func createStoredFileName(originalFileName string) (string, error) {
	randomBytes := make([]byte, 16)

	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}

	extension := strings.ToLower(filepath.Ext(originalFileName))

	return hex.EncodeToString(randomBytes) + extension, nil
}

func removeFiles(paths []string) {
	for _, path := range paths {
		if path != "" {
			_ = os.Remove(path)
		}
	}
}

func scanDiaryEntryPicture(scanner interface {
	Scan(dest ...any) error
}) (DiaryEntryPicture, error) {
	var picture DiaryEntryPicture
	var createdAt time.Time

	err := scanner.Scan(&picture.ID, &picture.DiaryEntryID, &picture.FileName, &picture.ContentType, &createdAt)

	if err != nil {
		return picture, err
	}

	picture.CreatedAt = formatDiaryTime(createdAt)
	picture.URL = pictureURL(picture.DiaryEntryID, picture.ID)

	return picture, nil
}

func GetDiaryEntryPictures(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	entryId, err := strconv.Atoi(ctx.Params.ByName("diaryEntryId"))

	if err != nil || entryId == 0 {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "No entry id provided"))
		return
	}

	rows, err := db.QueryContext(ctx.Request.Context(), `
		select diary_entry_pictures.id, diary_entry_pictures.diary_entry_id, files.file_name, files.content_type, diary_entry_pictures.created_at
		from diary_entry_pictures
		join files on files.id = diary_entry_pictures.file_id
		where diary_entry_pictures.diary_entry_id = $1
		order by diary_entry_pictures.created_at asc, diary_entry_pictures.id asc
	`, entryId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to fetch diary entry pictures"))
		return
	}

	defer rows.Close()

	pictures := []DiaryEntryPicture{}

	for rows.Next() {
		picture, err := scanDiaryEntryPicture(rows)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to read diary entry pictures"))
			return
		}

		pictures = append(pictures, picture)
	}

	ctx.JSON(http.StatusOK, pictures)
}

func UploadDiaryEntryPicture(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	userSession := GetUserSessionFromContext(ctx)
	entryId, err := strconv.Atoi(ctx.Params.ByName("diaryEntryId"))

	if err != nil || entryId == 0 {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "No entry id provided"))
		return
	}

	canEdit, err := userCanEditDiaryEntry(ctx, entryId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, CreateErrorResponse(NotFound, "Diary entry not found"))
		return
	}

	if !canEdit {
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(AuthorizationError, "Only the author can add pictures"))
		return
	}

	fileHeader, err := ctx.FormFile("picture")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "No picture provided"))
		return
	}

	if fileHeader.Size > 8*1024*1024 {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Picture must be 8 MB or smaller"))
		return
	}

	file, err := fileHeader.Open()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Could not read picture"))
		return
	}

	defer file.Close()

	data, err := io.ReadAll(io.LimitReader(file, 8*1024*1024+1))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Could not read picture"))
		return
	}

	if len(data) > 8*1024*1024 {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Picture must be 8 MB or smaller"))
		return
	}

	contentType := http.DetectContentType(data)

	if !strings.HasPrefix(contentType, "image/") {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Only image files are supported"))
		return
	}

	if err := os.MkdirAll(config.FileStoragePath, 0755); err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to prepare file storage"))
		return
	}

	storedFileName, err := createStoredFileName(fileHeader.Filename)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to create file name"))
		return
	}

	diskPath := filepath.Join(config.FileStoragePath, storedFileName)

	if err := os.WriteFile(diskPath, data, 0644); err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to write picture"))
		return
	}

	tx, err := db.BeginTx(ctx.Request.Context(), nil)

	if err != nil {
		_ = os.Remove(diskPath)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to save picture"))
		return
	}

	defer tx.Rollback()

	var fileId int
	err = tx.QueryRowContext(ctx.Request.Context(), `
		insert into files (uploaded_by_user_id, file_name, content_type, disk_path, size_bytes)
		values ($1, $2, $3, $4, $5)
		returning id
	`, userSession.UserID, fileHeader.Filename, contentType, diskPath, len(data)).Scan(&fileId)

	if err != nil {
		_ = os.Remove(diskPath)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to save file metadata"))
		return
	}

	row := tx.QueryRowContext(ctx.Request.Context(), `
		insert into diary_entry_pictures (diary_entry_id, file_id)
		values ($1, $2)
		returning id, diary_entry_id, $3::text as file_name, $4::text as content_type, created_at
	`, entryId, fileId, fileHeader.Filename, contentType)

	picture, err := scanDiaryEntryPicture(row)

	if err != nil {
		_ = os.Remove(diskPath)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to save diary entry picture"))
		return
	}

	if err := tx.Commit(); err != nil {
		_ = os.Remove(diskPath)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to finish picture upload"))
		return
	}

	ctx.JSON(http.StatusOK, picture)
}

func GetDiaryEntryPicture(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	entryId, entryErr := strconv.Atoi(ctx.Params.ByName("diaryEntryId"))
	pictureId, pictureErr := strconv.Atoi(ctx.Params.ByName("pictureId"))

	if entryErr != nil || pictureErr != nil || entryId == 0 || pictureId == 0 {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "No picture id provided"))
		return
	}

	var fileName string
	var contentType string
	var diskPath string
	var createdAt time.Time

	err := db.QueryRowContext(ctx.Request.Context(), `
		select files.file_name, files.content_type, files.disk_path, diary_entry_pictures.created_at
		from diary_entry_pictures
		join files on files.id = diary_entry_pictures.file_id
		where diary_entry_pictures.id = $1 and diary_entry_pictures.diary_entry_id = $2
	`, pictureId, entryId).Scan(&fileName, &contentType, &diskPath, &createdAt)

	if err != nil {
		ctx.JSON(http.StatusNotFound, CreateErrorResponse(NotFound, "Picture not found"))
		return
	}

	file, err := os.Open(diskPath)

	if err != nil {
		ctx.JSON(http.StatusNotFound, CreateErrorResponse(NotFound, "Picture file not found"))
		return
	}

	defer file.Close()

	ctx.Header("Content-Type", contentType)
	ctx.Header("Content-Disposition", "inline; filename=\""+strings.ReplaceAll(fileName, "\"", "")+"\"")
	http.ServeContent(ctx.Writer, ctx.Request, fileName, createdAt, file)
}

func DeleteDiaryEntryPicture(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	entryId, entryErr := strconv.Atoi(ctx.Params.ByName("diaryEntryId"))
	pictureId, pictureErr := strconv.Atoi(ctx.Params.ByName("pictureId"))

	if entryErr != nil || pictureErr != nil || entryId == 0 || pictureId == 0 {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "No picture id provided"))
		return
	}

	canEdit, err := userCanEditDiaryEntry(ctx, entryId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, CreateErrorResponse(NotFound, "Diary entry not found"))
		return
	}

	if !canEdit {
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(AuthorizationError, "Only the author can delete pictures"))
		return
	}

	var fileId int
	var diskPath string

	err = db.QueryRowContext(ctx.Request.Context(), `
		select files.id, files.disk_path
		from diary_entry_pictures
		join files on files.id = diary_entry_pictures.file_id
		where diary_entry_pictures.id = $1 and diary_entry_pictures.diary_entry_id = $2
	`, pictureId, entryId).Scan(&fileId, &diskPath)

	if err != nil {
		ctx.JSON(http.StatusNotFound, CreateErrorResponse(NotFound, "Picture not found"))
		return
	}

	tx, err := db.BeginTx(ctx.Request.Context(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to start picture delete"))
		return
	}

	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx.Request.Context(), "delete from diary_entry_pictures where id = $1 and diary_entry_id = $2", pictureId, entryId); err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to delete picture"))
		return
	}

	if _, err := tx.ExecContext(ctx.Request.Context(), "delete from files where id = $1", fileId); err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to delete file metadata"))
		return
	}

	if err := tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to finish picture delete"))
		return
	}

	removeFiles([]string{diskPath})

	ctx.JSON(http.StatusOK, struct{}{})
}

func GetDiaryEntries(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	offset, limit := GetOffsetAndLimit(ctx, 0, 30)

	rows, err := db.QueryContext(ctx.Request.Context(), `
		select diary_entries.id, title, body, who_posted_user_id, diary_entries.created_at, label, count(diary_entry_pictures.id) as picture_count
		from diary_entries
		left join diary_entry_pictures on diary_entry_pictures.diary_entry_id = diary_entries.id
		group by diary_entries.id
		order by diary_entries.created_at desc
		offset $1 limit $2
	`, offset, limit)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, ""))
		return
	}

	defer rows.Close()

	entries := []DiaryEntry{}

	for rows.Next() {
		entry, err := scanDiaryEntry(rows)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, ""))
			return
		}

		entries = append(entries, entry)
	}

	ctx.JSON(http.StatusOK, entries)
}

func GetDiaryLabels(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	rows, err := db.QueryContext(ctx.Request.Context(), `
		select distinct label
		from diary_entries
		where label is not null and label <> ''
		order by label asc
	`)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to fetch labels"))
		return
	}

	defer rows.Close()

	labels := []string{}

	for rows.Next() {
		var label string
		if err := rows.Scan(&label); err != nil {
			ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to read labels"))
			return
		}

		labels = append(labels, label)
	}

	ctx.JSON(http.StatusOK, labels)
}

func GetDiaryEntriesCount(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	var count int
	err := db.QueryRowContext(ctx.Request.Context(), "select count(*) from diary_entries").Scan(&count)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, ""))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}
