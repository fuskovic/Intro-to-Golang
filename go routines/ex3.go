package main

import (
	"fmt"
	"time"
)

//**************************Passing data into a go routine

func main() {
	var msg = "hello"

	go func(msg string) {
		//pass the variable into the func
		fmt.Println(msg)
	}(msg)

	/*
		Now that msg is passed directly into the go routine, this decouples
		the msg variable from the main func and the go routine.
	*/

	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)

	/*
		As the execution stack of main was running, it passed a copy
		of what msg's value was at the time the goroutine was hit.

		Which was "hello". Thus, output is "hello"

	*/
}
