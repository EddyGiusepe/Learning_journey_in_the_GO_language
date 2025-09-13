// Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro

/*
Integration with OpenAI
=======================
In this script we will integrate the OpenAI API with
the Go language.

RUN
---
go run integration_with_openai.go

FORMATTING THE CODE
-------------------
gofmt -w integration_with_openai.go

REMEMBERING THE INSTALLATION OF DEPENDENCIES
--------------------------------------
go get github.com/openai/openai-go/v2
go get github.com/openai/openai-go/v2/option

You can use the command <go mod tidy> to clean and reorganize your
dependencies.
*/
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")

	// Creating OpenAI client using environment variable:
	client := openai.NewClient(
		option.WithAPIKey(OPENAI_API_KEY),
	)
	fmt.Printf("OpenAI client created successfully: %T\n", client)

	// Making a chat completion request:
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Just say: 'Hello Dr. Eddy Giusepe'"),
		},
		Model: openai.ChatModelGPT4o,
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Response from the IA:")
	println(chatCompletion.Choices[0].Message.Content)
}
