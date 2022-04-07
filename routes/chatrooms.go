package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetChatrooms(ctx *gin.Context) {
	userSession := GetUserSessionFromContext(ctx)

	db := GetDBFromContext(ctx)

	chatrooms := []models.Chatroom{}

	err := models.NewQuery(
		qm.InnerJoin("chatroom_users on chatroom_users.chatroom_id = chatrooms.id"),
		qm.Where("chatroom_users.user_id = ?", userSession.UserID),
		qm.From("chatrooms"),
	).Bind(ctx.Request.Context(), db, &chatrooms)

	if err != nil {
		fmt.Println("Chatrooms fetch failed", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
		return
	}

	ctx.JSON(200, chatrooms)
}
