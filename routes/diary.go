package routes

import (
	"database/sql"
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
	Label          *string `json:"label"`
	StartsAt       *string `json:"startsAt"`
	EndsAt         *string `json:"endsAt"`
	AllDay         bool    `json:"allDay"`
}

type CreateDiaryEntryRequest struct {
	Title     string  `json:"title"`
	Body      string  `json:"body"`
	CreatedAt string  `json:"createdAt"`
	Label     *string `json:"label"`
	StartsAt  *string `json:"startsAt"`
	EndsAt    *string `json:"endsAt"`
	AllDay    bool    `json:"allDay"`
}

type UpdateDiaryEntryRequest struct {
	Title    string   `json:"title"`
	Body     string   `json:"body"`
	Offset   int      `json:"offset"`
	Mask     []string `json:"mask"`
	Label    *string  `json:"label"`
	StartsAt *string  `json:"startsAt"`
	EndsAt   *string  `json:"endsAt"`
	AllDay   bool     `json:"allDay"`
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
	var startsAt sql.NullTime
	var endsAt sql.NullTime
	var createdAt time.Time
	var postedByUserID int

	err := scanner.Scan(
		&entry.ID,
		&title,
		&body,
		&postedByUserID,
		&createdAt,
		&label,
		&startsAt,
		&endsAt,
		&entry.AllDay,
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

	if startsAt.Valid {
		formatted := formatDiaryTime(startsAt.Time)
		entry.StartsAt = &formatted
	}

	if endsAt.Valid {
		formatted := formatDiaryTime(endsAt.Time)
		entry.EndsAt = &formatted
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

	var startsAt *time.Time

	if createRequest.StartsAt != nil && *createRequest.StartsAt != "" {
		parsed, err := parseDiaryTime(*createRequest.StartsAt)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Invalid diary entry start time"))
			return
		}

		startsAt = &parsed
	}

	var endsAt *time.Time

	if createRequest.EndsAt != nil && *createRequest.EndsAt != "" {
		parsed, err := parseDiaryTime(*createRequest.EndsAt)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "Invalid diary entry end time"))
			return
		}

		if startsAt != nil && parsed.Before(*startsAt) {
			ctx.JSON(http.StatusBadRequest, CreateErrorResponse(InvalidInput, "End time must be after start time"))
			return
		}

		endsAt = &parsed
	}

	title := createRequest.Title
	if title == "" && createRequest.Label != nil {
		title = *createRequest.Label
	}

	row := db.QueryRowContext(ctx.Request.Context(), `
		insert into diary_entries (title, body, who_posted_user_id, created_at, label, starts_at, ends_at, all_day)
		values ($1, $2, $3, $4, $5, $6, $7, $8)
		returning id, title, body, who_posted_user_id, created_at, label, starts_at, ends_at, all_day
	`, title, createRequest.Body, userSession.UserID, createdAt, createRequest.Label, startsAt, endsAt, createRequest.AllDay)

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
		select id, title, body, who_posted_user_id, created_at, label, starts_at, ends_at, all_day
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
		    label = case when $6 then $7 else label end,
		    all_day = case when $8 then $9 else all_day end
		where id = $1
	`,
		entryId,
		DoesMaskHaveField(updateDiaryEntryRequest.Mask, "title"), updateDiaryEntryRequest.Title,
		DoesMaskHaveField(updateDiaryEntryRequest.Mask, "body"), updateDiaryEntryRequest.Body,
		DoesMaskHaveField(updateDiaryEntryRequest.Mask, "label"), updateDiaryEntryRequest.Label,
		DoesMaskHaveField(updateDiaryEntryRequest.Mask, "allDay"), updateDiaryEntryRequest.AllDay,
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

func GetDiaryEntries(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	offset, limit := GetOffsetAndLimit(ctx, 0, 30)

	rows, err := db.QueryContext(ctx.Request.Context(), `
		select id, title, body, who_posted_user_id, created_at, label, starts_at, ends_at, all_day
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
