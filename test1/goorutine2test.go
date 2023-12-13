/*package main

import ("fmt"
"time"
)
func PrintAlphabet(){
for x := 'A'; x <= 'Z'; x++ {
	time.Sleep(100* time.Millisecond)
	fmt.Println("the alphabets are :",string(x))
}
}
func PrintNbr(){
	for y := 0; y <= 26; y++{
		time.Sleep(105* time.Millisecond)
		fmt.Println("the numbers are :", y)
	}
}
func main(){
	go PrintAlphabet()
	go PrintNbr()

	time.Sleep(1 * time.Second)
}*/
package main

import (
	"fmt"
	"time"
)

func PrintAlphabet() {
	for x := 'A'; x <= 'Z'; x++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("the alphabets are :", string(x))
	}
}

func PrintNbr() {
	for y := 1; y <= 27; y++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("the numbers are :", y)
	}
}

func main() {
	go PrintAlphabet() // Start a goroutine to print the alphabet
	go PrintNbr()      // Start another goroutine to print numbers

	// Sleep for a sufficient duration to allow goroutines to complete
	time.Sleep(100 * time.Second)
}
