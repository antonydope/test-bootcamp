package main

import ("fmt"
"time")

func main(){
firstchannel := make(chan string)
secondchannel := make(chan string)

go func(){
	time.Sleep(4*time.Second)
	firstchannel <- "hello antony"
}()
go func (){
	time.Sleep(1*time.Second)
	secondchannel <- "hellobro"
}()

select{
case firstmessage := <- firstchannel: {
	fmt.Println("this is the firstmessage",firstmessage)
}
case secondmessage := <- secondchannel: {
	fmt.Println("this is the second message",secondmessage)
}

}
fmt.Println("the programme continues")

}