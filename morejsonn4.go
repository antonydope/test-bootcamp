package main

import (
	"fmt"
	"encoding/json"
	"log"
)
type Friends struct {
	FirstName string `json:"Firstname"`
	SecondName string `json:"Secondname"`
	Age int  `json:"age"`
	Hobby string `json:"Hobby"`
}

func main(){
	firstFriend := Friends{
		"antony","musumba",16,"coding",
	}
	secondFriend := Friends{
		"John","kiptoo",34,"driving",
	}
	thirdFriend := Friends{
		"kevin","kipchumba",23,"drinking",
	}
	friends := []Friends{firstFriend,secondFriend,thirdFriend,}

	data,err :=json.Marshal(friends)
	if err != nil {
		log.Println("There was an error during marshalling:",err)
		return
	}
	fmt.Println(string(data))

	//un-marshalling
	var friendsInfo []Friends
	err = json.Unmarshal(data,&friendsInfo)
	if err != nil {
		log.Println("there was an error during unmarshalling: ",err)
		return
	}
	fmt.Println("THE UNMARSHALLED FRIENDS DATA")
	for _, y := range friendsInfo {
		fmt.Printf("FirstName  : %s,\nSecondName : %s,\nAge        : %d,\nHobby      : %s\n  \n",y.FirstName,y.SecondName,y.Age,y.Hobby)
	}


}