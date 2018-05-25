package main

import (
	"fmt"
	"sync"
)

//************************************ Understanding Channel deadlock

var wg = sync.WaitGroup{}

func main() {

	ch := make(chan int)

	go func() {
		i := <-ch
		//receive from channel and hold in "i"
		fmt.Println(i)
		wg.Done()
		//lets waitgroup know its done
	}()

	for j := 0; j < 5; j++ {

		wg.Add(2)
		//go routines added to waitgroup but only one inside for loop
		go func() {
			ch <- 42
			/*
			 Keep in mind that only one message can be in the channel at a time
			 The message that was pushed to the i variable that printed
			 never left the channel. Therefore, this example will result in
			 a deadlock upon second interation of the for loop.
			*/
			wg.Done()
			//lets waitgroup know its done
		}()
	}
	wg.Wait()
	//main() will not complete until both goroutines are done
}
