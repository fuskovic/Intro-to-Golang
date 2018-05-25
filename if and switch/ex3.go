package main

import (
	"fmt"
)

//****************** Type switching and type switching w/ Arrays

func main() {
	/*
		type switching is a form of type validation.

		An interface can contain any "type" of data.

		In this case an int.

		Using i.(type) returns int and is then validated accordingly.
	*/
	var i interface{} = 1

	switch i.(type) {

	case int:
		fmt.Println("i is an int")
		//this is what will output
	case float64:
		fmt.Println("i is a float 64")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is another type")
	}

	//type switching arrays

	var k interface{} = [3]int{}
	//arrays must contain same length and data type to be evaluated as same type
	switch k.(type) {

	case int:
		fmt.Println("k is an int")
	case float64:
		fmt.Println("k is a float 64")
	case string:
		fmt.Println("k is a string")
	case [3]int:
		//below is the case that gets output
		fmt.Println("k is an array of three ints")
		break
		fmt.Println("I want to print too")
		/*
			this does not print because we used break to leave the case early
			could be necessary in some situations
		*/
	default:
		fmt.Println("k is another type")
	}
}
