# Desafio Ping Pong com Goroutines

O desafioPingPong é um programa em Go que trabalha com dois pacotes e concorrência. Ou sejá, utilizando canais e goroutines para que o programa exiba, em alternância, as palavras Ping e Pong.

## Código Original

O código original consiste em três goroutines: `Pingar`, `Pongar` e `Imprimir`. As duas primeiras goroutines enviam mensagens ("Ping" ou "Pong") para um canal compartilhado, e a terceira goroutine imprime essas mensagens no console com um intervalo de 1 segundo. O programa principal cria os canais, inicia as goroutines e aguarda a entrada do usuário antes de encerrar o programa.

```go
package main

import (
	"fmt"
	"time"
)

// (Código original aqui), vide https://academiapme-my.sharepoint.com/:b:/g/personal/nubia_dio_me/EeTaf_s-LpxPij0xjbQaBisB7wP3LUIhdwAC3UFdARR4uw?e=CMHccy

func main() {
	var c chan string = make(chan string)

	go Pingar(c)
	go Imprimir(c)
	go Pongar(c)

	var entrada string
	fmt.Scanln(&entrada)
}
```

## Melhorias Implementadas

1. **Encerramento Gracioso:**
    - Adicionamos um mecanismo de encerramento para garantir que as goroutines sejam encerradas adequadamente. Um canal adicional chamado `done` é utilizado para sinalizar o encerramento.
    - Cada goroutine (`Pingar`, `Pongar` e `Imprimir`) verifica regularmente o canal `done` e encerra a execução quando recebe um sinal de encerramento.

2. **Uso de `select`:**
    - Utilizamos a instrução `select` para tratar múltiplos canais. Isso torna o código mais conciso e eficiente.

3. **Encerramento Adequado:**
    - Fechamos os canais após o uso para evitar vazamento de recursos. O encerramento do canal `c` é feito quando a goroutine recebe o sinal de encerramento.

```go
package main

import (
	"fmt"
	"time"
)

// (Código melhorado aqui), vide main.go

func main() {
	c := make(chan string)
	done := make(chan bool)

	go Pingar(c, done)
	go Imprimir(c, done)
	go Pongar(c, done)

	var entrada string
	fmt.Scanln(&entrada)

	// Sinaliza o encerramento para as goroutines
	close(done)

	// Aguarda um pouco para garantir que todas as goroutines terminem
	time.Sleep(time.Second)
}
```

## Como Usar

1. Clone o repositório:
   ```bash
   git clone https://github.com/gatinhodev/desafioPingPong.git
   ```

2. Execute o programa:
   ```bash
   cd desafioPingPong
   go run main.go
   ```

3. Siga as instruções no console para encerrar o programa.

Sinta-se à vontade para ajustar o código conforme necessário ou adicionar recursos adicionais. Este é um ponto de partida simples para experimentar com goroutines e canais em Go.
