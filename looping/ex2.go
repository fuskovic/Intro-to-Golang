package main

import (
	"fmt"
)

//************************* Leaving out the counter variable

func main() {
	g := 0
	for ; g < 5; g++ {
		/*
			whitespace semi-colon means the loop will fill in the variable implicitly
			based on the other statements.

			In this example g is scoped to the main func
			where the first example i and j are scoped within their statement block
		*/
		fmt.Println(g)
	}
	fmt.Println(g)
	//outputs final value of g. Which did not get printed within the loop due to failing bool test

}
