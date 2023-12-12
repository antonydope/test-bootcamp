package main

import "fmt"

type Rectangle struct{
height int
width  int
}
func (y Rectangle) areaOne()int{
return y.height * y.width
}
func main(){
rectangle1 := Rectangle{10,20}
area := rectangle1.areaOne()
fmt.Println("the are of the rectangle :",area)

fmt.Println("the are of the rectangle :",area1)
}