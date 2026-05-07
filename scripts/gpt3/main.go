package main

import (
	"context"
	"fmt"
	"log"

	"github.com/j45k4/talktocow/config"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func main() {
	if config.OpenAIApiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	client := openai.NewClient(option.WithAPIKey(config.OpenAIApiKey))
	resp, err := client.Chat.Completions.New(
		context.Background(),
		openai.ChatCompletionNewParams{
			Model: openai.ChatModelGPT5_4Mini,
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage("Hello!"),
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
