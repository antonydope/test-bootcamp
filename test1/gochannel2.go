package main

import (
	"fmt"
	"time"
)

func sendData(ch chan<- int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		ch <- i
	}
	close(ch)
}

func receiveData(ch <-chan int) {
	for {
		data, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println("Received:", data)
	}
}

func main() {
	dataChannel := make(chan int)

	go sendData(dataChannel)
	go receiveData(dataChannel)

	// Allow some time for goroutines to complete
	time.Sleep(1 * time.Second)
}
