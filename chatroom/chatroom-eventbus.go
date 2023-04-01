package chatroom

import (
	"fmt"
	"log"

	"github.com/asaskevich/EventBus"
	"github.com/gin-gonic/gin"
)

type ChatroomMessage struct {
	ChatroomID       string `json:"chatroomId"`
	MessageID        string `json:"messageId"`
	UserID           string `json:"userId"`
	UserName         string `json:"userName"`
	MessageText      string `json:"messageText"`
	WrittenAt        string `json:"writtenAt"`
	TransmitedAt     string `json:"transmitedAt"`
	ServerReceivedAt string
	Platform         string `json:"platform"`
	Reference        string
}

type ChatroomMessageCallback = func(ChatroomMessage)

type ChatroomEventbus struct {
	eventbus EventBus.Bus
}

func NewChatroomEventbus() *ChatroomEventbus {
	bus := EventBus.New()

	return &ChatroomEventbus{
		eventbus: bus,
	}
}

func (this *ChatroomEventbus) SendChatroomMessage(chatroomID int, chatroomMessage ChatroomMessage) {
	log.Println("Send chatroomMessage", chatroomID, chatroomMessage)

	this.eventbus.Publish(fmt.Sprintf("chatroomMessage:%d", chatroomID), chatroomMessage)
}

func (this *ChatroomEventbus) SubscribeToChatroomMessages(chatroomID int, cb ChatroomMessageCallback) {
	this.eventbus.Subscribe(fmt.Sprintf("chatroomMessage:%d", chatroomID), cb)
}

func (this *ChatroomEventbus) UnsubscribeToChatroomMessages(chatroomID int, cb ChatroomMessageCallback) {
	err := this.eventbus.Unsubscribe(fmt.Sprintf("chatroomMessage:%d", chatroomID), cb)

	if err != nil {
		log.Println("Unsibscribe error", err)
	}
}

func GetChatroomEventbus(ctx *gin.Context) *ChatroomEventbus {
	e, _ := ctx.Get("chatroomEventbus")

	return e.(*ChatroomEventbus)
}

func SetChatroomEventbus(ctx *gin.Context, chatroomEventbus *ChatroomEventbus) {
	ctx.Set("chatroomEventbus", chatroomEventbus)
}
