package main

import "fmt"

// package level
var h,m string = "harish","madhavan"

func main() {
    // function level
	var lang string = "tamil"
	fmt.Println(h,m,lang)
}