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

func GetMyChatrooms(
	ctx *gin.Context,
) {
	userSession := GetUserSessionFromContext(ctx)

	db := GetDBFromContext(ctx)

	// chatrooms := []models.Chatroom{}

	// Find all distinct chatrooms where the user is a member and order them by the last message
	chatrooms, err := models.Chatrooms(
		qm.InnerJoin("chatroom_users on chatroom_users.chatroom_id = chatrooms.id"),
		qm.Where("chatroom_users.user_id = ?", userSession.UserID),
		qm.InnerJoin("(SELECT chatroom_id, MAX(created_at) AS latest_message FROM messages GROUP BY chatroom_id) AS latest_messages ON latest_messages.chatroom_id = chatrooms.id"),
		qm.InnerJoin("messages ON messages.chatroom_id = chatrooms.id AND messages.created_at = latest_messages.latest_message"),
		qm.OrderBy("latest_message DESC"),
	).All(ctx, db)

	if err != nil {
		fmt.Println("Chatrooms fetch failed", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
		return
	}

	ctx.JSON(200, chatrooms)
}
