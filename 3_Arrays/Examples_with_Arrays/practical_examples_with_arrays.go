// Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro

/*
Neste script vamos aprender como usar arrays em Go.
Ademais, vamos usar o "package utils" para operações
com arrays.

RUN
---
go run practical_examples_with_arrays.go

FORMATANDO O CÓDIGO
-------------------
gofmt -w practical_examples_with_arrays.go
*/
package main

import (
	"Learning_journey_in_the_GO_language/utils"
	"fmt"
)

func main() {
	// PASSO 8: EXEMPLOS PRÁTICOS
	fmt.Println("\n🔸 PASSO 8: Exemplos Práticos")
	fmt.Println("==============================")

	// Exemplo 1: Encontrar o maior e menor valor em um array:
	valores := [7]int{23, 45, 12, 67, 34, 89, 56}

	maior, menor := utils.EncontrarMaiorMenor(valores)
	fmt.Println("O array tem valores: ", valores)
	fmt.Printf("O maior valor é: %d e o menor valor é: %d\n", maior, menor)

	// Exemplo 2: Calcular a média de um array de números:
	temperaturas := [5]float64{23.5, 25.2, 22.8, 24.1, 26.3}

	fmt.Println("\nOs valores das 5 temperaturas são:", temperaturas)
	media_tem := utils.CalculaMedia(temperaturas)
	fmt.Printf("A média das temperaturas é: %.2f\n", media_tem)

}
