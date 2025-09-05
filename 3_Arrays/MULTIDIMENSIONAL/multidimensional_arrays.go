// Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro

/*
RUN
---
go run multidimensional_arrays.go

FORMATANDO O CÓDIGO
-------------------
gofmt -w multidimensional_arrays.go
*/
package main

import (
	"fmt"
)

func main() {
	// PASSO 5: ARRAYS MULTIDIMENSIONAIS
	fmt.Println("\n🔸 PASSO 5: Arrays Multidimensionais")
	fmt.Println("=====================================")

	// Array bidimensional (maztriz 3x3):
	var matriz [3][3]int

	// Preenchendo a matriz com valores:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matriz[i][j] = (i+1)*10 + (j + 1)
		}
	}
	fmt.Println("Matriz 3x3:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%4d", matriz[i][j])
		}
		fmt.Println()
	}

	// Array 3D inicializado
	cubo := [2][2][2]int{
		{{1, 2}, {3, 4}},
		{{5, 6}, {7, 8}},
	}
	fmt.Printf("Cubo 3D: %v\n", cubo)

	// PASSO 6: COMPARAÇÃO DE ARRAYS
	fmt.Println("\n🔸 PASSO 6: Comparação de Arrays")
	fmt.Println("=================================")

	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	array3 := [3]int{1, 2, 4}

	fmt.Printf("array1: %v\n", array1)
	fmt.Printf("array2: %v\n", array2)
	fmt.Printf("array3: %v\n", array3)
	fmt.Printf("array1 == array2: %t\n", array1 == array2)
	fmt.Printf("array1 == array3: %t\n", array1 == array3)

	// PASSO 7: PASSANDO ARRAYS PARA FUNÇÕES
	fmt.Println("\n🔸 PASSO 7: Arrays em Funções")
	fmt.Println("==============================")

	numeros_teste := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("Array original: %v\n", numeros_teste)

	// Passando array a uma função para calcular a soma de seus elementos:
	resultado_soma := somarElementos(numeros_teste)
	fmt.Printf("Soma dos elementos: %d\n", resultado_soma)

	// Passando array a uma função para calcular o dobro de cada elemento:
	result_dobro := dobrarElementos(numeros_teste)
	fmt.Printf("Array após dobrar: %v\n", result_dobro)
}

// FUNÇÕES AUXILIARES
// ===================

// Função que recebe um array e retorna a soma de seus elementos:
// somarElementos calcula a soma de todos os elementos de um array de 5 inteiros.
//
// Parâmetros:
//   - array: array de 5 números inteiros
//
// Retorna:
//   - int: soma total de todos os elementos
//
// Exemplo:
//
//	numeros := [5]int{10, 20, 30, 40, 50}
//	resultado := somarElementos(numeros) // resultado = 150
func somarElementos(array [5]int) int {
	soma := 0
	for _, valor := range array {
		soma = soma + valor
	}
	return soma
}

// Função que recebe um array e retorna um novo array com cada elemento dobrado
// dobrarElementos dobra cada elemento de um array de 5 inteiros.
//
// Parâmetros:
//   - array: array de 5 números inteiros
//
// Retorna:
//   - [5]int: novo array com cada elemento dobrado
//
// Exemplo:
//
//	numeros := [5]int{10, 20, 30, 40, 50}
//	dobro := dobrarElementos(numeros) // dobro = [20, 40, 60, 80, 100]
func dobrarElementos(array [5]int) [5]int {
	var result [5]int
	for i := range array {
		result[i] = array[i] * 2
	}
	return result
}
