package main

import (
	"fmt"
)

//************************** Creating a copy Vs pointing to address of

func main() {

	//a is an array that holds however many values you give it
	a := [...]int{1, 2, 3}
	//b is a copy of a in its original state
	b := a
	//index 2 is changed to 5
	b[1] = 5
	fmt.Println(a)
	//above outputs 1,2,3
	fmt.Println(b)
	/*
		above outputs 1,5,3

		in order for b to print the changed "a" array ,it would need to point to &a
		as demsontrated below by var c
	*/
	c := &a
	fmt.Println(c)
	//outputs 1,2,4
	//capacity function reveals maximum amount of indices the array can hold
	fmt.Printf("Capacity of C: %v\n", cap(c))
}
