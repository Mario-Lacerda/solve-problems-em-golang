# Desafio Dio - Resolvendo Problemas Numéricos com Go



**Projeto: Resolvendo Problemas Numéricos com Go**

**Objetivo:**

Demonstrar como usar a biblioteca math do Go para resolver problemas numéricos.



**Código:**

go

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    // Calcular a raiz quadrada de 2
    raizQuadrada := math.Sqrt(2)

    // Calcular o seno de π/2
    seno := math.Sin(math.Pi / 2)

    // Calcular a tangente de π/4
    tangente := math.Tan(math.Pi / 4)

    // Imprimir os resultados
    fmt.Println("Raiz quadrada de 2:", raizQuadrada)
    fmt.Println("Seno de π/2:", seno)
    fmt.Println("Tangente de π/4:", tangente)
}
```



#### **Como executar:**



1. Crie um arquivo chamado `main.go` e copie o código acima para ele.

2. Abra um terminal e navegue até o diretório onde o arquivo `main.go` está localizado.

3. Execute o seguinte comando:

   

```plaintext
go run main.go
```



**Saída:**

```plaintext
Raiz quadrada de 2: 1.4142135623730951
Seno de π/2: 1
Tangente de π/4: 1
```



#### **Explicação:**

Este programa demonstra como usar as seguintes funções da biblioteca math:

- `Sqrt()` para calcular a raiz quadrada de um número.

- `Sin()` para calcular o seno de um ângulo em radianos.

- `Tan()` para calcular a tangente de um ângulo em radianos.

  

Você pode usar este programa como base para resolver problemas numéricos mais complexos em Go.


