// Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro

/*
Integration with OpenAI - Interactive Chat
==========================================
In this script, I will integrate the OpenAI API with
the Go language for interactive conversations.

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
go get github.com/joho/godotenv

You can use the command <go mod tidy> to clean and reorganize your
dependencies.
*/
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"Learning_journey_in_the_GO_language/utils"
	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(utils.Red + "Error loading the .env file" + utils.Reset)
	}

	// Creating OpenAI client using environment variable:
	client := openai.NewClient(option.WithAPIKey(os.Getenv("OPENAI_API_KEY")))
	log.Printf(utils.Green+"OpenAI client created successfully: %T"+utils.Reset, client)

	// Initializing the conversation history:
	var messages []openai.ChatCompletionMessageParamUnion

	// Initial system message (optional):
	messages = append(messages, openai.SystemMessage(`You're a helpful and friendly assistant.
	                                                  If you don't know the answer, say: 'Sorry, I don't know the answer for that.'
	                                                  Moreover, always answer in Spanish and at the end add an emoji depending on the response.`))

	fmt.Println(utils.Blue + "💬 Interactive Chat with OpenAI started!" + utils.Reset)
	fmt.Println(utils.Yellow + "Type 'quit' or 'exit' to end the conversation." + utils.Reset)
	fmt.Println(strings.Repeat("-", 47))

	// Scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Request user input:
		fmt.Print(utils.Green + "You: " + utils.Reset)
		if !scanner.Scan() {
			break
		}

		userInput := strings.TrimSpace(scanner.Text())

		// Check if the user wants to quit
		if strings.ToLower(userInput) == "quit" ||
			strings.ToLower(userInput) == "exit" {
			fmt.Println(utils.Blue + "👋 Bye! It was a pleasure talking to you!" + utils.Reset)
			break
		}

		// Check if the input is not empty:
		if userInput == "" {
			continue
		}

		// Add the user's message to the history:
		messages = append(messages, openai.UserMessage(userInput))

		// Make the request to the OpenAI API:
		fmt.Print(utils.Blue + "AI: " + utils.Reset)
		chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
			Messages: messages,
			Model:    openai.ChatModelGPT4o,
		})

		if err != nil {
			fmt.Printf(utils.Red+"❌ Error communicating with the AI: %v"+utils.Reset+"\n", err)
			continue
		}

		// Get and display the AI's response:
		aiResponse := chatCompletion.Choices[0].Message.Content
		fmt.Println(aiResponse)

		// Add the AI's response to the history:
		messages = append(messages, openai.AssistantMessage(aiResponse))

		fmt.Println(strings.Repeat("-", 50))
	}

	// Check for scanner errors:
	if err := scanner.Err(); err != nil {
		log.Printf(utils.Red+"Error reading input: %v"+utils.Reset, err)
	}
}
