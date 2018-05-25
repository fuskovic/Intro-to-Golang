package main

import (
	"fmt"
)

//******************************* Slices- their indices and bounds

func main() {
	//slice bounds

	h := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//index
	//pos.     0  1  2  3  4  5  6  7  8  9

	i := h[:]   //slice of all elements
	j := h[3:]  //slice of all elements from position 3 ( including position 3 )
	k := h[:6]  //slice of all elements up to position 6( without including position 6 )
	l := h[3:6] //slice of elements from position 3 (including position 4) up to position 6( without including position 6)
	fmt.Printf(" original array: %v\n", h)
	//outputs entire slice
	fmt.Printf(" h[:] is equal to : %v\n", i)
	//outputs entire slice
	fmt.Printf(" h[3:] is equal to: %v\n", j)
	//output is 4,5,6,7,8,9,10
	fmt.Printf(" h[:6] is equal to: %v\n", k)
	//output is 1,2,3,4,5,6
	fmt.Printf(" h[3:6] is equal to: %v\n", l)
	//output is 4,5,6

}
