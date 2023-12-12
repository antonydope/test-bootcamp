package main

import "fmt"

func main(){
var t string
var p *string
t = "awesome"
p = &t
fmt.Println(p)
fmt.Println(t)
fmt.Println(*p)

x := "am awesome yo!"
y := &x
fmt.Println(y)
fmt.Println(*y)
changeValue(y)
fmt.Println("updated x: ",x)
updateV(p)
fmt.Println("the upadated value: ",t)

}
func changeValue(x *string){
	*x = "thank you!"
}
func updateV(p *string){
	*p = "am glad"
}