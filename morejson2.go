package main

import (
	"encoding/json"
	"fmt"
	
)
 type Car struct{
	Make string `json:"carMake"`
	Colour string `json:"carColour"`
	Numberplate int `json:"carNumberplate"`

 }
 func main(){
	car1 := Car{
		Make : "toyota",
		Colour : "black",
		Numberplate : 4765,
	}
	data, err := json.Marshal(car1)
	if err != nil {
		fmt.Println("there was an error",err)
		return
	}
	fmt.Println(string(data))

	var yoCar Car
	err = json.Unmarshal(data, &yoCar)
	if err != nil {
		fmt.Println("the error during unmarshalling",err)
		return
	}
	fmt.Println("UNMARSHALLED STRUCT ")
	fmt.Printf("Make: %s,\nColour: %s,\nNumberplate: %d,\n",yoCar.Make,yoCar.Colour,yoCar.Numberplate)
 }
