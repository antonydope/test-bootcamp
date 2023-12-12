package main

import "fmt"

func main(){
admissionNumber := map[string]int{
"John omondi": 29,
"Sarah aketch": 45,
"Daisy atieno": 35,
"Amon oyoo": 56,
"Jacob oluoch": 46,
}
fmt.Println(admissionNumber)
delete(admissionNumber,"Daisy atieno")
fmt.Println(admissionNumber)
fmt.Println("the admission  number for amon is: ",admissionNumber["Amon oyoo"])

// looping a map
for k, v := range admissionNumber {
	fmt.Println(k,"-",v)
}
parentsPhone := map[int]string{
	72367789 : "John omondi",
	73465647 : "Sarah aketch",
	87677657 : "Daisy atieno",
	993747873 : "Amon oyoo",
	97646457 :"Jacob oluoch",
}
for k,v := range parentsPhone{
	fmt.Println(k," - ",v)
	
}
fmt.Println("the parents ",parentsPhone[87677657])
}