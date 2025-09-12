// Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro

/*
Integration with OpenAI
=======================
Neste script vamos integrar a API da OpenAI com
a linguagem Go.

RUN
---
go run integration_with_openai.go

FORMATANDO O CÓDIGO
-------------------
gofmt -w integration_with_openai.go

LEMBRANDO A INSTALAÇÃO DE DEPENDÊNCIAS
--------------------------------------
go get github.com/openai/openai-go/v2
go get github.com/openai/openai-go/v2/option

Pode usar o comando <go mod tidy> para limpar e reorganizar suas
dependências.
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
		log.Fatal("Erro ao carregar o arquivo .env")
	}
	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")

	// Criando cliente OpenAI usando variável de ambiente
	client := openai.NewClient(
		option.WithAPIKey(OPENAI_API_KEY),
	)
	fmt.Printf("Cliente OpenAI criado com sucesso: %T\n", client)

	// Fazendo uma requisição de chat completion
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Apenas diga: 'Olá Dr. Eddy Giusepe'"),
		},
		Model: openai.ChatModelGPT4o,
	})
	if err != nil {
		panic(err.Error())
	}

	// Imprimindo a resposta da IA
	fmt.Println("Resposta da OpenAI:")
	println(chatCompletion.Choices[0].Message.Content)
}
