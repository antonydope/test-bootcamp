package main

import (
    "fmt"
    
)

func main() {
    // Creating a buffered channel with a capacity of 3
    ch := make(chan int, 3)

    // Send values to the channel without blocking until the buffer is full
    ch <- 1
    ch <- 2
    ch <- 3

    // Receiving values from the channel without blocking until the buffer is empty
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
