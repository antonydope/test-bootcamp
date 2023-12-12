package main

import "fmt"

func main(){
john := "software developer"
johnPointer := &john
fmt.Println("CAREER : ", john)
fmt.Println("address :",johnPointer)
updateJohn(johnPointer)
fmt.Println("CAREER UPDATE :",john)
changedCareerAgain(johnPointer)
fmt.Println("TRANSFORMATION :",john)

}

func updateJohn(john *string){

*john = "lecturer of Software engineering" 

}
func changedCareerAgain(john *string){
	*john = "Pastor"
}