package main

import (
    "fmt"
    "time"
)

func main() {
    // Create two channels
    ch1 := make(chan string)
    ch2 := make(chan string)

    // Goroutine 1: Send "Hello" to ch1 after 2 seconds
    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "Hello"
    }()

    // Goroutine 2: Send "World" to ch2 after 1 second
    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "World"
    }()

    // Use select to wait for either ch1 or ch2 to receive a value
    select {
    case msg1 := <-ch1:
        fmt.Println("Received from ch1:", msg1)
    case msg2 := <-ch2:
        fmt.Println("Received from ch2:", msg2)
    }
}
