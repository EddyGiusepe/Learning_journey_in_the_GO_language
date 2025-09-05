// Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro

/*
RUN
---
go run practical_examples_with_arrays.go

FORMATANDO O CÓDIGO
-------------------
gofmt -w practical_examples_with_arrays.go
*/
package main

import "fmt"

func main() {
	// PASSO 8: EXEMPLOS PRÁTICOS
	fmt.Println("\n🔸 PASSO 8: Exemplos Práticos")
	fmt.Println("==============================")

	// Exemplo 1: Encontrar o maior e menor valor em um array:
	valores := [7]int{23, 45, 12, 67, 34, 89, 56}

	maior, menor := encontrarMaiorMenor(valores)
	fmt.Println("O array tem valores: ", valores)
	fmt.Printf("O maior valor é: %d e o menor valor é: %d\n", maior, menor)

	// Exemplo 2: Calcular a média de um array de números:
	temperaturas := [5]float64{23.5, 25.2, 22.8, 24.1, 26.3}

	fmt.Println("\nOs valores das 5 temperaturas são:", temperaturas)
	media_tem := calculaMedia(temperaturas)
	fmt.Printf("A média das temperaturas é: %.2f\n", media_tem)

}

// FUNÇÕES AUXILIARES
// ===================

func encontrarMaiorMenor(array [7]int) (int, int) {
	maior := array[0]
	menor := array[0]

	for _, valor := range array {
		if valor > maior {
			maior = valor
		}
		if valor < menor {
			menor = valor
		}
	}
	return maior, menor
}

func calculaMedia(array [5]float64) float64 {
	soma := 0.0
	for _, valor := range array {
		soma = soma + valor
	}
	return soma / float64(len(array))
}
