package main

import (
	"encoding/json"
	"log"
	//"fmt"
)

type Person struct {
	Name  string `json:"CustomName"`
	Age   int    `json:"age,omitempty"` // omits when empty
	//Email string `json:"-"` // the "-" escapes 
}

func main() {
	// Encoding JSON
	//person := Person{Name: "John Doe", Age: 30, Email: "john@example.com"}
	p1 := Person{Name: "antony"}
	//encoded, err := json.Marshal(person)
	pBytes , err := json.Marshal(p1)
	/*if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println("Encoded JSON:", string(encoded))*/
	//log.Println(err)
	//log.Println(string(pBytes))
	//fmt.Println(err)
	//fmt.Println(string(pBytes))

	// Decoding JSON
	var decoded Person
	err = json.Unmarshal(pBytes, &decoded)
	/*if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("Decoded Person:", decoded)
*/
log.Println(err)
log.Println(decoded)
}