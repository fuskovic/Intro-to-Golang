package main

import (
	"fmt"
)

//**************************** Syntax for Arrays of Arrays/ multi-dimensional arrays

func main() {
	/*
	 var identityMatrix [3][3]int= [3][3]int{ [3]int{1,0,0}, [3]int{0,1,0}, [3]int{0,0,1} }
	 so Identity matrix holds three indices in which each index holds three indices itself
	*/
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

}
