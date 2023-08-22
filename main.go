package main

import (
	"fmt"
	"time"
)

// input channel
func sendNumber(sendCh chan<- int) {
	fmt.Println("Send Number channel")
	sendCh <- 4
}

// output channel
func getNumber(getCh <-chan int) {
	num := <- getCh
	fmt.Printf("Recieved %v", num)
}

func main() {
	ch := make(chan int)
	go sendNumber(ch)
	go getNumber(ch)
	time.Sleep(1 * time.Second)
}
