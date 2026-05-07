package chatroom

import (
	"fmt"
	"log"
	"reflect"
	"sync"

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
	lock        sync.RWMutex
	subscribers map[string][]ChatroomMessageCallback
}

func NewChatroomEventbus() *ChatroomEventbus {
	return &ChatroomEventbus{
		subscribers: make(map[string][]ChatroomMessageCallback),
	}
}

func (this *ChatroomEventbus) SendChatroomMessage(chatroomID int, chatroomMessage ChatroomMessage) {
	log.Println("Send chatroomMessage", chatroomID, chatroomMessage)

	topic := chatroomMessageTopic(chatroomID)

	this.lock.RLock()
	subscribers := append([]ChatroomMessageCallback(nil), this.subscribers[topic]...)
	this.lock.RUnlock()

	for _, cb := range subscribers {
		cb(chatroomMessage)
	}
}

func (this *ChatroomEventbus) SubscribeToChatroomMessages(chatroomID int, cb ChatroomMessageCallback) {
	this.lock.Lock()
	defer this.lock.Unlock()

	topic := chatroomMessageTopic(chatroomID)
	this.subscribers[topic] = append(this.subscribers[topic], cb)
}

func (this *ChatroomEventbus) UnsubscribeToChatroomMessages(chatroomID int, cb ChatroomMessageCallback) {
	this.lock.Lock()
	defer this.lock.Unlock()

	topic := chatroomMessageTopic(chatroomID)
	subscribers := this.subscribers[topic]
	cbPointer := reflect.ValueOf(cb).Pointer()

	for i, subscriber := range subscribers {
		if reflect.ValueOf(subscriber).Pointer() != cbPointer {
			continue
		}

		this.subscribers[topic] = append(subscribers[:i], subscribers[i+1:]...)
		if len(this.subscribers[topic]) == 0 {
			delete(this.subscribers, topic)
		}
		return
	}

	log.Println("Unsubscribe error: chatroom message subscriber not found")
}

func chatroomMessageTopic(chatroomID int) string {
	return fmt.Sprintf("chatroomMessage:%d", chatroomID)
}

func GetChatroomEventbus(ctx *gin.Context) *ChatroomEventbus {
	e, _ := ctx.Get("chatroomEventbus")

	return e.(*ChatroomEventbus)
}

func SetChatroomEventbus(ctx *gin.Context, chatroomEventbus *ChatroomEventbus) {
	ctx.Set("chatroomEventbus", chatroomEventbus)
}
