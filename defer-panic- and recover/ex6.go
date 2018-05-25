package main

import (
	"fmt"
	"log"
)

//************************** Using recover to save application execution
func main() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			/*
				Recover is useful as a deferred function.

				Looks for a panic and if a function is panicking, it will decide
				what to do with it from there.

				If you recover succesfully, current funcs will not continue but,
				higher funcs in the call stack will.
			*/
			log.Println("Error:", err)
			panic(err)
		}
	}()

	panic("somthing bad happened")
	fmt.Println("done panicking")
}

/*
	Notice the }() syntax. Thats what makes the function execute. Defers don't
	take a function themselves, they take a function call.

	So in the above example we create an anonyous function that sets err to the recover(),
	which can return an error itself.

	In the event that err, or in this case recover(), is not nil(it returns
	either nil or the error itself),we print the error causing the program to panic.

	At this point the func is only set and needs the additional () after the defer func
	to actually invoke it.

	This all an attempt to recover before the actual panic.
*/

/*
	ORDER OF EXECUTION

		the func
		defers
		panics
		return values.
*/
