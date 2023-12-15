package main

import (
	"fmt"
	"encoding/json"
)

type Tuiton struct{
	StudentName string `json:"Student name"`
	StudentAge int   `json:"Student age"`
	StudentHobby string    `json:"Student hobby"`
	StudentBestsubject string   `json:"Student best subject"`
}
func main(){
	student1 := Tuiton{
		StudentName: "Christian",
		StudentAge: 9,
		StudentHobby: "Football",
		StudentBestsubject: "Mathematics",
	}
	student2 := Tuiton{
		StudentName: "Breyden",
		StudentAge: 7,
		StudentHobby: "studying",
		StudentBestsubject: "all subjects",
	}
	student3 := Tuiton{
		StudentName: "James",
		StudentAge: 10,
		StudentHobby: "cycling",
		StudentBestsubject: "English",
	}
	students := []Tuiton{student1,student2,student3}
	data , err := json.Marshal(students)
	if err != nil {
		fmt.Println("there was an error while marshalling:",err)
		return
	}
	// Unmarshal the JSON array into a slice of structs
	var formattedStudents []Tuiton
	err = json.Unmarshal(data, &formattedStudents)
	if err != nil {
		fmt.Println("there was an error while unmarshalling:", err)
		return
	}

	// Print each student's information on a separate line
	for _, studentInfo := range formattedStudents {
		fmt.Printf("StudentName: %s,\nStudentAge: %d,\nStudentHobby: %s,\nStudentBestsubject: %s\n\n",
			studentInfo.StudentName, studentInfo.StudentAge, studentInfo.StudentHobby, studentInfo.StudentBestsubject)
	}
	// unmarshalling
	/*var yoStudent1 Tuiton
	err = json.Unmarshal(data, &yoStudent1)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("THE UNMARSHALLED STRUCT")
	fmt.Printf("StudentName: %s,\nStudentAge: %d,\nStudentHobby: %s,\nStudentBestsubject: %s",yoStudent1.StudentName,yoStudent1.StudentAge,yoStudent1.StudentHobby,yoStudent1.StudentBestsubject)
*/}