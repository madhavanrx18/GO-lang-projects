package main

import "fmt"

func swap (a,b string) (string , string){
	return b, a
}

func main(){
	a, b := swap("hello", "world")
	fmt.Println("Swapped:" ,a,b)
}
// := declares and initializes a variable