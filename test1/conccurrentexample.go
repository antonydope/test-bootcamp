package main

import (
    "fmt"
    "sync"
    "time"
)

func process(data int, resultChan chan int) {
    // Simulate some processing time
    time.Sleep(time.Millisecond * time.Duration(data))

    // Send the processed result to the result channel
    resultChan <- data * 2
}

func main() {
    // Create a buffered channel to hold processed results
    resultChan := make(chan int, 10)

    // Create a wait group to wait for all goroutines to finish
    var wg sync.WaitGroup

    // Start multiple goroutines to process data concurrently
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go func(data int) {
            defer wg.Done()
            process(data, resultChan)
        }(i)
    }

    // Close the result channel when all processing is done
    go func() {
        wg.Wait()
        close(resultChan)
    }()

    // Use select to wait for results from the result channel
    for {
        select {
        case result, ok := <-resultChan:
            if !ok {
                // The channel is closed, all processing is done
                fmt.Println("All processing is done.")
                return
            }
            fmt.Println("Processed result:", result)
        }
    }
}
