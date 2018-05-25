package main

//******************************Using iota and common use case example
import (
	"fmt"
)

const (
	_ = iota // Using a write only var as an easy way to get out of zero-indexing

	/* << is a bit shifter

	<< is used for "times 2" and >> is for "divided by 2" - and the number after it is how many times.

	*/
	KB = 1 << (10 * iota) // so this is 1*2......10 times
	MB                    // this is 1*2..........20 times and so forth
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {

	fmt.Printf("%v\n", KB)
	fmt.Printf("%v\n", MB)
	fmt.Printf("%v\n", GB)
	fmt.Printf("%v\n", TB)
	fmt.Printf("%v\n", PB)
	fmt.Printf("%v\n", EB)

	/*
		KB 1024
		MB 1048576
		GB 1073741824
		TB 1099511627776
		PB 1125899906842624
		EB 1152921504606846976
	*/

	filesize := 4000000000.
	//don't forget to close with a "." since we are dealing with decimals in our arithmetic
	fmt.Printf("%.2fGB\n", filesize/GB)
	//formats result by 2 decimal places and writes GB next to it
	//output is 3.73GB because you're doing 4*1000000000/3*(10*3)

}
