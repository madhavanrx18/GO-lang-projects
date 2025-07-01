package main

import "fmt"

// naked return

func multiply (num int) (num2, num3 int){
	num2 = num*2
	num3 = num*3
	return 
}

func main(){
	fmt.Println(multiply(3))
}