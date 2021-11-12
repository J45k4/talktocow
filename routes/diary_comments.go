package routes

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type DiaryEntryComment struct {
	ID           int    `json:"id",boil:"id"`
	DiaryEntryID int    `json:"diaryEntryId",boil:"diary_entry_id"`
	UserID       int    `json:"userId",boil:"user_id"`
	UserName     string `json:"userName",boil:"user_name"`
	CommentText  string `json:"commentText",boil:"comment_text"`
	CreatedAt    string `json:"createdAt",boil:"created_at"`
	UpdatedAt    string `json:"updatedAt",boil:"updated_at"`
}

func DiaryEntryComments(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

type CreateDiaryEntryCommentRequest struct {
	CommentText string `json:"commentText"`
}

func CreateDiaryEntryComment(ctx *gin.Context) {
	userSession := GetUserSessionFromContext(ctx)

	db := GetDBFromContext(ctx)

	diaryEntryIdStr := ctx.Param("diaryEntryId")

	diaryEntryId, err := strconv.Atoi(diaryEntryIdStr)

	if err != nil {
		ctx.JSON(400, CreateErrorResponse(InvalidInput, "Invalid payload"))
		return
	}

	var request CreateDiaryEntryCommentRequest

	err = ctx.BindJSON(&request)

	if err != nil {
		ctx.JSON(400, CreateErrorResponse(InvalidInput, "Invalid payload"))
		return
	}

	diaryEntryComment := models.DiaryEntryComment{
		DiaryEntryID: diaryEntryId,
		UserID:       int(userSession.UserID),
		CreatedAt:    time.Now(),
		CommentText:  request.CommentText,
	}

	err = diaryEntryComment.Insert(ctx.Request.Context(), db, boil.Infer())

	if err != nil {
		log.Printf("Error inserting diary entry comment: %s", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Failed to create diary entry comment"))
		return
	}

	ctx.JSON(200, diaryEntryComment)
}

type UpdateDiaryEntryCommentRequest struct {
	CommentText string `json:"comment"`
}

func UpdateDiaryEntryComment(ctx *gin.Context) {
	userSession := GetUserSessionFromContext(ctx)

	db := GetDBFromContext(ctx)

	diaryEntryCommentIdStr := ctx.Param("commentId")

	diaryEntryCommentId, err := strconv.Atoi(diaryEntryCommentIdStr)

	if err != nil {
		ctx.JSON(400, CreateErrorResponse(InvalidInput, "Invalid payload"))
		return
	}

	var request UpdateDiaryEntryCommentRequest

	err = ctx.BindJSON(&request)

	if err != nil {
		ctx.JSON(400, CreateErrorResponse(InvalidInput, "Invalid payload"))
		return
	}

	diaryEntryComment, err := models.FindDiaryEntryComment(ctx.Request.Context(), db, diaryEntryCommentId)

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))
		return
	}

	if diaryEntryComment.UserID != int(userSession.UserID) {
		ctx.JSON(403, CreateErrorResponse(AuthorizationError, "You are not allowed to edit this diary entry comment"))
		return
	}

	diaryEntryComment.CommentText = request.CommentText

	diaryEntryComment.Update(ctx.Request.Context(), db, boil.Infer())

	ctx.JSON(200, diaryEntryComment)
}

type RemoveDiaryEntryCommentRequest struct {
	CommendID int `json:"commentId"`
}

func DeleteDiaryEntryComment(ctx *gin.Context) {
	userSession := GetUserSessionFromContext(ctx)

	db := GetDBFromContext(ctx)

	diaryEntryCommentIdStr := ctx.Param("commentId")

	diaryEntryCommentId, err := strconv.Atoi(diaryEntryCommentIdStr)

	if err != nil {
		ctx.JSON(400, CreateErrorResponse(InvalidInput, "Invalid payload"))
		return
	}

	diaryEntryComment, err := models.FindDiaryEntryComment(ctx.Request.Context(), db, diaryEntryCommentId)

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))
		return
	}

	if diaryEntryComment.UserID != int(userSession.UserID) {
		ctx.JSON(403, CreateErrorResponse(AuthorizationError, "You are not allowed to delete this diary entry comment"))
		return
	}

	_, err = diaryEntryComment.Delete(ctx.Request.Context(), db)

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))
		return
	}

	ctx.JSON(200, gin.H{})
}

type DiaryEntryCommentAndUser struct {
	models.DiaryEntryComment `boil:",bind"`
	models.User              `boil:",bind"`
}

func GetDiaryEntryComments(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	diaryEntryIDStr := ctx.Param("diaryEntryId")

	diaryEntryID, err := strconv.Atoi(diaryEntryIDStr)

	if err != nil {
		ctx.JSON(400, CreateErrorResponse(InvalidInput, "Invalid payload"))
		return
	}

	offset, limit := GetOffsetAndLimit(ctx, 0, 5)

	rows := []DiaryEntryComment{}

	err = models.NewQuery(
		qm.Select(
			"d.id as id",
			"d.diary_entry_id as diary_entry_id",
			"d.user_id as user_id",
			"u.name as user_name",
			"d.comment_text as comment_text",
			"d.created_at as created_at",
			"d.updated_at as updated_at",
		),
		qm.Where("diary_entry_id = ?", diaryEntryID),
		qm.Offset(offset),
		qm.Limit(limit),
		qm.OrderBy("d.created_at asc"),
		qm.InnerJoin("users u on u.id = d.user_id"),
		qm.From("diary_entry_comments as d"),
	).Bind(ctx.Request.Context(), db, &rows)

	if err != nil {
		log.Printf("Failed to get diary entry comments: %v", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))
		return
	}

	ctx.JSON(200, rows)
}

func GetDiaryEntryCommentsCount(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	diaryEntryIDStr := ctx.Param("diaryEntryId")

	diaryEntryID, err := strconv.Atoi(diaryEntryIDStr)

	if err != nil {
		ctx.JSON(400, CreateErrorResponse(InvalidInput, "Invalid payload"))
		return
	}

	count, err := models.DiaryEntryComments(
		qm.Where("diary_entry_id = ?", diaryEntryID),
	).Count(ctx.Request.Context(), db)

	if err != nil {
		log.Printf("Failed to get diary entry comments count: %v", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, ""))
		return
	}

	ctx.JSON(200, gin.H{
		"count": count,
	})
}
