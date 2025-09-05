// Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro

/*
Variables, Strings and Numbers
==============================
Neste script vamos aprender como declarar variáveis, strings e números em Go.

RUN
---
go run variables_strings_numbers.go
*/
package main

import "fmt"

func main() {

	// Strings:
	var nameOne string = "Eddy Giusepe"
	var nameTwo = "Karina"
	var nameThree string // É uma variável sem valor
	var nameFour string = "Brasil"
    
	// A seguir usamos o `Printf` para formatar a saída:
	fmt.Printf("Me chamo %s e sou casado com a %s e moramos em %s %s\n", nameOne, nameTwo, nameThree, nameFour) // %d para números, %f para decimais, etc.

	// A seguir usamos o `Println` para formatar a saída:
	fmt.Println("Me chamo " + nameOne + " e sou casado com a " + nameTwo + " e moramos em " + nameFour)

	// Outro exemplo de declaração de variáveis:
	var (
		a int = 10
		b int = 20
		c int = 30
	)
	// Mostramos os valores das variáveis:
	fmt.Println(a, b, c)
	// Fazemos a soma entre as variáveis:
	fmt.Println("A soma das variáveis é:", a + b + c)

	// Agora, usando booleanos:
	var d bool = true
	fmt.Printf("d: %v e %T\n", d, d)
	fmt.Println("A variável 'd' tem o seguinte valor:", d)
}
