package main

import (
	"fmt"
	"encoding/json"
)
type Student struct{
	Name string   `json:"Customname"`
	Age int       `json:"studentAge"`
	Race string   `json:"studentrace"`
	Taste string  `json:"studenttaste"`
	Strength int  `json:"studentstrength"`

}
func main(){
	luke := Student{
		Name : "lukeshaw",
		Age : 24,
		Race : "american",
		Taste : "coder",
		Strength : 678,
	}

	data, err :=json.Marshal(luke)
	if err != nil {
		fmt.Println("there was no error")
		return
	}
	fmt.Println(string(data))
}