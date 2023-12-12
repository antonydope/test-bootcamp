package main

import "fmt"

func main(){

type Student struct{
fname string
sname string
age int 
colour string
}
stu1 := Student{"anton","musumba", 56,"brown",}
fmt.Println(stu1)
stu2 := Student{"faith","cherono",47,"black"}
fmt.Println(stu2)

}