package main

import "fmt"

func main(){
num := []string{"tom","cylnton","sharon","adanda","sande",}
fmt.Println(num)
//a := num[5:]
//fmt.Println(a)
//b :=num[:]
//fmt.Println(b)
//c := num[:4]
//fmt.Println(c)
num = append([]string{"antony"}, num...)
fmt.Println(num)
num = append(num, "george")
fmt.Println(num)
num = num[:len(num)-1]
fmt.Println(num)
index := 3 // index where "sam" should be inserted
num = append(num[:index+1], append([]string{"sam"}, num[index+1:]...)...)
fmt.Println(num)
nums := make([]int, 2)
	nums[0], nums[1] = 0, 1
	fmt.Println(nums, len(nums), cap(nums)) // Prints [0 1] 2 2
num[1] = "job"
fmt.Println(num) 
length := len(num) //length of an array or slice can be accessed using the len function.
fmt.Println(length)
//changeFirst(num, "john")
}
/*func changeFirst(num []string, john string){
	if (len(num) > 0){
		num[4] = john
	}
}*/