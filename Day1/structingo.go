package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// always variable name datatype

func (p *Person) greetings() {

	fmt.Println("Your name", p.Name)
	fmt.Println("Your age", p.Age)

}

func main() {
	var dood = new(Person)
	dood.Name = "HM"
	dood.Age = 20
	dood.greetings()
}
