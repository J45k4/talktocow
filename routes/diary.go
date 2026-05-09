package routes

import (
	"bytes"
	"database/sql"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type DiaryEntry struct {
	ID             int     `json:"id"`
	Title          string  `json:"title"`
	Body           string  `json:"body"`
	PostedByUserID string  `json:"postedByUserId"`
	CreateAt       string  `json:"createdAt"`
	Label          *string `json:"label"`
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
		returning id, title, body, who_posted_user_id, created_at, label
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
		select id, title, body, who_posted_user_id, created_at, label
		from diary_entries
		where id = $1
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

	_, err = tx.ExecContext(ctx.Request.Context(), "delete from diary_entry_pictures where diary_entry_id = $1", entryId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to delete diary entry pictures"))
		return
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
		select id, diary_entry_id, file_name, content_type, created_at
		from diary_entry_pictures
		where diary_entry_id = $1
		order by created_at asc, id asc
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

	row := db.QueryRowContext(ctx.Request.Context(), `
		insert into diary_entry_pictures (diary_entry_id, uploaded_by_user_id, file_name, content_type, image_data)
		values ($1, $2, $3, $4, $5)
		returning id, diary_entry_id, file_name, content_type, created_at
	`, entryId, userSession.UserID, fileHeader.Filename, contentType, data)

	picture, err := scanDiaryEntryPicture(row)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to save diary entry picture"))
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
	var data []byte
	var createdAt time.Time

	err := db.QueryRowContext(ctx.Request.Context(), `
		select file_name, content_type, image_data, created_at
		from diary_entry_pictures
		where id = $1 and diary_entry_id = $2
	`, pictureId, entryId).Scan(&fileName, &contentType, &data, &createdAt)

	if err != nil {
		ctx.JSON(http.StatusNotFound, CreateErrorResponse(NotFound, "Picture not found"))
		return
	}

	ctx.Header("Content-Type", contentType)
	ctx.Header("Content-Disposition", "inline; filename=\""+strings.ReplaceAll(fileName, "\"", "")+"\"")
	http.ServeContent(ctx.Writer, ctx.Request, fileName, createdAt, bytes.NewReader(data))
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

	result, err := db.ExecContext(ctx.Request.Context(), `
		delete from diary_entry_pictures
		where id = $1 and diary_entry_id = $2
	`, pictureId, entryId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to delete picture"))
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil || rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, CreateErrorResponse(NotFound, "Picture not found"))
		return
	}

	ctx.JSON(http.StatusOK, struct{}{})
}

func GetDiaryEntries(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	offset, limit := GetOffsetAndLimit(ctx, 0, 30)

	rows, err := db.QueryContext(ctx.Request.Context(), `
		select id, title, body, who_posted_user_id, created_at, label
		from diary_entries
		order by created_at desc
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
