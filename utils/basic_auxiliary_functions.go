// Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro

/*
Package utils fornece funções utilitárias para operações com arrays e slices.
Este pacote é usado em outros scripts para realizar operações com arrays.

FORMATANDO O CÓDIGO
-------------------
gofmt -w basic_auxiliary_functions.go
*/

package utils

// EncontrarMaiorMenor encontra o maior e menor valor em um array de inteiros.
// Parâmetros:
//   - array: array de 7 inteiros
//
// Retorna o maior e menor valor encontrados.
func EncontrarMaiorMenor(array [7]int) (int, int) {
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

// CalculaMedia calcula a média aritmética de um array de float64.
// Parâmetros:
//   - array: array de 5 float64
//
// Retorna a média dos valores.
func CalculaMedia(array [5]float64) float64 {
	soma := 0.0
	for _, valor := range array {
		soma = soma + valor
	}
	return soma / float64(len(array))
}

// SomarElementos retorna a soma de todos os elementos de um array de inteiros.
// Parâmetros:
//   - array: array de inteiros
//
// Retorna a soma dos valores.
func SomarElementos(array []int) int {
	soma := 0
	for _, valor := range array {
		soma += valor
	}
	return soma
}

// ContarElementos conta quantas vezes um elemento aparece no array de inteiros.
// Parâmetros:
//   - array: array de inteiros
//   - elemento: elemento a ser contado
//
// Retorna o número de vezes que o elemento aparece no array.
func ContarElementos(array []int, elemento int) int {
	contador := 0
	for _, valor := range array {
		if valor == elemento {
			contador++
		}
	}
	return contador
}
