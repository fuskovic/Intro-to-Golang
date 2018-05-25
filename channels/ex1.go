package main

import (
	"fmt"
	"sync"
)

//*********************** Using sending and receiving channels

var wg = sync.WaitGroup{}

//waitgroups are used to moderate goroutines
func main() {
	ch := make(chan int)
	/*
		have to use make function to create channel
		channels only work with the data type you declare
	*/
	wg.Add(2)
	//two go routines added to wait group
	go func() {
		/*
			this is the receiving go routine channel because the syntax below shows ch on the right with arrow pointing
			away from it to the var so i holds what is received from the channel
		*/
		i := <-ch
		fmt.Println(i)
		wg.Done()
		//this is a way of the go routine letting the wait group know its done
	}()
	go func() {
		/*
			this is the sending go routine channel because ch is on the left and arrow is pointing to it
			so we are sending data to the channel
		*/
		i := 42
		ch <- i
		i = 27
		wg.Done()
	}()
	wg.Wait()
	//main() will not complete until both goroutines are done
}
