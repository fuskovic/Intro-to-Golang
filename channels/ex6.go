package main

import (
	"fmt"
	"sync"
)

/*
***********For looping through channels to use bools to check if they're open first
*/

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
			/*
			Designated this as a receiving channel by passing into the go routine func.
			Receiving because "<-" points away from chan and into ch.
		
		
			Receiving channels do have the ability
			to check if the sending channel is open
			before receiving
		*/
		for {
			if i, ok := <-ch; ok {
				/*
				 You can return a bool value to check if the channel is open or not before
				 receiving however, you can't do this for 
				*/
				fmt.Println(i)
			} else {
				break
				// break out of loop if channel is closed and nothing is being received
			}
		}
		wg.Done()
		//lets waitgroup know its done
	}(ch)
	go func(ch chan<- int) {
		/*
			Designated this a sending channel as "<-" points into chan
			away from data means we are sending data into the channel
		
			 Sending channels Don't have the ability to check whether
			or not a channel is open before sending data to it
		*/
		ch <- 42
		ch <- 27
		close(ch)
		//when the channel closes the for loop uptop will recognize it and break
		wg.Done()
	}(ch)
	wg.Wait()
}
