package main

import (
	"fmt"
)

//****************************** multi-variable looping

func main() {

	//counter variable, bool tester, incrementer
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		/*
		   i and j set to 0

		   Notice the incrementer syntax format matches the
		   counter variable setting format
		*/
		fmt.Println(i, j)
		//output 0,1,2,3,4 and 0,2,4,6,8
	}
}
