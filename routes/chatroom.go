package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ChatroomMessage struct {
	ChatroomID   string      `json:"chatroomId" boil:"chatroom_id"`
	UserID       string      `json:"userId" boil:"user_id"`
	UserName     string      `json:"userName" boil:"user_name"`
	MessageID    string      `json:"messageId" boil:"message_id"`
	MessageText  string      `json:"messageText" boil:"message_text"`
	WritenAt     string      `json:"writenAt" boil:"writen_at"`
	TransmitedAt string      `json:"transmitedAt" boil:"transmited_at"`
	Platform     string      `json:"platform" boil:"platform"`
	Reference    null.String `json:"reference" boil:"reference"`
}

type MessageAndUser struct {
	models.Message `boil:",bind"`
	models.User    `boil:",bind"`
}

func GetChatroomMessages(ctx *gin.Context) {
	// chatroomId := ctx.Param("chatroomId")

	db := GetDBFromContext(ctx)

	rows := []ChatroomMessage{}

	err := models.NewQuery(
		qm.Select(
			"messages.chatroom_id as chatroom_id",
			"messages.user_id as user_id",
			"users.name as user_name",
			"messages.id as message_id",
			"messages.message_text as message_text",
			"messages.written_at as writen_at",
			"messages.transmited_at as transmited_at",
			"messages.platform as platform",
			"messages.reference as reference"),
		qm.OrderBy("transmited_at desc"),
		qm.Limit(35),
		qm.From("messages"),
		qm.InnerJoin("users on messages.user_id = users.id"),
		// qm.Where("chatroom_id = ?", chatroomId),
	).Bind(ctx.Request.Context(), db, &rows)

	if err != nil {
		fmt.Println("Messages fetch failed", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
	}

	ctx.JSON(200, rows)
}
