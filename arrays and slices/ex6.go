package main

import (
	"fmt"
)

//*****************Slices-removing inner indices by appending

func main() {

	n := []int{1, 2, 3, 4, 5, 6}
	//index
	//pos.     0  1  2  3  4  5
	o := append(n[:2], n[3:]...)
	//take slice up to pos two (without including pos 2) and append everything from pos 3(including pos 3) onward
	fmt.Printf("%v\n", o)
	/*
		outputs [1,2,4,5,6]

		remember all these slices point to an underlying array
		watch what happened to n
	*/
	fmt.Printf("%v\n", n)
	/*
		outputs [1,2,4,5,6,6], this unexpected behavior can wreak havoc
		NOTE:always remember not to reuse slices after operations like the one above
	*/

}
