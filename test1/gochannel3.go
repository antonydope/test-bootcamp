package main

import ("fmt"
"time")
func sendsomeData(dataA chan<- string){
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
}