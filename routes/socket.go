package routes

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/j45k4/talktocow/chatroom"
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
	TransmitedAt string `json:"transmitedAt"`
	Reference    string `json:"reference"`
}

type WebsocketMessageToServer struct {
	MessageToChatroom *MessageToChatroom `json:"messageToChatroom"`
}

type UserStatus struct {
	Online    bool   `json:"online"`
	UserID    string `json:"userId"`
	Username  string `json:"username"`
	Lastseen  string `json:"lastseen"`
	Sleeping  bool   `json:"sleeping"`
	Timestamp string `json:"timestamp"`
}

type NewChatroomMessage struct {
	ChatroomID       string `json:"chatroomId"`
	UserID           string `json:"userId"`
	UserName         string `json:"userName"`
	MessageID        string `json:"messageId"`
	MessageText      string `json:"messageText"`
	WritenAt         string `json:"writenAt"`
	TransmitedAt     string `json:"transmitedAt"`
	ServerReceivedAt string `json:"serverReceivedAt"`
	Platform         string `json:"platform"`
	Reference        string `json:"reference"`
}

type WebsocketMessageToClient struct {
	ChangedUserStatus  *UserStatus         `json:"changedUserStatus"`
	NewChatroomMessage *NewChatroomMessage `json:"newChatroomMessage"`
}

func processMessageRead(
	ws *websocket.Conn,
	db *sql.DB,
	chatroomEventbus *chatroom.ChatroomEventbus,
	userSession UserSession,
) {
	ctx := context.Background()

	for {
		var msg WebsocketMessageToServer

		err := ws.ReadJSON(&msg)

		if err != nil {
			return
		}

		if msg.MessageToChatroom != nil {
			log.Println("Received new chatroom message %v", msg.MessageToChatroom)

			chatroomID, _ := strconv.Atoi(msg.MessageToChatroom.ChatroomID)

			transmittedAt, _ := time.Parse(time.RFC3339Nano, msg.MessageToChatroom.TransmitedAt)
			writenAt, _ := time.Parse(time.RFC3339Nano, msg.MessageToChatroom.WritenAt)
			//createdAt, _ := time.Parse(time.RFC3339Nano, msg.MessageToChatRoom.CreateTime)

			serverReceivedAt := time.Now()

			platform := "talktocow"

			newMessage := models.Message{
				MessageText:      null.NewString(msg.MessageToChatroom.MessageText, true),
				ServerReceivedAt: serverReceivedAt,
				UserID:           int(userSession.UserID),
				Platform:         null.StringFrom(platform),
				ChatroomID:       1,
				TransmitedAt:     transmittedAt,
				WrittenAt:        writenAt,
				Reference:        null.StringFrom(msg.MessageToChatroom.Reference),
			}

			messageInserErr := newMessage.Insert(ctx, db, boil.Infer())

			if messageInserErr != nil {
				fmt.Println("Message insert failed", messageInserErr)
			}

			newChatroomEvent := models.ChatroomEvent{
				ChatroomID: chatroomID,
				MessageID:  null.IntFrom(newMessage.ID),
				EventType:  1,
			}

			newChatroomEvent.Insert(ctx, db, boil.Infer())

			chatroomEventbus.SendChatroomMessage(chatroomID, chatroom.ChatroomMessage{
				MessageID:        fmt.Sprint(newMessage.ID),
				ChatroomID:       fmt.Sprint(chatroomID),
				UserID:           fmt.Sprint(userSession.UserID),
				UserName:         userSession.UserName,
				MessageText:      msg.MessageToChatroom.MessageText,
				WritenAt:         msg.MessageToChatroom.WritenAt,
				TransmitedAt:     msg.MessageToChatroom.TransmitedAt,
				ServerReceivedAt: serverReceivedAt.Format(time.RFC3339),
				Platform:         platform,
				Reference:        msg.MessageToChatroom.Reference,
			})

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
	chatroomEventbus := chatroom.GetChatroomEventbus(ctx)

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

	log.Printf("New socket connection from {%s}", userSession.UserName)

	ctx.Request.Context()

	go processMessageRead(ws, db, chatroomEventbus, userSession)

	var handle chatroom.ChatroomMessageCallback

	handle = func(chatroomMessage chatroom.ChatroomMessage) {
		log.Println("Chatroom subscriber received message", chatroomMessage)

		messageToClient := WebsocketMessageToClient{
			NewChatroomMessage: &NewChatroomMessage{
				ChatroomID:       chatroomMessage.ChatroomID,
				UserID:           chatroomMessage.UserID,
				UserName:         chatroomMessage.UserName,
				MessageID:        chatroomMessage.MessageID,
				MessageText:      chatroomMessage.MessageText,
				WritenAt:         chatroomMessage.WritenAt,
				TransmitedAt:     chatroomMessage.TransmitedAt,
				ServerReceivedAt: chatroomMessage.ServerReceivedAt,
				Platform:         chatroomMessage.Platform,
				Reference:        chatroomMessage.Reference,
			},
		}

		err := ws.WriteJSON(messageToClient)

		if err != nil {
			log.Println("Sending message to websocket channel failed", err)

			// chatroomEventbus.UnsubscribeToChatroomMessages(1, handle)
			//ctx.Abort()
		}
	}

	log.Println("Subscribing to channel")

	chatroomEventbus.SubscribeToChatroomMessages(1, handle)

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
