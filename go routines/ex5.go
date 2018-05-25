package main

import (
	"fmt"
	"sync"
)

//****************************************Synchronizing using mutex locks to run in parallel

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

/*
 Simple mutex's are locked or unlocked. If something tries to
 manipulate a locked value, it has to wait until its unlocked.


 In an RWMutex You can have an infinite number of
 readers but only one can write at a time. 
 
 Also, if anything is reading,
 we can't write to it until the readers are done. 
 
 Nothing can read or write
 until a writer is done.

 Execution is labeled in comments in sequential order
*/
func main() {

	for i := 0; i < 10; i++ {
		wg.Add(2)
		// 1. add two goroutines below to waitgroup
		m.RLock()
		/* 2. Read lock is placed on go routine with sayHello() to prevent 
		other go routines from reading any variables in it.
		*/
		go sayHello()

		m.Lock()
		/* 5. This go routine locks any variables in its increment() func to
		prevent other go routines from reading any variables in it
		*/
		go increment()
		
		//these locks are necessary in order to print the incremented counter variable in sequential order
	}
	wg.Wait()
	//main won't finish executing until all go routines in wait group are done
}

func sayHello() {
	// 3. this func locks the counter variable in order to print it before any other go routines do anything else to it
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	//4. go routine is unlocked once counter is printed this way counter can't be incremented until its printed first
	wg.Done()
}

func increment() {
	//6. this func locks the counter variable in order to increment it before any other go routines do anything else to it
	counter++
	//7. now counter can increment since the it was unlocked on line 56
	m.Unlock()
	//8. now counter is completely unlocked until line 39 locks it again on next interation of for loop
	wg.Done()
}
