package main

import "fmt"

func modifyArray(arr [3]int) {
    arr[0] = 100
}

func modifySlice(slice []int) {
    slice[0] = 100
}

func main() {
    // Passing an array
    arr := [3]int{1, 2, 3}
    modifyArray(arr)
    fmt.Println("Array after modification:", arr)

    // Passing a slice
    slice := []int{1, 2, 3}
    modifySlice(slice)
    fmt.Println("Slice after modification:", slice)
}
