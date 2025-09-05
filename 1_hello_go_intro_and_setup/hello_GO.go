/*
Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro

Script: hello_GO.go
===================
Este é um programa simples que exibe uma mensagem de boas-vindas no console.
Ele utiliza a função Println do pacote fmt para imprimir a mensagem no console.

Para executar o programa, você pode usar o seguinte comando ----> go run main.go

Link de estudo: https://go.dev/doc/effective_go

RUN:
---
go run hello_GO.go

FORMATANDO O CÓDIGO
-------------------
gofmt -w hello_GO.go
*/
package main // Declaração do pacote principal. Deve ser "main" para ser um programa executável.

import "fmt" // Importa o pacote fmt (format). É um pacote que contém funções para formatação de strings.

func main() { // Função principal. É a função que será executada automaticamente quando o programa for executado.

	fmt.Println("Olá, Mundo!")
	fmt.Println("Olá, mundo! Bem-vindo ao Go!")

	// Chamando a outra função:
	outraFuncao()
}

func outraFuncao() {
	fmt.Println("Aqui estou testando a outra função!")
}
