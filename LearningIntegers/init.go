package main

import "fmt"


func main(){
	fmt.Println("hello")


	var myint int16

	for i := 0;i<129; i++ {
		myint += 1
		fmt.Println(myint)
	}
	
}