package main

import (
	"fmt"

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
	if apiKey == "" {
		panic("Missing API Key")
	}
	fmt.Println(apiKey)

	// if there is no api key, throw an error

	// create a context

	// initialize the client
	// client := openai.NewClient("")

	//set a preamble, which in this case is "give me a sql query"

	// use the client library to generate text completion for the given prompt

	// when we are loading / waiting for the response, show a spinner

	// client := openai.NewClient("")
}
