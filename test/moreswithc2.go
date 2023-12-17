package main

import (
	"fmt"
	"time"
)

func main(){
	/*Antony := "software developer"
	fmt.Print("Mr ",Antony, "nowadays")
	switch Antony {
	case 1:
		fmt.Println("decidede to try python programmin")
	case 2:
		fmt.Println("went in for a job as a tech specialist")
	case 3:
		fmt.Println("am so cool in golang")
	//default:
		//fmt.Println("now i am a software developer in go ")
	}*/
	i := 3
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }
	t := time.Now()
	fmt.Println("the time now is :",t)
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }
}