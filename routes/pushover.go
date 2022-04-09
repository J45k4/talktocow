package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type CreatePushoverTokenBody struct {
	UserToken string `json:"userToken"`
	Token     string `json:"token"`
	UserID    int    `json:"userId"`
}

func CreatePushoverToken(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	var body CreatePushoverTokenBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pushoverToken := models.PushoverToken{
		UserToken: body.UserToken,
		Token:     body.Token,
		UserID:    body.UserID,
	}

	err := pushoverToken.Insert(ctx.Request.Context(), db, boil.Infer())

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pushoverToken)
}

type UpdatePushoverTokenBody struct {
	UserToken string `json:"userToken"`
	Token     string `json:"token"`
	UserID    int    `json:"userId"`
}

func UpdatePushoverToken(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	var body UpdatePushoverTokenBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pushoverToken := models.PushoverToken{
		UserToken: body.UserToken,
		Token:     body.Token,
		UserID:    body.UserID,
	}

	_, err := pushoverToken.Update(ctx.Request.Context(), db, boil.Infer())

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pushoverToken)
}

func DeletePushoverToken(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	pushoverTokenID := ctx.Params.ByName("pushoverTokenId")

	_, err := models.PushoverTokens(qm.Where("id = ?", pushoverTokenID)).DeleteAll(ctx.Request.Context(), db)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func GetPushoverTokens(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	pushoverTokens, err := models.PushoverTokens().All(ctx.Request.Context(), db)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pushoverTokens)
}

func GetPushoverToken(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	pushoverTokenID := ctx.Params.ByName("pushoverTokenId")

	pushoverToken, err := models.PushoverTokens(qm.Where("id = ?", pushoverTokenID)).One(ctx.Request.Context(), db)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pushoverToken)
}
