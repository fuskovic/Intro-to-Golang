package main

import (
	"fmt"
)

//********************************* Complex Integer types, strings, and runes

func main() {
	/*
		complex integar types(imaginary nums)
		complex64
		complex128
	*/

	var g complex128 = 1 + 2i
	var f complex128 = 2 + 5.2i
	fmt.Println(g + f)
	fmt.Println(g - f)
	fmt.Println(g * f)
	fmt.Println(g / f)
	//isolating real and imaginary numbers
	fmt.Printf("%v, %T\n", real(g), real(g))
	fmt.Printf("%v, %T\n", imag(f), imag(f))
	//output of last two lines will read 1, float 64 and 5.2,float 64
	var h complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", h, h)

	//text types -strings are UTF-8 encoded
	s := "this is a string"
	fmt.Printf("%v, %T\n", s[2], s[4])
	//strings are treated as []bytes so output is 105, uint8
	fmt.Printf("%v, %T\n", s, s)
	//convert to bytes
	bites := []byte(s)
	fmt.Printf("%v, %T\n", bites, bites)
	// will output the array of numerical bytes

	//rune string types- same thing as int32
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
	//If necessary, there are shpecial functions available to decode utf-32 like readRune()

}
