package main

import (
	"fmt"
)

//************************************ Methods

type greeter struct {
	greeting string
	name     string
}

func main() {
	ex()
}

func (g *greeter) greet() {

	/*	pointer creates AND passes an instance of greeter

		This is what makes this a method.

		In this case the greeter struct is called the receiver. g* greeter is kind of
		like "this" in Java.
	*/
	fmt.Println(g.greeting, g.name)
	//prints field values that belong to the g greeter struct
}

func ex() {
	//function to set g's field values and execute its greet method
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}

	g.greet()
}

/*

NOTE: you can create methods out of any instatiated types so if you had
type counter int{} instead of type greeter struct{},
you could pass that into func(g greeter) just the same
in order to access that types fields

*/
