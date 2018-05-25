package main

import (
	"fmt"
)

//******************** Booleans, integar types and basic arithmetic

func main() {
	//booleans are false by default
	n := 1 == 1
	fmt.Printf("%v, %T\n", n, n)
	//output will be true,bool
	m := 1 == 2
	fmt.Printf("%v, %T\n", m, m)
	//output will be false, bool

	/*
		different integar types
		Signed integars
		-int8 -128 to 127
		-int16 -32,768 to 32,767
		-int32 -2,124,483,648 to 2,147,483,647
		-int 64 -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
		Unsigned Integars(only positive)
		uint8 -  0 to 255
		uint16 -  0 to 67,535
		uint32 - 0 to 4,294,967


		byte= 8bit unsigned integar(very common in data stream encoding)
	*/
	//basic arithmetic

	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

}
