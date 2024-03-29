package routes

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MessageAndUser struct {
	models.Message `boil:",bind"`
	models.User    `boil:",bind"`
}

func GetChatroomMessages(ctx *gin.Context) {
	// chatroomId := ctx.Param("chatroomId")

	chatroomId := ctx.Param("chatroomId")

	chatroomIdNum, err := strconv.Atoi(chatroomId)

	if err != nil {
		fmt.Println("Chatroom ID is not a number", err)
	}

	db := GetDBFromContext(ctx)

	rows := []ChatroomMessage{}

	err = models.NewQuery(
		qm.Select(
			"messages.chatroom_id as chatroom_id",
			"messages.user_id as user_id",
			"users.name as user_name",
			"messages.id as message_id",
			"messages.message_text as message_text",
			"messages.written_at as written_at",
			"messages.transmited_at as transmited_at",
			"messages.reference as reference"),
		qm.OrderBy("written_at desc"),
		qm.Limit(35),
		qm.From("messages"),
		qm.Where("messages.chatroom_id = ?", chatroomIdNum),
		qm.InnerJoin("users on messages.user_id = users.id"),
		// qm.Where("chatroom_id = ?", chatroomId),
	).Bind(ctx.Request.Context(), db, &rows)

	if err != nil {
		fmt.Println("Messages fetch failed", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
	}

	ctx.JSON(200, rows)
}

func GetChatroomMembers(ctx *gin.Context) {
	chatroomId := ctx.Param("chatroomId")

	chatroomIdNum, err := strconv.Atoi(chatroomId)

	if err != nil {
		fmt.Println("Chatroom ID is not a number", err)
	}

	db := GetDBFromContext(ctx)

	rows := []models.User{}

	err = models.NewQuery(
		qm.Select("users.*"),
		qm.From("chatroom_users"),
		qm.Where("chatroom_users.chatroom_id = ?", chatroomIdNum),
		qm.InnerJoin("users on chatroom_users.user_id = users.id"),
	).Bind(ctx.Request.Context(), db, &rows)

	if err != nil {
		fmt.Println("Chatroom members fetch failed", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
	}

	ctx.JSON(200, rows)
}

type CreateChatroomBody struct {
	Name    string   `json:"name"`
	UserIds []uint32 `json:"userIds"`
}

func CreateChatroom(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	session := GetUserSessionFromContext(ctx)

	body := CreateChatroomBody{}

	err := ctx.BindJSON(&body)

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
		return
	}

	chatroom := models.Chatroom{
		Name: null.NewString(body.Name, true),
	}

	err = chatroom.Insert(ctx.Request.Context(), db, boil.Infer())

	if err != nil {
		log.Printf("Creating new chatroom failed %v", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
		return
	}

	processedUserIds := map[uint32]bool{}

	// Add the creator of the chatroom to the chatroom

	processedUserIds[uint32(ctx.GetInt("userId"))] = true

	chatroomUser := models.ChatroomUser{
		ChatroomID: chatroom.ID,
		UserID:     int(session.UserID),
	}

	err = chatroomUser.Insert(ctx.Request.Context(), db, boil.Infer())

	if err != nil {
		log.Printf("Inserting request user to chatroom failed %v", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
		return
	}

	for _, userId := range body.UserIds {
		_, ok := processedUserIds[userId]

		if ok {
			continue
		}

		processedUserIds[userId] = true

		// userIdNum, err := strconv.Atoi(userId)

		// if err != nil {
		// 	ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
		// 	return
		// }

		chatroomUser := models.ChatroomUser{
			ChatroomID: chatroom.ID,
			UserID:     int(userId),
		}

		err = chatroomUser.Insert(ctx.Request.Context(), db, boil.Infer())

		if err != nil {
			log.Printf("inserting user %v to chatroom %v failed %v", userId, chatroom.ID, err)

			ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
			return
		}
	}

	ctx.JSON(200, chatroom)
}

func GetChatroom(ctx *gin.Context) {
	chatroomId := ctx.Param("chatroomId")

	chatroomIdNum, err := strconv.Atoi(chatroomId)

	if err != nil {
		fmt.Println("Chatroom ID is not a number", err)
	}

	db := GetDBFromContext(ctx)

	chatroom, err := models.FindChatroom(ctx.Request.Context(), db, chatroomIdNum)

	if err != nil {
		fmt.Println("Chatroom fetch failed", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
	}

	ctx.JSON(200, chatroom)
}

type PatchChatroomBody struct {
	Name string `json:"name"`
}

func PatchChatroom(ctx *gin.Context) {
	chatroomId := ctx.Param("chatroomId")

	chatroomIdNum, err := strconv.Atoi(chatroomId)

	if err != nil {
		fmt.Println("Chatroom ID is not a number", err)
	}

	db := GetDBFromContext(ctx)

	body := PatchChatroomBody{}

	err = ctx.BindJSON(&body)

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
		return
	}

	chatroom, err := models.FindChatroom(ctx.Request.Context(), db, chatroomIdNum)

	if err != nil {
		fmt.Println("Chatroom fetch failed", err)

		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
	}

	chatroom.Name = null.NewString(body.Name, true)

	_, err = chatroom.Update(ctx.Request.Context(), db, boil.Infer())

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
		return
	}

	ctx.JSON(200, chatroom)
}

type AddChatroomMemberBody struct {
	UserId uint32 `json:"userId"`
}

func AddChatroomMember(ctx *gin.Context) {
	chatroomId := ctx.Param("chatroomId")

	chatroomIdNum, err := strconv.Atoi(chatroomId)

	if err != nil {
		fmt.Println("Chatroom ID is not a number", err)
	}

	db := GetDBFromContext(ctx)

	body := AddChatroomMemberBody{}

	err = ctx.BindJSON(&body)

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
		return
	}

	//Check if user exists
	_, err = models.FindUser(ctx.Request.Context(), db, int(body.UserId))

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
		return
	}

	chatroomUser := models.ChatroomUser{
		ChatroomID: chatroomIdNum,
		UserID:     int(body.UserId),
	}

	err = chatroomUser.Insert(ctx.Request.Context(), db, boil.Infer())

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
		return
	}

	user, err := models.FindUser(ctx.Request.Context(), db, int(body.UserId))

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
		return
	}

	ctx.JSON(200, user)
}

func RemoveChatroomMember(ctx *gin.Context) {
	chatroomId := ctx.Param("chatroomId")

	chatroomIdNum, err := strconv.Atoi(chatroomId)

	if err != nil {
		fmt.Println("Chatroom ID is not a number", err)
	}

	db := GetDBFromContext(ctx)

	userId := ctx.Param("userId")

	userIdNum, err := strconv.Atoi(userId)

	if err != nil {
		fmt.Println("User ID is not a number", err)
	}

	chatroomUser, err := models.ChatroomUsers(
		qm.Where("chatroom_id = ?", chatroomIdNum),
		qm.And("user_id = ?", userIdNum),
	).One(ctx.Request.Context(), db)

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
		return
	}

	_, err = chatroomUser.Delete(ctx.Request.Context(), db)

	if err != nil {
		ctx.JSON(500, CreateErrorResponse(InternalServerError, "Internal server error"))
		return
	}

	ctx.JSON(200, gin.H{})
}
