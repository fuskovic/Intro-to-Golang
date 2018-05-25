package main

import (
	"fmt"
	"sync"
)

/*
**************************** Ranging over channels
 */
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		/*
			Designated this as a receiving channel by passing into the go routine func.
			Receiving because "<-" points away from chan and into ch.
		*/

		for i := range ch {
			/*
			   you don't need to use two variables to hold what is returned from range
			   like usual when ranging over something with key/value pairs.

			   This is because only the channel message is returned.
			*/
			fmt.Println(i)
		}
		wg.Done()
		//lets waitgroup know its done
	}(ch)

	go func(ch chan<- int) {
		/*
			Designated this a sending channel as "<-" points into chan
			away from data means we are sending data into the channel
		*/
		ch <- 42
		ch <- 27
		ch <- 50
		ch <- 63
		/*
			The channel is not blocked and I have no idea why ranging
			over the channel makes it unaware of the fact that multiple
			messages have been sent to it?????????????????????????????
		*/
		close(ch)
		//notice we are now closing the send channel elegantly as opposed to a "forced shutdown"
		wg.Done()
		//lets waitgroup know its done
	}(ch)
	wg.Wait()
	//main() can't finish until both go routines in the wait group have called done()
}
