package main

import (
	"fmt"
)

//*****leaving out the incrementer statement (do while effect), infinite loop, and continue

func main() {
	g := 0
	for g < 5 {

		/*
			this is like a do while loop in Java. Doesn't incremenet.
			As long as logical test passes loop continues
		*/

		fmt.Println(g)
		break
		//you could optionally increment inside here with "g++"" if you want
	}

	//infinite looping
	for {

		fmt.Println(g)
		g++

		if g == 100 {
			break
		}
	}

	//continue statement can help determine if you want to process a record or not

	for i := 0; i < 10; i++ {
		//We only want to print iterations of i that do not have a remainder 0 when /2
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	//output is only odd numbers 1,3,5,7,9

}
