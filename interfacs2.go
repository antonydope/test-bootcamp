package main

import "fmt"

type circle struct{
	radius float64
}
type rectangle struct{
	height , width int
}
type HelloShape interface{
	areaYEAH()
}
func printArea(s HelloShape){
	fmt.Println("OKAY BRO THE AREA IS :",s.areaYEAH())
}
func (s circle) areaYEAH() float64{
	return 3.14 * s.radius
}
func (y rectangle) areaYEAH() int {
	return y.width*y.height
}
func main(){
	newCircle := circle{21.7}
	newRectangle := rectangle{20,10}

	printArea(newCircle)
	printArea(newRectangle)

}