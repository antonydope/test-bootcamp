package main

import (
	"fmt"
	"math"
)
type square struct{
	height float64
	width float64
}
type circle struct{
	radius float64
}
type Shape interface{
	getArea() float64
	getPerimeter() float64
}
func (y  circle) getArea() float64{
	return math.Pi * y.radius*y.radius
}
func (z circle) getPerimeter() float64{
	return math.Pi*z.radius*2
}
func (w square) getArea() float64 {
	return w.height*w.width
}
func (x square) getPerimeter() float64{
	var a float64
	a = x.height+x.width
	return 2*a
}
func geometry(s Shape) {
	fmt.Println(s)
	fmt.Println("the area is: ",s.getArea())
	fmt.Println("the perimeter is: ",s.getPerimeter())
}
func main(){
	newCircle := circle{14}
	newSquare := square{10,20}
	geometry(newCircle)
	geometry(newSquare)
}
/*package main

import (
	"fmt"
	"math"
)

type square struct {
	height int
	width  int
}

type circle struct {
	radius float64
}

// Shape interface
type Shape interface {
	getArea() float64
	getPerimeter() float64
}

func (y circle) getArea() float64 {
	return math.Pi * y.radius * y.radius
}

func (z circle) getPerimeter() float64 {
	return math.Pi * z.radius * 2
}

func (w square) getArea() float64 {
	return float64(w.height * w.width)
}

func (x square) getPerimeter() float64 {
	return float64(2*x.height + 2*x.width)
}

func geometry(s Shape) {
	fmt.Println("Area:", s.getArea())
	fmt.Println("Perimeter:", s.getPerimeter())
}

func main() {
	newCircle := circle{radius: 7}
	newSquare := square{height: 2, width: 4}

	geometry(newCircle)
	geometry(newSquare)
}
*/