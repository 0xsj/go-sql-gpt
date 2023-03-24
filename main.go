package main

import (
	"github.com/spf13/viper"
)

// define the openai choices struct

type OpenAIChoices struct{}

// define the openai respones struct

type OpenAIResponses struct{}

// define the openai request struct

type OpenAIRequest struct{}

// main entry point
func main() {
	// using viper to get the appropriate envs

	// if there is no api key, throw an error

	// create a context

	// initialize the client

	//set a preamble, which in this case is "give me a sql query"

	// use the client library to generate text completion for the given prompt

	// when we are loading / waiting for the response, show a spinner

	viper.SetConfigFile(".env")
	// client := openai.NewClient("")
}
