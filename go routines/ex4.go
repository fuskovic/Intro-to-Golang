package main

import (
	"fmt"
	"sync"
)

//****************** Synchronizing goroutines with WaitGroups

var wg = sync.WaitGroup{}

func main() {
	var msg = "hello"
	wg.Add(1)
	//Add takes number of goroutines to add
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
		/* 
		Done will decrement waitgroup by one and main() will
		continue as it is 0 by default. Tells the waitgroup
		its routine is done
		*/
	}(msg)
	msg = "Goodbye"
	wg.Wait()
	//main won't finish executing until all the go routines in the waitgroup have called Done()
}
