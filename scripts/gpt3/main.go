package main

import (
	"context"
	"fmt"
	"log"

	"github.com/j45k4/talktocow/config"
	"github.com/sashabaranov/go-openai"
)

func main() {
	if config.OpenAIApiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	client := openai.NewClient(config.OpenAIApiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		return
	}

	for _, choise := range resp.Choices {
		fmt.Println(choise.Message.Content)
	}
}
