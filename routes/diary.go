package routes

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type DiaryEntry struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	CreateAt time.Time `json:"createdAt"`
}

type CreateDiaryEntryRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func CreateDiaryEntry(ctx *gin.Context) {
	userSession := GetUserSessionFromContext(ctx)

	db := GetDBFromContext(ctx)

	var createRequest CreateDiaryEntryRequest

	ctx.BindJSON(&createRequest)

	entry := models.DiaryEntry{
		Title:           null.StringFrom(createRequest.Title),
		Body:            null.StringFrom(createRequest.Body),
		WhoPostedUserID: int(userSession.UserID),
	}

	err := entry.Insert(ctx.Request.Context(), db, boil.Infer())

	if err != nil {
		panic(err)
	}

	ctx.JSON(200, entry)
}

func GetDiaryEntry(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	entryId, err := strconv.Atoi(ctx.Params.ByName("diaryEntryId"))

	if entryId == 0 {
		fmt.Errorf("no entry id provided")

		ctx.JSON(400, CreateErrorResponse(InvalidInput, "No entry id provided"))
		return
	}

	fmt.Printf("entryId %v", entryId)

	diaryEntry, err := models.FindDiaryEntry(ctx.Request.Context(), db, entryId)

	if err != nil {
		fmt.Errorf("Finding diary entry failed %v", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))
		return
	}

	ctx.JSON(200, diaryEntry)
}

type UpdateDiaryEntryRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Offset int    `json:"offset"`
	Mask   []string
}

func UpdateDiaryEntry(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	entryId, err := strconv.Atoi(ctx.Params.ByName("diaryEntryId"))

	if entryId == 0 {
		fmt.Errorf("no entry id provided")

		ctx.JSON(400, CreateErrorResponse(InvalidInput, "No entry id provided"))
		return
	}

	var updateDiaryEntryRequest UpdateDiaryEntryRequest

	err = ctx.BindJSON(&updateDiaryEntryRequest)

	if err != nil {
		fmt.Errorf("parsing update payload failed %v", err)

		ctx.JSON(400, CreateErrorResponse(InvalidInput, "Invalid payload"))
		return
	}

	diaryEntry, _ := models.FindDiaryEntry(ctx.Request.Context(), db, entryId)

	if DoesMaskHaveField(updateDiaryEntryRequest.Mask, "title") {
		diaryEntry.Title = null.StringFrom(updateDiaryEntryRequest.Title)
	}

	if DoesMaskHaveField(updateDiaryEntryRequest.Mask, "body") {
		newBody := updateDiaryEntryRequest.Body[:updateDiaryEntryRequest.Offset]
		newBody += updateDiaryEntryRequest.Body

		diaryEntry.Body = null.StringFrom(newBody)

		diaryEntry.Update(ctx.Request.Context(), db, boil.Infer())
	}

	ctx.JSON(200, diaryEntry)
}

func DeleteDiaryEntry(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	entryId, _ := strconv.Atoi(ctx.Params.ByName("diaryEntryId"))

	if entryId == 0 {
		fmt.Errorf("no entry id provided")

		ctx.JSON(400, CreateErrorResponse(InvalidInput, "No entry id provided"))
		return
	}

	diaryEntry, err := models.FindDiaryEntry(ctx.Request.Context(), db, entryId)

	if err != nil {
		fmt.Errorf("fetching diary entry failed %v", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))
		return
	}

	_, err = diaryEntry.Delete(ctx.Request.Context(), db)

	if err != nil {
		fmt.Errorf("deleting diaryentry failed %v", err)
		return
	}

	ctx.JSON(200, struct{}{})
}

func GetDiaryEntries(ctx *gin.Context) {
	// userSession := GetUserSessionFromContext(ctx)

	db := GetDBFromContext(ctx)

	dbEntries, err := models.DiaryEntries(
		// qm.Where("who_posted_user_id = ?", userSession.UserID),
		qm.Limit(30),
		qm.OrderBy("created_at desc"),
	).All(ctx.Request.Context(), db)

	entries := []DiaryEntry{}

	for _, dbEntry := range dbEntries {
		entries = append(entries, DiaryEntry{
			ID:       dbEntry.ID,
			Title:    dbEntry.Title.String,
			Body:     dbEntry.Body.String,
			CreateAt: dbEntry.CreatedAt,
		})
	}

	if err != nil {
		fmt.Errorf("Entries fetch failed %v", err)
	}

	if entries == nil {
		ctx.JSON(200, []struct{}{})
		return
	}

	ctx.JSON(200, entries)
}
