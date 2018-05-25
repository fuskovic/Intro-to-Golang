package main

import (
	"fmt"
	"sync"
)

/*
****************Using buffered channels-in case your sender or receiver
needs more time to process but you don't want to block the other side.

NOTE: Order of execution is labeled

*/

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	//add two goroutines to waitgroup
	go func(ch <-chan int) {

		/*
			Designated this as a receiving channel by passing into the go routine func.
			Receiving because "<-" points away from chan and into ch.
		*/

		i := <-ch
		fmt.Println(i)
		//2. 42 is received from channel and printed
		i = <-ch
		fmt.Println(i)
		//4. 27 is received from channel and printed

		// if you did not have two i variables but only one than this would result in deadlock

		wg.Done()
		//5. go routine finished
	}(ch)

	go func(ch chan<- int) {

		/*
			Designated this a sending channel as "<-" points into chan
			away from data means we are sending data into the channel
		*/

		ch <- 42
		//1. 42 sent into channel
		ch <- 27
		//2.  27 sent into channel
		wg.Done()
		//6. go routine finished
	}(ch)
	wg.Wait()
	//main() complete
}
