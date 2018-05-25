package main

import (
	"fmt"
)

//*********** Underlying value relationships of arrays and slices

func main() {

	g := [3]int{1, 2, 3}
	h := g
	fmt.Println(g, h)
	g[1] = 42
	fmt.Println(g, h)
	//first line will print identical arrays so output will be [1,2,3],[1,2,3]
	/*
		second line will print g which has index position one changed to 42 so [1,42,3] and
		h will still print [1,2,3] because it is a created copy of the first values of g.

		Also notice g holds the length of the array as well as the values.
		however,if you remove the number of indices in the first g declaration
		than the behavior changes. Remember the only time you are dealing with an array,
		as opposed to a slice, is when you set the size.
	*/

	i := []int{1, 2, 3}
	j := i
	fmt.Println(i, j)
	i[1] = 42
	fmt.Println(i, j)

	/*
		Output is 1,42,3 for both because slices point to the underlying data by default
		So don't pass arrays. Pass slices instead.
	*/
}
