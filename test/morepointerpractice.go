package main 

import ("fmt"
"time")
func hellobro(x *string){
	time.Sleep(100*time.Millisecond)
	fmt.Println("The address for antony is :",&antony)
}
func amcool(x *string){
	time.Sleep(101*time.Millisecond)
	fmt.Println("also i can print the address: ",antonyPointer)
}
func heyB(x string){
	time.Sleep(102*time.Millisecond)
	fmt.Println("I am antony and ",antony)
}
func main(){
antony := "am awesome"
antonyPointer := &antony
go hellobro(antony)
go amcool(antonyPointer)
go heyB(antony)
time.Sleep(4*time.Second)
}