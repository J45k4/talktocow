package routes

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DiaryEntry struct {
	ID             int     `json:"id"`
	Title          string  `json:"title"`
	Body           string  `json:"body"`
	PostedByUserID string  `json:"postedByUserId"`
	CreateAt       string  `json:"createdAt"`
	CanEditDate    bool    `json:"canEditDate"`
	Label          *string `json:"label"`
	PictureCount   int     `json:"pictureCount"`
}

type DiaryEntryPicture struct {
	ID           int    `json:"id"`
	DiaryEntryID int    `json:"diaryEntryId"`
	FileID       int    `json:"fileId"`
	FileName     string `json:"fileName"`
	ContentType  string `json:"contentType"`
	CreatedAt    string `json:"createdAt"`
	URL          string `json:"url"`
}

type CreateDiaryEntryRequest struct {
	Title          string  `json:"title"`
	Body           string  `json:"body"`
	CreatedAt      string  `json:"createdAt"`
	Label          *string `json:"label"`
	PictureFileIDs []int   `json:"pictureFileIds"`
}

type UpdateDiaryEntryRequest struct {
	Title          string   `json:"title"`
	Body           string   `json:"body"`
	CreatedAt      string   `json:"createdAt"`
	Offset         int      `json:"offset"`
	Mask           []string `json:"mask"`
	Label          *string  `json:"label"`
	PictureFileIDs []int    `json:"pictureFileIds"`
}

const diaryEntryDateEditWindow = 7 * 24 * time.Hour

var (
	errDiaryEntryDateEditExpired = errors.New("diary entry date edit window expired")
	errInvalidDiaryEntryDate     = errors.New("invalid diary entry date")
)

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

func canEditDiaryEntryDate(createdAt time.Time, now time.Time) bool {
	return now.Sub(createdAt) < diaryEntryDateEditWindow
}

func validateDiaryEntryDateEdit(existingCreatedAt time.Time, requestedCreatedAt string, now time.Time) (time.Time, error) {
	if !canEditDiaryEntryDate(existingCreatedAt, now) {
		return time.Time{}, errDiaryEntryDateEditExpired
	}

	parsed, err := parseDiaryTime(requestedCreatedAt)
	if err != nil {
		return time.Time{}, errInvalidDiaryEntryDate
	}

	return parsed, nil
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
	entry.CanEditDate = canEditDiaryEntryDate(createdAt, time.Now())

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

	tx, err := db.BeginTx(ctx.Request.Context(), nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to create diary entry"))
		return
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(ctx.Request.Context(), `
		insert into diary_entries (title, body, who_posted_user_id, created_at, label)
		values ($1, $2, $3, $4, $5)
		returning id, title, body, who_posted_user_id, created_at, label, 0 as picture_count
	`, title, createRequest.Body, userSession.UserID, createdAt, createRequest.Label)

	entry, err := scanDiaryEntry(row)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to create diary entry"))
		return
	}

	pictureCount, err := attachFilesToDiaryEntry(ctx, tx, entry.ID, createRequest.PictureFileIDs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Invalid diary entry picture files"))
		return
	}

	entry.PictureCount = pictureCount

	if err := tx.Commit(); err != nil {
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
	userSession := GetUserSessionFromContext(ctx)

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

	tx, err := db.BeginTx(ctx.Request.Context(), nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to update diary entry"))
		return
	}
	defer tx.Rollback()

	var postedByUserID int
	var existingCreatedAt time.Time

	err = tx.QueryRowContext(ctx.Request.Context(), `
		select who_posted_user_id, created_at
		from diary_entries
		where id = $1
	`, entryId).Scan(&postedByUserID, &existingCreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, CreateErrorResponse(NotFound, "Diary entry not found"))
			return
		}

		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to update diary entry"))
		return
	}

	if postedByUserID != int(userSession.UserID) {
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(AuthorizationError, "Only the author can edit this diary entry"))
		return
	}

	shouldUpdateCreatedAt := DoesMaskHaveField(updateDiaryEntryRequest.Mask, "createdAt")
	updatedCreatedAt := time.Time{}

	if shouldUpdateCreatedAt {
		updatedCreatedAt, err = validateDiaryEntryDateEdit(existingCreatedAt, updateDiaryEntryRequest.CreatedAt, time.Now())

		if err != nil {
			if errors.Is(err, errDiaryEntryDateEditExpired) {
				ctx.JSON(http.StatusForbidden, CreateErrorResponse(AuthorizationError, "Diary entry dates can only be changed during the first 7 days"))
				return
			}

			ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Invalid diary entry date"))
			return
		}
	}

	_, err = tx.ExecContext(ctx.Request.Context(), `
		update diary_entries
		set title = case when $2 then $3 else title end,
		    body = case when $4 then $5 else body end,
		    label = case when $6 then $7 else label end,
		    created_at = case when $8 then $9 else created_at end
		where id = $1
	`,
		entryId,
		DoesMaskHaveField(updateDiaryEntryRequest.Mask, "title"), updateDiaryEntryRequest.Title,
		DoesMaskHaveField(updateDiaryEntryRequest.Mask, "body"), updateDiaryEntryRequest.Body,
		DoesMaskHaveField(updateDiaryEntryRequest.Mask, "label"), updateDiaryEntryRequest.Label,
		shouldUpdateCreatedAt, updatedCreatedAt,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to update diary entry"))
		return
	}

	if DoesMaskHaveField(updateDiaryEntryRequest.Mask, "pictureFileIds") {
		if _, err := tx.ExecContext(ctx.Request.Context(), "delete from diary_entry_pictures where diary_entry_id = $1", entryId); err != nil {
			ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to update diary entry pictures"))
			return
		}

		if _, err := attachFilesToDiaryEntry(ctx, tx, entryId, updateDiaryEntryRequest.PictureFileIDs); err != nil {
			ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Invalid diary entry picture files"))
			return
		}
	}

	if err := tx.Commit(); err != nil {
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
		select files.id, files.content_hash
		from files
		join diary_entry_pictures on diary_entry_pictures.file_id = files.id
		where diary_entry_pictures.diary_entry_id = $1
	`, entryId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to find diary entry pictures"))
		return
	}

	fileIds := []int{}
	contentHashes := []string{}

	for fileRows.Next() {
		var fileId int
		var contentHash string

		if err := fileRows.Scan(&fileId, &contentHash); err != nil {
			fileRows.Close()
			ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to read diary entry pictures"))
			return
		}

		fileIds = append(fileIds, fileId)
		contentHashes = append(contentHashes, contentHash)
	}

	fileRows.Close()

	_, err = tx.ExecContext(ctx.Request.Context(), "delete from diary_entry_pictures where diary_entry_id = $1", entryId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to delete diary entry pictures"))
		return
	}

	for _, fileId := range fileIds {
		if err := deleteFileRecordIfNoDiaryReferences(ctx, tx, fileId); err != nil {
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

	removeUnreferencedFiles(ctx.Request.Context(), db, contentHashes)

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

func uniquePositiveInts(values []int) []int {
	seen := map[int]bool{}
	result := []int{}

	for _, value := range values {
		if value <= 0 || seen[value] {
			continue
		}

		seen[value] = true
		result = append(result, value)
	}

	return result
}

func attachFilesToDiaryEntry(ctx *gin.Context, tx *sql.Tx, entryId int, fileIDs []int) (int, error) {
	userSession := GetUserSessionFromContext(ctx)
	uniqueFileIDs := uniquePositiveInts(fileIDs)

	for _, fileID := range uniqueFileIDs {
		var exists bool
		if err := tx.QueryRowContext(ctx.Request.Context(), `
			select exists(
				select 1
				from files
				where id = $1
					and uploaded_by_user_id = $2
					and content_type like 'image/%'
			)
		`, fileID, userSession.UserID).Scan(&exists); err != nil {
			return 0, err
		}

		if !exists {
			return 0, errors.New("file not found")
		}

		if _, err := tx.ExecContext(ctx.Request.Context(), `
			insert into diary_entry_pictures (diary_entry_id, file_id)
			values ($1, $2)
			on conflict do nothing
		`, entryId, fileID); err != nil {
			return 0, err
		}
	}

	return len(uniqueFileIDs), nil
}

func deleteFileRecordIfNoDiaryReferences(ctx *gin.Context, tx *sql.Tx, fileID int) error {
	var referenceCount int
	if err := tx.QueryRowContext(ctx.Request.Context(), "select count(*) from diary_entry_pictures where file_id = $1", fileID).Scan(&referenceCount); err != nil {
		return err
	}

	if referenceCount > 0 {
		return nil
	}

	_, err := tx.ExecContext(ctx.Request.Context(), "delete from files where id = $1", fileID)
	return err
}

func scanDiaryEntryPicture(scanner interface {
	Scan(dest ...any) error
}) (DiaryEntryPicture, error) {
	var picture DiaryEntryPicture
	var createdAt time.Time

	err := scanner.Scan(&picture.ID, &picture.DiaryEntryID, &picture.FileID, &picture.FileName, &picture.ContentType, &createdAt)

	if err != nil {
		return picture, err
	}

	picture.CreatedAt = formatDiaryTime(createdAt)
	picture.URL = fileURL(picture.FileID)

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
		select diary_entry_pictures.id, diary_entry_pictures.diary_entry_id, diary_entry_pictures.file_id, files.file_name, files.content_type, diary_entry_pictures.created_at
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

	storedFile, routeErr := storeUploadedFile(ctx.Request.Context(), db, userSession.UserID, fileHeader, storeUploadedFileOptions{
		AllowedContentTypePrefixes: []string{"image/"},
	})
	if routeErr != nil {
		ctx.JSON(routeErr.status, CreateErrorResponse(routeErr.code, routeErr.message))
		return
	}

	tx, err := db.BeginTx(ctx.Request.Context(), nil)

	if err != nil {
		deleteStoredFileRecord(ctx.Request.Context(), db, storedFile.ID)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to save picture"))
		return
	}

	defer tx.Rollback()

	row := tx.QueryRowContext(ctx.Request.Context(), `
		insert into diary_entry_pictures (diary_entry_id, file_id)
		values ($1, $2)
		returning id, diary_entry_id, file_id, $3::text as file_name, $4::text as content_type, created_at
	`, entryId, storedFile.ID, storedFile.FileName, storedFile.ContentType)

	picture, err := scanDiaryEntryPicture(row)

	if err != nil {
		_ = tx.Rollback()
		deleteStoredFileRecord(ctx.Request.Context(), db, storedFile.ID)
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to save diary entry picture"))
		return
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		deleteStoredFileRecord(ctx.Request.Context(), db, storedFile.ID)
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
	var contentHash string
	var createdAt time.Time

	err := db.QueryRowContext(ctx.Request.Context(), `
		select files.file_name, files.content_type, files.content_hash, diary_entry_pictures.created_at
		from diary_entry_pictures
		join files on files.id = diary_entry_pictures.file_id
		where diary_entry_pictures.id = $1 and diary_entry_pictures.diary_entry_id = $2
	`, pictureId, entryId).Scan(&fileName, &contentType, &contentHash, &createdAt)

	if err != nil {
		ctx.JSON(http.StatusNotFound, CreateErrorResponse(NotFound, "Picture not found"))
		return
	}

	serveStoredFile(ctx, StoredFile{
		FileName:    fileName,
		ContentType: contentType,
		ContentHash: contentHash,
		CreatedAt:   createdAt,
	})
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
	var contentHash string

	err = db.QueryRowContext(ctx.Request.Context(), `
		select files.id, files.content_hash
		from diary_entry_pictures
		join files on files.id = diary_entry_pictures.file_id
		where diary_entry_pictures.id = $1 and diary_entry_pictures.diary_entry_id = $2
	`, pictureId, entryId).Scan(&fileId, &contentHash)

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

	if err := deleteFileRecordIfNoDiaryReferences(ctx, tx, fileId); err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to delete file metadata"))
		return
	}

	if err := tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, "Failed to finish picture delete"))
		return
	}

	removeUnreferencedFiles(ctx.Request.Context(), db, []string{contentHash})

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
