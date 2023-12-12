package main

import"fmt"

func main(){
myCar := car{}
myCar.make = "subaru"
myCar.colour = "black"
myCar.wheel = 4
myCar.noOfseats.frontseat = 6
myCar.noOfseats.backseat = 7
fmt.Println(myCar)

}
type car struct {
	make string
	colour string
	wheel int
	noOfseats seats
}
type seats struct{
frontseat int
backseat int


}
