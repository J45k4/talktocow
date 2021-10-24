package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

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
	entryId := ctx.Params.ByName("diaryEntryId")

	fmt.Printf("entryId %v", entryId)
}

func UpdateDiaryEntry(ctx *gin.Context) {

}

func DeleteDiaryEntry(ctx *gin.Context) {

}

func GetDiaryEntries(ctx *gin.Context) {
	userSession := GetUserSessionFromContext(ctx)

	db := GetDBFromContext(ctx)

	entries, err := models.DiaryEntries(
		qm.Where("who_posted_user_id = ?", userSession.UserID),
		qm.Limit(30),
		qm.OrderBy("created_at desc"),
	).All(ctx.Request.Context(), db)

	if err != nil {
		fmt.Errorf("Entries fetch failed %v", err)
	}

	ctx.JSON(200, entries)
}
