package main

import"fmt"

func main(){
car := []string{"BMW","TOYOTA","LANDCRUISER","SUZUKI"}
for index := 0; index < len(car); index++ {
	if car[index] == "MASSEY" {
		
		continue
	}
	fmt.Println("found you baby!", car[index])
}

}