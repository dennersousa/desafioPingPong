package main

import (
	"fmt"
	"time"
)

func Pingar(c chan string, done chan bool) {
	for i := 0; ; i++ {
		select {
		case c <- "Ping":
		case <-done:
			close(c)
			return
		}
	}
}

func Pongar(c chan string, done chan bool) {
	for i := 0; ; i++ {
		select {
		case c <- "Pong":
		case <-done:
			close(c)
			return
		}
	}
}

func Imprimir(c chan string, done chan bool) {
	for {
		select {
		case msg, ok := <-c:
			if !ok {
				// Canal fechado, encerrar a goroutine
				return
			}
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
		case <-done:
			close(c)
			return
		}
	}
}

func main() {
	c := make(chan string)
	done := make(chan bool)

	go Pingar(c, done)
	go Imprimir(c, done)
	go Pongar(c, done)

	var entrada string
	fmt.Scanln(&entrada)

	// Sinalize o encerramento para as goroutines
	close(done)

	// Aguarde um pouco para garantir que todas as goroutines terminem
	time.Sleep(time.Second)
}
