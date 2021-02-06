package routes

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/j45k4/talktocow/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ChatroomEventType uint32

const (
	ChatroomEventTypeMessage ChatroomEventType = 1
)

type UserSession struct {
	UserID   int32  `json:"userId"`
	UserName string `json:"userName"`
}

type MessageToChatroom struct {
	MessageText  string `json:"messageText"`
	ChatroomID   string `json:"chatroomId"`
	WritenAt     string `json:"writenAt"`
	TransmitedAt string `json:"trasmitedAt"`
}

type WebsocketMessageToServer struct {
	MessageToChatroom *MessageToChatroom `json:"messageToChatroom"`
}

type Message struct {
	ID           string `json:"id"`
	MessageText  string `json:"messageText"`
	WritenAt     string `json:"writenAt"`
	TransmitedAt string `json:"transmitedAt"`
	Platform     string `json:"platform"`
	ChatroomID   string `json:"chatroomId"`
}

type ChatroomEvent struct {
	ID         string  `json:"id"`
	ChatroomID string  `json:"chatroomId"`
	Message    Message `json:"message"`
	CreatedAt  string  `json:"createdAt"`
}

type UserStatus struct {
	Online    bool   `json:"online"`
	UserID    string `json:"userId"`
	Username  string `json:"username"`
	Lastseen  string `json:"lastseen"`
	Sleeping  bool   `json:"sleeping"`
	Timestamp string `json:"timestamp"`
}

type WebsocketMessageToClient struct {
	ChangedUserStatus UserStatus `json:"changedUserStatus"`
}

func processMessageRead(ws *websocket.Conn, db *sql.DB, userSession UserSession) {
	ctx := context.Background()

	for {
		var msg WebsocketMessageToServer

		ws.ReadJSON(&msg)

		if msg.MessageToChatroom != nil {
			chatroomId, _ := strconv.Atoi(msg.MessageToChatroom.ChatroomID)

			transmittedAt, _ := time.Parse(time.RFC3339Nano, msg.MessageToChatroom.TransmitedAt)
			writenAt, _ := time.Parse(time.RFC3339Nano, msg.MessageToChatroom.WritenAt)
			//createdAt, _ := time.Parse(time.RFC3339Nano, msg.MessageToChatRoom.CreateTime)

			newMessage := models.Message{
				MessageText:      null.NewString(msg.MessageToChatroom.MessageText, true),
				ServerReceivedAt: time.Now(),
				UserID:           int(userSession.UserID),
				Platform:         null.StringFrom("talktocow"),
				ChatroomID:       1,
				TransmitedAt:     transmittedAt,
				WrittenAt:        writenAt,
			}

			messageInserErr := newMessage.Insert(ctx, db, boil.Infer())

			if messageInserErr != nil {
				fmt.Println("Message insert failed", messageInserErr)
			}

			newChatroomEvent := models.ChatroomEvent{
				ChatroomID: chatroomId,
				MessageID:  null.IntFrom(newMessage.ID),
				EventType:  1,
			}

			newChatroomEvent.Insert(ctx, db, boil.Infer())

			// newChatroomMessage := NewChatroomMessage{
			// 	MessageText:   msg.MessageToChatRoom.MessageText,
			// 	FromUserName:  userSession.Name,
			// 	TransmittedAt: msg.MessageToChatRoom.TransmitTime,
			// }

			// transmitMessage := WebsocketTransmitMessage{
			// 	NewChatroomMessage: &newChatroomMessage,
			// }

			// fmt.Printf("Sending message %v to %v", transmitMessage, userSession)

			// for c, _ := range connections {
			// 	c.WriteJSON(transmitMessage)
			// }
		}
	}
}

func HandleSocket(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	w := ctx.Writer
	r := ctx.Request

	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)

	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not websocket handshake", 400)
	} else if err != nil {
		return
	}

	userSession := GetUserSessionFromContext(ctx)

	fmt.Printf("New socket connection from {%s}", userSession.UserName)

	ctx.Request.Context()

	go processMessageRead(ws, db, userSession)

	// connections[ws] = true

	// go func() {
	// 	for {
	// 		msg := WebsocketReceiveMessage{}

	// 		err = ws.ReadJSON(&msg)

	// 		if err != nil {
	// 			delete(connections, ws)

	// 			return
	// 		}

	// 		fmt.Printf("new message %v\n", msg)

	// 		if msg.MessageToChatRoom != nil {
	// 			transmittedAt, _ := time.Parse(time.RFC3339Nano, msg.MessageToChatRoom.TransmitTime)
	// 			//createdAt, _ := time.Parse(time.RFC3339Nano, msg.MessageToChatRoom.CreateTime)

	// 			newMessage := models.Message{
	// 				MessageText:      null.NewString(msg.MessageToChatRoom.MessageText, true),
	// 				ServerReceivedAt: time.Now(),
	// 				UserID:           int(userSession.UserID),
	// 				Platform:         null.StringFrom("talktocow"),
	// 				ChatroomID:       1,
	// 				TransmitedAt:     transmittedAt,
	// 			}

	// 			messageInserErr := newMessage.Insert(context.Background(), db, boil.Infer())

	// 			if messageInserErr != nil {
	// 				fmt.Println("Message insert failed", messageInserErr)
	// 			}

	// 			newChatroomMessage := NewChatroomMessage{
	// 				MessageText:   msg.MessageToChatRoom.MessageText,
	// 				FromUserName:  userSession.Name,
	// 				TransmittedAt: msg.MessageToChatRoom.TransmitTime,
	// 			}

	// 			transmitMessage := WebsocketTransmitMessage{
	// 				NewChatroomMessage: &newChatroomMessage,
	// 			}

	// 			fmt.Printf("Sending message %v to %v", transmitMessage, userSession)

	// 			for c, _ := range connections {
	// 				c.WriteJSON(transmitMessage)
	// 			}
	// 		}

	// 		// for c, _ := range connections {
	// 		// 	c.WriteJSON(WebsocketSendMessage{
	// 		// 		MessageText: msg.MessageText,
	// 		// 	})
	// 		// }
	// 	}
	// }()
}
