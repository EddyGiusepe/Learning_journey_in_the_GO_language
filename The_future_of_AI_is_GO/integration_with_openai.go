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
*/
package main

import (
	"fmt"
	"os"

	"github.com/openai/openai-go/v2"
)

func main() {
	fmt.Println("Integrando a API da OpenAI com a linguagem Go!")

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	fmt.Printf("Cliente OpenAI criado com sucesso: %T\n", client)
}
