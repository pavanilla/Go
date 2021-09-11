package main

import "fmt"

//engineer type --> contains name and age of the engineer
// here engineer is the user defined data type
type Engineer struct {
	age  int
	name string
}

func (e *Engineer) UpdateAge() { // any return tyep
	e.age += 1 // no need of return type as we are updating here using pointers
}

func (e *Engineer) UpdateName() {
	e.name = "teja"
}

func main() {

	//pavan is the pointer to the type engineer
	pavan := &Engineer{
		age:  21,
		name: "pavan",
	}
	fmt.Println(pavan)
	pavan.UpdateAge()
	pavan.UpdateName()
	fmt.Println(pavan)
}
