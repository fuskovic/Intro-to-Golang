package main

import (
	"fmt"
)

//******************* Basics of pointers, address of, and dereferencing
var a int
var b int

func main() {

	var a int = 42
	fmt.Println(a)
	//output is 42
	var b *int = &a
	fmt.Println(a, b)

	/*
		output:42 0xc420090008
		var b points to an int type and its value points
		to the address of a in memory
	*/

	fmt.Println(&a, b)

	//output: 0xc420090008 0xc420090008

	/*
		in order to get the value at that particular address in memory
		you must use dereferencing
	*/

	fmt.Println(a, *b)
	/*
		output: 42 42, now these variables are tied together
		now because the b var is pointing to an int type at the
		address in memory of a and then, dereferencing to return whatever
		value is stored at that address in mem. So no matter what
		we change a to now, *b == a.
	*/

	a = 27
	fmt.Println(a, *b)

}
