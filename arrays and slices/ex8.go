package main

import (
	"fmt"
)

//************************Differences between Arrays and Slices

func main() {

	//********************************** Arrays
	d := [3]int{1, 2, 3}
	//notice e points to d's address in memory
	e := d
	d[1] = 5
	//index position 1 changed after e is declared
	fmt.Printf("output of d is : %v\n", d)
	fmt.Printf("length: %v\n", len(d))
	fmt.Printf("capacity: %v\n", cap(d))
	// d outputs 1 ,5, 3
	fmt.Printf(" output of e is %v\n", e)
	fmt.Printf("length: %v\n", len(e))
	fmt.Printf("capacity: %v\n", cap(e))
	//e outputs  1, 2, 3 because its value holds a copy of d that it received at its time of creation

	/*
		if you did not point to address of and dereference,
		then you would be pointing to d's original state
		as shown in the example below
	*/

	z := &d
	fmt.Printf(" output of z: %v\n", *z)
	fmt.Printf("length: %v\n", len(*z))
	fmt.Printf("capacity: %v\n", cap(*z))
	//output of z will now be equal to output of d

	//*********************************** Slices
	f := []int{1, 2, 3}
	g := f
	f[1] = 5
	//index of position 1 changed after g is declared

	fmt.Printf("output of f: %v\n", f)
	fmt.Printf("length: %v\n", len(f))
	fmt.Printf("capacity: %v\n", cap(f))
	fmt.Printf("output of g: %v\n", g)
	fmt.Printf("length: %v\n", len(g))
	fmt.Printf("capacity: %v\n", cap(g))

	/*
		slices do not need that type of pointer logic because they are
		pointing to the underlying data by default.

		Try to use slices more often than arrays. If Arrays are necessary, be conscious of
		using pointers and dereferencers with your array.
	*/
}
