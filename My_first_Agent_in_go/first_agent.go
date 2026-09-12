/*
Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro

Script: first_agent.go
======================
Este é um programa que cria um agente de pesquisa usando o modelo
Gemini Flash da Google. O agente é capaz de pesquisar tópicos de
forma detalhada e retornar os resultados.

Link de estudo:
- https://adk.dev/
- https://adk.dev/get-started/go/

Para interagir no terminal
--------------------------
go run first_agent.go


RUN (API para o chat React; sem ADK Web):
---
go run first_agent.go web api

UI:
---
cd web && npm install && npm run dev

ADK Web (só debug):
---
go run first_agent.go web api webui

FORMATANDO O CÓDIGO
--------------------
gofmt -w first_agent.go
*/
package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/adk/v2/agent"
	"google.golang.org/adk/v2/agent/llmagent"
	"google.golang.org/adk/v2/cmd/launcher"
	"google.golang.org/adk/v2/cmd/launcher/full"
	"google.golang.org/adk/v2/model/gemini"
	"google.golang.org/adk/v2/tool"
	"google.golang.org/adk/v2/tool/geminitool"
	"google.golang.org/genai"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		if err = godotenv.Load(".env"); err != nil {
			log.Fatal("Error loading the .env file")
		}
	}

	ctx := context.Background()

	model, err := gemini.NewModel(ctx, "gemini-flash-latest", &genai.ClientConfig{
		APIKey: os.Getenv("GOOGLE_API_KEY"),
	})
	if err != nil {
		log.Fatalf("Failed to create model: %v", err)
	}

	researcher, err := llmagent.New(llmagent.Config{
		Name:        "researcher",
		Model:       model,
		Instruction: "You help users to search topics in a detailed way.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
	})
	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	config := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(researcher),
	}

	l := full.NewLauncher()
	if err = l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
