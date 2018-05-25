package main

import (
	"fmt"
	"sync"
)

//************************** Making Channel data flows more clear

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		/*
			Designated this as a receiving channel by passing into the go routine func.
			Receiving because "<-" points away from chan and into ch.
		*/
		i := <-ch
		//ch on right and variable on left is proper syntax for receiving data
		ch <- 27
		//this would throw an error as we are sending data to a receive only channel
		fmt.Println(i)
		/*

		 */
		wg.Done()
		//lets waitgroup know its done
	}(ch)
	go func(ch chan<- int) {
		/*
			Designated this a sending channel as "<-" points into chan
			away from data means we are sending data into the channel
		*/
		ch <- 42
		//ch on left and data on right is proper syntax for sending data
		wg.Done()
		//lets wait group know its done
	}(ch)
	wg.Wait()
	//main wont finish till wait group goroutines all call done
}

/*
A single go routine can do both reading and writing but, its common
to have different go routines designated for each
*/
