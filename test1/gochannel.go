package main

import ("fmt"
"math/rand"
"time"
)
func doWork()int{
	time.Sleep(time.Second)
	return rand.Intn(100)
}
func main(){
myChan := make(chan int)
go func(){
	for i := 0; i < 1000; i++{
		result := doWork()
		myChan <- result
	}
	close(myChan)
	//myChan <- 456
}() // the go statement should be followed by a function call, 
//n := <- myChan
for n := range myChan{


fmt.Println("n =",n)
}
}
