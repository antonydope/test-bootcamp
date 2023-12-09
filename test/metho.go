package main

import "fmt"

type Rectangle struct {
	length int
	width int
}
 
func (rectangle Rectangle) getArea() int {
	return rectangle.length * rectangle.width
}

func (circum Rectangle) getCircum()int{
	return circum.length + circum.width 
}
func main (){
rect := Rectangle{4,5}
area := rect.getArea()
fmt.Println("the area of the rectangle is ",area)
halfCircumference := rect.getCircum()
fmt.Println("THE HALF CIRCUMFERENCE IS ", halfCircumference)
circumference := halfCircumference*2
fmt.Println("the cicumference is ", circumference)

}