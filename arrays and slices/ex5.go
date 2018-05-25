package main

import (
	"fmt"
)

//******************************* Appending and concatenating slices

func main() {
	/*
		slices don't have to have fixed sizes
		this makes appending and removing indices convenient
	*/
	m := []int{}
	fmt.Println(m)
	fmt.Printf("length of : %v\n", len(m))
	fmt.Printf("Capacity of : %v\n", cap(m))
	//outputs 0,0,0
	m = append(m, 1)
	// takes what you are appending to and what to append to it
	fmt.Println(m)
	fmt.Printf("length of : %v\n", len(m))
	fmt.Printf("Capacity of : %v\n", cap(m))
	//outputs 1,1,1,
	m = append(m, 2, 3, 4, 5)
	//appending multiple indices
	fmt.Println(m)
	fmt.Printf("Length of : %v\n", len(m))
	fmt.Printf("Capacity of: %v\n", cap(m))
	//outputs [1,2,3,4,5] len of 5 cap of 6

	//concatenating slices
	m = append(m, []int{6, 7, 8, 9}...)
	//won't work without the three dots. Spreads slice into individual arguments
	//above syntax is the same as m= append(m, 6,7,8,9)
	fmt.Println(m)
	fmt.Printf("%v\n Length of %v, capacity of: %v\n", m, len(m), cap(m))
	//outputs [1,2,3,4,5,6,7,8,9]
}
