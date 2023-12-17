package main

import (
	"fmt"
	"time"
	
)
/*func sendsomeData(dataA chan<- string){
	for y := 'A'; y <= 'Z'; y++ {
		time.Sleep(100*time.Millisecond)
		dataA <- string(y)
	}
}
func receivesomeData(dataA <-chan string){
	for {
		receiver, ok := <- dataA
		if !ok {
			fmt.Println("The channel is closed")
			break
		}
		fmt.Println("te data being received is :",receiver)
	}
}
func printNumbers(){
	for i := 1; i <= 27; i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println("and the numbers correspending to are :",i)
	}
}
func main(){
	chan1 := make(chan string)
	go sendsomeData(chan1)
	go receivesomeData(chan1)
	go printNumbers()
	time.Sleep(4*time.Second)
}*/

func printAlphabet(mycha chan<- string) {
	for i := 'a'; i <= 'z'; i++ {
		time.Sleep(100*time.Millisecond)
		mycha <- string(i)
	
	}
	close(mycha)
}
/*func receiveAlphabet(mycha <-chan string){
		for {
			dataA ,ok := <-chan
		if !ok {
			fmt.Println("there was no data")
			break
		}
		fmt.Println(dataA)
	}
}*/

func receiveAlphabet(mycha <-chan string) {
	for {
		dataA, ok := <-mycha
		if !ok {
			fmt.Println("The channel is closed")
			break
		}
		fmt.Println("The data being received is:", dataA)
	}
}
func printNumber(){
	for x := 1; x <= 26; x++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(x)
	}
}
func main(){
	alphabetChannel := make(chan string)
	go receiveAlphabet(alphabetChannel)
	go printAlphabet(alphabetChannel)
	go printNumber()

	time.Sleep(5*time.Second)
}