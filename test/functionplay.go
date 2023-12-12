package main

import( "fmt"
"math"
)
func sayGreeting(n string) {
	fmt.Printf("Goodmorning my friend %v \n", n)
}
func sayHello(n string) {
	fmt.Printf("Hello my long lost friend: %v from way back in campus \n ", n)
}
func triumphant(n []string, f func(string)){ // takes in a slice of the type string and a function
for _, v := range n { // v is each element in the slice n
	f(v) // f is the name of the function
}
}
func circleArea(r float64)float64{
	return math.Pi *r*r
}
func main(){
	//sayGreeting("shadrack")
	//sayHello("antony musumba the baddest")
//triumphant([]string{"jon","joel","jane","jack"},sayHello)
//triumphant([]string{"jon","joel","jane","jack"},sayGreeting)
a1 := circleArea(47)
a2 := circleArea(21)
fmt.Printf("the area of circle a1 is %0.2f \n",a1)
fmt.Printf("the area of circle a2 is %0.2f \n",a2)
}


