package main

import (
	"fmt"
	"time"
)

//************************** Goroutines and anonymous funcs

func main() {
	var msg = "hello"

	go func() {
		fmt.Println(msg)
	}()

	/*
		even though the msg variable is declared outside
		the go routines' funcs' scope, the go routine
		still catches is due to something called closures.
	*/

	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)

	/*
	  output is Goodbye because go routine doesn't interrupt the main() until it hits the sleep()
	*/
}
