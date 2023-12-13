package main

import (
	"fmt"
	"time"
)
func main(){
	dtaChannel := make(chan int)
	go SendData(dtaChannel)
	go ReceiverData(dtaChannel)
	time.Sleep(3*time.Second)
}
func SendData(ch chan<- int){
	for i := 0; i < 26; i++{
		time.Sleep(100 * time.Millisecond)
		ch <- i
	}
}
func ReceiverData(ch <-chan int) {
for {
	newdata, ok := <- ch
	if !ok {
		fmt.Println("data not received and the channel is closed")
		break
	}
	fmt.Println("the data received is:",newdata)
} 
}
/*func receiveData(ch <-chan int) {
	for {
		data, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println("Received:", data)
	}
}*/


