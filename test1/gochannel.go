package main

import "fmt"

func main(){
myChan := make(chan int)
go func(){
	for i := 0; i < 1000; i++{
		myChan <- i
	}
	close(myChan)
	//myChan <- 456
}() // the go statement should be followed by a function call, 
//n := <- myChan
for n := range myChan{


fmt.Println("n =",n)
}
}
