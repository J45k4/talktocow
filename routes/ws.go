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
	"github.com/j45k4/talktocow/auth"
	"github.com/j45k4/talktocow/eventbus"
	"github.com/j45k4/talktocow/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type WsMessageFromClientType string

const (
	WsSendMessage     WsMessageFromClientType = "sendMessage"
	WsAuthenticate    WsMessageFromClientType = "authenticate"
	WsAskQuestion     WsMessageFromClientType = "askQuestion"
	WsGetQuestions    WsMessageFromClientType = "getQuestions"
	WsMakeVideoOffer  WsMessageFromClientType = "makeVideoOffer"
	WsMakeVideoAnswer WsMessageFromClientType = "makeVideoAnswer"
	WsNewIceCandidate WsMessageFromClientType = "newIceCandidate"
)

type WsMsgFromClient struct {
	Type         WsMessageFromClientType `json:"type"`
	Token        *string                 `json:"token"`
	MessageText  *string                 `json:"messageText"`
	ChatroomID   *string                 `json:"chatroomID"`
	UserID       *string                 `json:"userID"`
	UserName     *string                 `json:"userName"`
	WrittenAt    *string                 `json:"writtenAt"`
	TransmitedAt *string                 `json:"transmitedAt"`
	Reference    *string                 `json:"reference"`
	SDP          *string                 `json:"sdp"`
	Candidate    *string                 `json:"candidate"`
}

type Authenticated struct {
	Type string `json:"type"`
}

type WsMessageToClientType string

const (
	WsAuthenticated        WsMessageToClientType = "authenticated"
	WsUnauthenticated      WsMessageToClientType = "unauthenticated"
	WsChatroomMessagesType WsMessageToClientType = "chatroomMessages"
)

type ChatroomMessages struct {
	Type     WsMessageToClientType `json:"type"`
	Messages []ChatroomMessage     `json:"messages"`
}

type WsMsgToClient struct {
	Type      WsMessageToClientType `json:"type"`
	Questions *[]string             `json:"questions"`
}

func wsChan(ws *websocket.Conn) chan WsMsgFromClient {
	c := make(chan WsMsgFromClient, 200)

	go func() {
		defer close(c)
		for {
			var msg WsMsgFromClient
			err := ws.ReadJSON(&msg)

			fmt.Printf("received %v", msg.Type)

			if err != nil {
				log.Println("error:", err)
				break
			}

			c <- msg

			fmt.Println("value sent to channel")
		}
	}()

	return c
}

type WsHandler struct {
	ws       *websocket.Conn
	db       *sql.DB
	eventbus *eventbus.Eventbus
	ctx      context.Context
	userID   int32
	userName string
}

func (h *WsHandler) sendMessage(msg interface{}) {
	h.ws.WriteJSON(msg)
}

func (h *WsHandler) handleAuthenticate(msg WsMsgFromClient) bool {
	var userSession UserSession

	err := auth.DecodeObjectFromToken(*msg.Token, &userSession)

	if err != nil {
		h.sendMessage(WsMsgToClient{
			Type: WsUnauthenticated,
		})

		return true
	}

	fmt.Printf("user %v authenticated", userSession.UserID)

	h.userID = userSession.UserID
	h.userName = userSession.UserName

	h.sendMessage(WsMsgToClient{
		Type: WsAuthenticated,
	})

	return false
}

func (h *WsHandler) handleSendMessage(msg WsMsgFromClient) bool {
	fmt.Println("handleSendMessage", msg)

	chatroomID, err := strconv.Atoi(*msg.ChatroomID)

	if err != nil {
		fmt.Println("error converting chatroomID to int: ", err)
		return false
	}

	chatroom, err := models.FindChatroom(h.ctx, h.db, chatroomID)

	if err != nil {
		fmt.Println("error finding chatroom: ", err)
		return false
	}

	if chatroom == nil {
		fmt.Println("chatroom not found")
		return false
	}

	var writtenAtTime time.Time

	if msg.WrittenAt != nil {
		writtenAtTime, _ = time.Parse(time.RFC3339, *msg.WrittenAt)
	}

	var transmitedAtTime time.Time

	if msg.TransmitedAt != nil {
		transmitedAtTime, _ = time.Parse(time.RFC3339, *msg.TransmitedAt)
	}

	reference := null.NewString("", false)

	if msg.Reference != nil {
		reference = null.NewString(*msg.Reference, true)
	}

	newMessage := models.Message{
		MessageText:      null.NewString(*msg.MessageText, true),
		ChatroomID:       chatroomID,
		UserID:           int(h.userID),
		WrittenAt:        writtenAtTime,
		TransmitedAt:     transmitedAtTime,
		Reference:        reference,
		CreatedAt:        time.Now(),
		ServerReceivedAt: time.Now(),
	}

	err = newMessage.Insert(h.ctx, h.db, boil.Infer())

	if err != nil {
		fmt.Println("error inserting new message: ", err)
		return false
	}

	log.Printf("ws publish message")

	h.eventbus.Publish(eventbus.Event{
		ChatroomMessage: &eventbus.ChatroomMessage{
			ChatroomID:   chatroomID,
			ID:           newMessage.ID,
			MessageText:  *msg.MessageText,
			UserID:       int(h.userID),
			WrittenAt:    writtenAtTime,
			TransmitedAt: transmitedAtTime,
			Reference:    *msg.Reference,
		},
	})

	log.Printf("ws message sent message handled")

	return false
}

func (h *WsHandler) handleAskQuestion(msg WsMsgFromClient) bool {
	newChatroom := models.Chatroom{
		Name: null.NewString("new question", true),
	}

	fmt.Printf("handleAskQuestion: %v\n", newChatroom.Name)

	err := newChatroom.Insert(h.ctx, h.db, boil.Infer())

	if err != nil {
		fmt.Println("error inserting new question: ", err)
		return false
	}

	fmt.Printf("giving access to chatroom %v to user %v\n", newChatroom.ID, h.userID)

	chatroomUser := models.ChatroomUser{
		ChatroomID: newChatroom.ID,
		UserID:     int(h.userID),
	}

	err = chatroomUser.Insert(h.ctx, h.db, boil.Infer())

	if err != nil {
		fmt.Println("error inserting chatroomUser: ", err)
	}

	return false
}

func (h *WsHandler) handleMakeVideoOffer(msg WsMsgFromClient) bool {
	h.eventbus.Publish(eventbus.Event{
		VideoOffer: &eventbus.VideoOfferEvent{
			SDP:    *msg.SDP,
			UserID: int(h.userID),
		},
	})

	return false
}

func (h *WsHandler) handleMakeVideoAnswer(msg WsMsgFromClient) bool {
	h.eventbus.Publish(eventbus.Event{
		VideoAnswer: &eventbus.VideoAnswerEvent{
			SDP:    *msg.SDP,
			UserID: int(h.userID),
		},
	})

	return false
}

func (h *WsHandler) handleNewIceCandidate(msg WsMsgFromClient) bool {
	h.eventbus.Publish(eventbus.Event{
		NewIceCandidate: &eventbus.NewIceCandidateEvent{
			Candidate: *msg.Candidate,
			UserID:    int(h.userID),
		},
	})

	return false
}

func (h *WsHandler) handleWsMsg(msg WsMsgFromClient) bool {
	switch msg.Type {
	case WsAuthenticate:
		return h.handleAuthenticate(msg)
	}

	if h.userID == 0 {
		return true
	}

	switch msg.Type {
	case WsSendMessage:
		return h.handleSendMessage(msg)
	case WsAskQuestion:
		return h.handleAskQuestion(msg)
	case WsMakeVideoOffer:
		return h.handleMakeVideoOffer(msg)
	case WsMakeVideoAnswer:
		return h.handleMakeVideoAnswer(msg)
	case WsNewIceCandidate:
		return h.handleNewIceCandidate(msg)
	default:
		return false
	}
}

func (h *WsHandler) handleEvent(event eventbus.Event) {
	fmt.Println("ws handler got event: ", event)

	if event.ChatroomMessage != nil {
		chatroomMessage := ChatroomMessage{
			ChatroomID:   strconv.Itoa(event.ChatroomMessage.ChatroomID),
			MessageText:  event.ChatroomMessage.MessageText,
			UserID:       strconv.Itoa(event.ChatroomMessage.UserID),
			WrittenAt:    event.ChatroomMessage.WrittenAt.Format(time.RFC3339),
			TransmitedAt: event.ChatroomMessage.TransmitedAt.Format(time.RFC3339),
			Reference:    event.ChatroomMessage.Reference,
		}

		chatroomMessages := ChatroomMessages{
			Type:     WsChatroomMessagesType,
			Messages: []ChatroomMessage{chatroomMessage},
		}

		h.sendMessage(chatroomMessages)
	}

	if event.VideoAnswer != nil {
		if event.VideoAnswer.UserID != int(h.userID) {
			return
		}

		h.sendMessage(
			VideoAnswer{
				SDP:    event.VideoAnswer.SDP,
				UserID: strconv.Itoa(event.VideoAnswer.UserID),
			},
		)
	}

	if event.VideoOffer != nil {
		if event.VideoOffer.UserID != int(h.userID) {
			return
		}

		h.sendMessage(
			VideoOffer{
				SDP:    event.VideoOffer.SDP,
				UserID: strconv.Itoa(event.VideoOffer.UserID),
			},
		)
	}
}

func (h *WsHandler) Run() {
	fmt.Println("running ws handler")

	readerChan := wsChan(h.ws)
	eventChan := h.eventbus.Subscribe()

	fmt.Println("starting event loop")

	for {
		log.Printf("ws waiting for next")

		select {
		case msg, ok := <-readerChan:
			if !ok {
				fmt.Println("readerChan closed")

				return
			}

			fmt.Println("received wsMsg: ", msg)

			if h.handleWsMsg(msg) {
				log.Printf("ws handler exiting")

				return
			}
		case event := <-eventChan:
			h.handleEvent(event)
		}
	}

	fmt.Println("exiting ws handler")
}

func HandleWs(ctx *gin.Context) {
	db := GetDBFromContext(ctx)
	eventbus := GetEventbusFromContext(ctx)

	w := ctx.Writer
	r := ctx.Request

	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)

	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not websocket handshake", 400)
	} else if err != nil {
		return
	}

	wsHandler := WsHandler{
		ws:       ws,
		db:       db,
		eventbus: eventbus,
		ctx:      ctx.Request.Context(),
	}

	wsHandler.Run()
}
