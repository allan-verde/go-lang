package main

import (
	"fmt"
	"time"
)

const (
	repeticoes int = 5
	tempoEspera time.Duration = time.Second
)

func pingar(c chan string) {
	for i := 0; i < repeticoes ; i++ {
		time.Sleep(1 * tempoEspera)
		c <- "ping"
	}
}

func pongar(c chan string) {
	for i := 0; i < repeticoes ; i++ {
		time.Sleep(1 * tempoEspera)
		c <- "pong"
	}
}

func main() {

	var c chan string = make(chan string)

	go pingar(c)
	go pongar(c)

	for i := 1; i <= (repeticoes * 2); i++  {
		fmt.Println(<- c)
		time.Sleep(tempoEspera * 1)
	}

	fmt.Println("Fim!")
}