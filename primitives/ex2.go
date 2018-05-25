package main

import (
	"fmt"
)

//********************************* Arithmetic between different int types and bit operators

func main() {
	//arithmetic between different int types must be converted
	var x int = 10
	var y int8 = 3
	fmt.Println(x + int(y))

	//use below binary for bit operators examples
	x = 10 //1010
	var z int
	z = 3 // 0011

	//bit operators-checks based comparisons or mismatches of equality using bit positions

	//and looks for which bit position is set on both variables
	fmt.Println(x & z) //0010
	//or looks for whether one or the other has that bit position set. If they both have it set it will still record it.
	fmt.Println(x | z) //1011
	//exclusive or looks for only one to have bit position set. If both do its "0".
	fmt.Println(x ^ z) //1001
	//and not- looks for when they both do not have the bit set
	fmt.Println(x &^ z) // 0100

	/*
		floating point types- decimals numbers
		float32
		 which means (pos or neg)1.18*10^-38)
		 to
		 ((pos or neg)3.4*10^38)

		float64
		which means
		(pos or neg)2.23*10^308)
		to
		(pos or neg)1.8E^308)
	*/

	var q float64 = 3.14
	q = 13.7e72 //  means 13.7 * 10^72
	q = 2.1E14  // means 2.1 * 10^14
	fmt.Printf("%v, %T\n", q, q)

	//output is 2.1E14, float64

	/*
		arithmetic can be done between floating
		points long as they are both in fact
		floating points.

		NO remainder operation can be used
	*/

	var p float64 = 2.78
	fmt.Println(p + q)
	fmt.Println(p - q)
	fmt.Println(p * q)
	fmt.Println(p / q)
}
