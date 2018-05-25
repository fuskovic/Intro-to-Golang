package main

import (
	"fmt"
)

//***************** Creating slices with make function

func main() {

	/*
		Takes two arguments. Type of slice to create and length.
		capacity is actually a third optional argument
	*/

	p := make([]int, 3, 5)
	fmt.Println(p)
	fmt.Printf("length is : %v\n", len(p))
	fmt.Printf("capacity is : %v\n", cap(p))
	/*
		outputs 0, 3 ,5

			means a p var was set to make a []int
			p holds 3 indices currently
			p can hold up to 5
			because indices were not set, they equal 0 by default
	*/
}
