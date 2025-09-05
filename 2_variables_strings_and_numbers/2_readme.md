# <h1 align="center"><font color="gree">GO Learning Journey</font></h1>

<font color="pink">Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro</font>

## <font color="red">`Estudo 2:` Variáveis, Strings e Números</font>

Neste estudo vamos aprender como criar e manipular variáveis, strings e números em `Go`.

### <font color="blue">`2.1` Variáveis</font>
Uma variável é um nome dado a uma área de armazenamento que os programas podem manipular. O nome de uma 
variável pode ser composto por letras, dígitos e o caractere sublinhado. Ele deve começar com uma letra ou sublinhado.

O código geral para declarar uma variável em golang é:

```go
var <nome da variável> <tipo da variável>
```

### <font color="blue">`2.2` Números</font>
Variáveis são tipos aritméticos e representam os seguintes valores em todo o programa: 

a) tipos inteiros

b) tipos de ponto flutuante

Para definir um inteiro, use a seguinte sintaxe:

```go
var a int=4
var b,c int
b=5
c=10
fmt.Println(a)
fmt.Println(b + c)
```

Para definir um número de ponto flutuante, use a seguinte sintaxe:

```go
var d float64 = 9.14
fmt.Println(d)
```

### <font color="blue">`2.3` Strings</font>
Strings em `GO`são definidas com aspas duplas.

```go
var s string = "Não se preocupe com apóstrofos"
fmt.Println(s)
```
Também, podemos definir várias strings de linha envolvendo a string entre `` aspas.

```go
var s string = `Esta é uma string que abrange várias linhas
Esta é a segunda linha
E esta é a terceira`

fmt.Println(s)
```

### <font color="blue">`2.4` Booleanos</font>
Golang suporta valores booleanos com as palavras-chave `true` e `false`.

```go
var b bool = true
fmt.Println(b)
```

### <font color="blue">`2.5` Declaração abreviada</font>
A notação `:=` serve tanto como declaração quanto como inicialização. Por exemplo: `foo := "bar"` é equivalente a `var foo string = "bar"`.

```go
a := 9
b := "golang"
c := 4.17
d := false
e := "Hello"
f := `Do you like golang, so far?`
g := 'M'
h := true

fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
fmt.Println(d)
fmt.Println(e)
fmt.Println(f)
fmt.Println(g)
fmt.Println(h)
```



















Thank God 🤓!