package bot

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/j45k4/talktocow/eventbus"
	"github.com/j45k4/talktocow/models"
	"github.com/sashabaranov/go-openai"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type CowGPT struct {
	Eventbus   *eventbus.Eventbus
	Client     *openai.Client
	CowGPTUser *models.User
	Ctx        context.Context
	DB         *sql.DB
}

// func NewCowGPT(eb *eventbus.Eventbus) *CowGPT {
// 	return &CowGPT{
// 		eventbus: eb,
// 	}
// }

const systemMessage = "You are helpfull bot developed by COW named CowGPT. You can ask me anything and I will try to answer you."

func (cowGPT *CowGPT) handleChatroomMessage(msg *eventbus.ChatroomMessage) {
	if msg.UserID == cowGPT.CowGPTUser.ID {
		return
	}

	// Fetch last messages
	messages, err := models.Messages(
		qm.Where("chatroom_id = ?", msg.ChatroomID),
		qm.OrderBy("written_at desc"),
		qm.Limit(10),
	).All(cowGPT.Ctx, cowGPT.DB)

	if err != nil {
		log.Printf("error fetching messages: %v", err)
		return
	}

	chatcompletionMessages := []openai.ChatCompletionMessage{}

	contextLenght := len(systemMessage)

	for _, message := range messages {
		contextLenght += len(message.MessageText.String)

		if contextLenght > 2000 {
			break
		}

		role := openai.ChatMessageRoleUser

		if message.UserID == cowGPT.CowGPTUser.ID {
			role = openai.ChatMessageRoleAssistant
		}

		chatcompletionMessages = append([]openai.ChatCompletionMessage{{
			Role:    role,
			Content: message.MessageText.String,
		}}, chatcompletionMessages...)
	}

	chatcompletionMessages = append([]openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are helpfull bot developed by COW named CowGPT. You can ask me anything and I will try to answer you.",
		}},
		chatcompletionMessages...,
	)

	res, err := cowGPT.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: chatcompletionMessages,
		},
	)

	if err != nil {
		fmt.Printf("error creating chat completion: %v", err)

		return
	}

	fmt.Println("ChatCompletionResponse: ", res)

	choice := res.Choices[0].Message.Content

	// use v4 uuid as id
	refence := uuid.New().String()

	nowUTC := time.Now().UTC()

	//save chatroom message to db
	chatroomMessage := models.Message{
		ChatroomID:       msg.ChatroomID,
		MessageText:      null.NewString(choice, true),
		WrittenAt:        nowUTC,
		TransmitedAt:     nowUTC,
		ServerReceivedAt: nowUTC,
		UserID:           cowGPT.CowGPTUser.ID,
		CreatedAt:        nowUTC,
		Reference:        null.NewString(refence, true),
	}

	err = chatroomMessage.Insert(cowGPT.Ctx, cowGPT.DB, boil.Infer())

	if err != nil {
		fmt.Printf("error inserting chatroom message: %v", err)

		return
	}

	responseEvent := eventbus.Event{
		ChatroomMessage: &eventbus.ChatroomMessage{
			ChatroomID:       msg.ChatroomID,
			MessageText:      choice,
			WrittenAt:        nowUTC,
			TransmitedAt:     nowUTC,
			ServerReceivedAt: nowUTC,
			UserID:           cowGPT.CowGPTUser.ID,
			CreatedAt:        nowUTC,
			Reference:        refence,
			Bot:              true,
		},
	}

	fmt.Println("responseEvent: ", responseEvent)

	cowGPT.Eventbus.Publish(responseEvent)
}

func (c *CowGPT) Run() {
	channel := c.Eventbus.Subscribe()

	for event := range channel {
		log.Printf("CowGPT received event: %+v\n", event)

		if event.ChatroomMessage != nil {
			go c.handleChatroomMessage(event.ChatroomMessage)
		}
	}

	log.Println("CowGPT stopped")
}
