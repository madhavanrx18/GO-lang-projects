package main

import "fmt"

func add (x int,y int) int{
	return x+y
}

func main(){
	fmt.Println("sum of 4 and 5",add(4,5))
}