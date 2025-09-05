// Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro
/*
RUN
---
go run arrays.go
*/

package main

import "fmt"

func main() {
	fmt.Println("=== EXEMPLOS PRÁTICOS DE ARRAYS EM GO ===")

	// PASSO 1: DECLARAÇÃO BÁSICA DE ARRAYS
	fmt.Println("\n🔸 PASSO 1: Declaração Básica de Arrays")
	fmt.Println("=========================================")

	// Declarando um array de 5 inteiros (inicializado com zeros):
	var numeros [5]int
	fmt.Printf("Array declarado e iniciado com zeros: %v\n", numeros)
	fmt.Println("Tamanho do array:", len(numeros))

	// PASSO 2: INICIALIZAÇÃO DE ARRAYS
	fmt.Println("\n🔸 PASSO 2: Inicialização de Arrays")
	fmt.Println("=========================================")

	// Forma 1: Inicialização com valores específicos:
	//idades := [5]int{25, 30, 35, 40, 45} // forma 1
	var idades = [5]int{25, 30, 35, 40, 45} // forma 2
	fmt.Printf("Array de idades inicializado (todos os valores): %v\n", idades)

	// Forma 2: Inicialização parcial (resto fica zero):
	notas := [5]float64{8.5, 9.2}
	fmt.Println("Array de notas parcialmente inicializado:", notas)

	// Forma 3: Go determina o tamanho automaticamente:
	frutas := [...]string{"maçã", "banana", "laranja", "uva", "manga"}
	fmt.Println("Array de frutas (tamanho automático):", frutas)
	fmt.Println("Tamanho do array de frutas:", len(frutas))

	// Forma 4: Inicialização com índices específicos:
	cores := [5]string{0: "vermelho", 2: "azul", 3: "verde"}
	fmt.Printf("Array com índices específicos: %v\n", cores)

	// PASSO 3: ACESSO E MODIFICAÇÃO DE ELEMENTOS
	fmt.Println("\n🔸 PASSO 3: Acesso e Modificação de Elementos")
	fmt.Println("=============================================")

	// Acessando elementos individuais:
	fmt.Printf("Primeira idade: %d\n", idades[0])
	fmt.Println("Terceira idade:", idades[2])
	fmt.Println("")
	fmt.Println("Última fruta do meu array de frutas:", frutas[len(frutas)-1])

	// Modificando elementos:
	idades[2] = 33 // Alterando o terceiro elemento
	fmt.Println("Idades após modificação: ", idades)

	// PASSO 4: ITERAÇÃO ATRAVÉS DE ARRAYS
	fmt.Println("\n🔸 PASSO 4: Iteração Através de Arrays")
	fmt.Println("======================================")

	// Método 1: For loop tradicional:
	fmt.Println("Método 1 - For loop tradicional:")
	for i := 0; i < len(frutas); i++ {
		fmt.Printf("  Índice %d: %s\n", i, frutas[i])
	}

	// Método 2: Range apenas com valores (mais idiomático em Go):
	fmt.Println("Método 2 - Using range (recomendado):")
	for indice, fruta := range frutas {
		fmt.Printf("  Posição %d: %s\n", indice, fruta)
	}

	// Método 3: Range apenas com valores (ignorando índice):
	fmt.Println("Método 3 - Range apenas com valores:")
	for _, fruta := range frutas { // range --> É uma Palavra-chave para iteração
		fmt.Println("  Fruta:", fruta)
	}
}
