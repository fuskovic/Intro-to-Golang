package main

import (
	"fmt"
)

//******************** How values are passed into deferred functions

func main() {

	a := "start"
	defer fmt.Println(a)
	/*
		In this case, "end" value gets ignored and "start" is printed.
		This occurs because the defer retains the value of the function
		when the defer was called.
	*/
	a = "end"
}
