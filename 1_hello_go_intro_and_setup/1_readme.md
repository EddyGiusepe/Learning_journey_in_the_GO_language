# <h1 align="center"><font color="gree">GO Learning Journey</font></h1>

<font color="pink">Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro</font>

## <font color="red">`Estudo 1:` Conceitos básicos e setup em GO</font>

Neste estudo vamos aprender como criar um programa simples em Go que exibe uma mensagem de boas-vindas no console.

### <font color="blue">`1.1` Primeiros Passos com Go</font>

O primeiro é instalar o Go:

```bash
sudo apt update
sudo apt upgrade golang-go
```

Logo verificamos a instalação:

```bash
go version
```

### <font color="blue">`1.2` Estrutura do Primeiro Programa (`hello_GO.go`)</font>

A seguir vamos a explicar cada parte do código:

```go
package main 

import "fmt"

func main() { 
	
	fmt.Println("Olá, Mundo!")
	fmt.Println("Olá, mundo! Bem-vindo ao Go!")
}
```

#### Explicação:

- **`package main`**: Todo programa `Go` começa declarando a qual pacote pertence. O pacote `main` é especial - indica que este arquivo é um ponto de entrada executável, não uma biblioteca.

- **`import "fmt"`**: Importa o pacote de formatação padrão do Go, necessário para funções de entrada/saída como `Println`. Os imports são explícitos em `Go`.

- **`func main()`**: Ponto de entrada do programa. Quando executamos um programa `Go`, a função `main()` é chamada automaticamente. Todo programa executável precisa ter uma função `main` no pacote `main`.


### <font color="blue">`1.3` Executando o Programa</font>

Para executar o programa, você pode usar o seguinte comando:

```bash
go run hello_GO.go
```


Thank God 🤓!
