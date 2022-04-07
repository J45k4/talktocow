package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetUsers(ctx *gin.Context) {
	limitStr, _ := ctx.GetQuery("limit")

	limit := 10

	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}

	db := GetDBFromContext(ctx)
	users := []models.User{}

	err := models.Users(
		qm.Limit(limit),
	).Bind(ctx.Request.Context(), db, &users)

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
		return
	}

	ctx.JSON(200, users)
}
