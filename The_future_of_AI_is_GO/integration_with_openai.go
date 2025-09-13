// Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro

/*
Integration with OpenAI - Interactive Chat
==========================================
In this script we will integrate the OpenAI API with
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

	// Inicializar o histórico da conversa:
	var messages []openai.ChatCompletionMessageParamUnion

	// Mensagem inicial do sistema (opcional):
	messages = append(messages, openai.SystemMessage(`Você é um assistente útil e amigável.
	                                                  Se não souber a resposta, diga: 'Desculpe, não sei a resposta para isso.'.
	                                                  Ademais, sempre responda em Espanhol e  no final adicione 
													  um emoji dependendo da resposta.`))

	fmt.Println(utils.Blue + "💬 Chat Interativo com OpenAI iniciado!" + utils.Reset)
	fmt.Println(utils.Yellow + "Digite 'sair', 'quit' ou 'exit' para encerrar a conversa." + utils.Reset)
	fmt.Println(strings.Repeat("-", 50))

	// Scanner para ler entrada do usuário
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Solicitar entrada do usuário:
		fmt.Print(utils.Green + "Você: " + utils.Reset)
		if !scanner.Scan() {
			break
		}

		userInput := strings.TrimSpace(scanner.Text())

		// Verificar se o usuário quer sair
		if strings.ToLower(userInput) == "sair" ||
			strings.ToLower(userInput) == "quit" ||
			strings.ToLower(userInput) == "exit" {
			fmt.Println(utils.Blue + "👋 Tchau! Foi um prazer conversar com você!" + utils.Reset)
			break
		}

		// Verificar se a entrada não está vazia
		if userInput == "" {
			continue
		}

		// Adicionar a mensagem do usuário ao histórico
		messages = append(messages, openai.UserMessage(userInput))

		// Fazer a requisição para a OpenAI API
		fmt.Print(utils.Blue + "IA: " + utils.Reset)
		chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
			Messages: messages,
			Model:    openai.ChatModelGPT4o,
		})

		if err != nil {
			fmt.Printf(utils.Red+"❌ Erro ao comunicar com a IA: %v"+utils.Reset+"\n", err)
			continue
		}

		// Obter e exibir a resposta da IA
		aiResponse := chatCompletion.Choices[0].Message.Content
		fmt.Println(aiResponse)

		// Adicionar a resposta da IA ao histórico
		messages = append(messages, openai.AssistantMessage(aiResponse))

		fmt.Println(strings.Repeat("-", 50))
	}

	// Verificar erros do scanner
	if err := scanner.Err(); err != nil {
		log.Printf(utils.Red+"Erro ao ler entrada: %v"+utils.Reset, err)
	}
}
