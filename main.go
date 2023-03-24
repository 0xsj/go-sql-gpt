package main

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

// define the openai choices struct

type OpenAIChoices struct {
	Text         string
	Index        uint8
	Prob         *uint8
	FinishReason string
}

// define the openai respones struct
type OpenAIResponses struct {
	Id      *string
	Object  *string
	Created *uint64
	Model   *string
	Choices []OpenAIChoices
}

// define the openai request struct
type OpenAIRequest struct {
	Prompt     string
	Max_tokens uint64
	// Temperature *float32
}

// main entry point
func main() {
	// using viper to get the appropriate envs
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apiKey := viper.GetString("OPENAI_KEY")
	if apiKey == "" { // if there is no api key, throw an error
		panic("Missing API Key")
	}
	fmt.Println(apiKey)

	// initialize the client
	client := openai.NewClient(apiKey)
	response, err := client.CreateChatCompletion(
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
		fmt.Println("Completion error: &v\n", err)
		return
	}
	fmt.Println(response.Choices[0].Message.Content)

	// create a context

	//set a preamble, which in this case is "give me a sql query"

	// use the client library to generate text completion for the given prompt

	// when we are loading / waiting for the response, show a spinner

	// client := openai.NewClient("")
}
