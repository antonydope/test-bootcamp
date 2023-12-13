package main

import (
    "fmt"
    "time"
)

func main() {
    go printNumbers()
    go printLetters()

    // Keep the main goroutine running for a while
    time.Sleep(2 * time.Second)
}

func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Printf("%d ", i)
        time.Sleep(200 * time.Millisecond)
    }
}

func printLetters() {
    for i := 'a'; i <= 'e'; i++ {
        fmt.Printf("%c ", i)
        time.Sleep(200 * time.Millisecond)
    }
}
