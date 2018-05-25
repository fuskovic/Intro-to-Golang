package main

import (
	"fmt"
)

//************************ Order of execution

func main() {

	// Go does not throw exceptions. Only error values.

	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	fmt.Println("end")

	/*Panics happen after defer statements are executed

	ORDER OF EXECUTION

		the func
		defers
		panics
		return values.

	Panics occur when program can't continue at all. Don't use when file can't
	be opened. Use for unrecoverable events like not being able to obtain a TCP port
	for a web server for example.
	*/
}
