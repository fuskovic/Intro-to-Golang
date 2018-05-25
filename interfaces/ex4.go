package main

import (
	"fmt"
)

//*********************************** Any custom types with methods can use interfaces

type Incrementer interface {
	Increment() int
}

type IntCounter int

//custom type Intcounter

func (ic *IntCounter) Increment() int {
	/*
	   Create method for the IntCounter type
	   ic points to type IntCounter
	   Incrementer interface is incremented by adding the Increment()
		method. 
	*/
	*ic++
	//points to the ic passed
	return int(*ic)
	
}

func main() {

	myInt := IntCounter(0)

	/*
		Field values for structs are set within {} but, custom
		types like intcounter use ().
		
		like IntCounter(0) above when creating an instance of the IntCounter type.
	*/

	var inc Incrementer = &myInt

	/*
	inc is of type Incrementor so it can use Increment() method

	Value points to &myInt which is 0
	
	*/

	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}
