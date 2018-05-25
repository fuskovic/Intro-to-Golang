package main

import (
	"fmt"
)

//****************** Type switching empty interfaces to test types

func main() {
	ts()
}

func ts() {
	var i interface{} = 0

	// the below is typically done to verify the type you are dealing with....
	switch i.(type) {

	case int:
		fmt.Println("i is an integar")

	case string:
		fmt.Println(" i is a string")

	default:
		fmt.Println("I don't know what i is ")
	}
}

/*
	then you would add your following logic here to continue after you
	have made sure of the empty interface's type.
/*
	***************************************BEST PRACTICES FOR INTERFACES***********************

	The smaller your interfaces the better.

	If you need a robust interface, it's better to compose a bunch of smaller ones togethor.

	3 of the most powerful interfaces within the standard libraries
		io.Writer
		io.Reader
		interface{}

	Don't export interfaces for types that will be consumed.

	Do export interfaces for types that will be used by package

	Design funcs and methods to receive interfaces whenever possible

*/
