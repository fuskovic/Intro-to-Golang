package main

import (
	"fmt"
)

//******* Breaking from nested loops using labels and "for ranging over collections"

func main() {

Loop:
	//above is whats called a label. Tells program where to break from nested loop

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				fmt.Println("break to label")
				break Loop
				//this break exits both for loops as label sits outside both
			}
		}
	}

	//for loops and collections- for range loop
	s := []int{1, 2, 3}
	fmt.Println("****Slice*****")
	for k, v := range s {
		//k and v are equal to the index and its value
		fmt.Printf(" index: %v , value: %v\n", k, v)

	}

	//works the same for arrays
	t := [3]int{1, 2, 3}
	fmt.Println("****Array*****")
	for k, v := range t {
		fmt.Printf("index: %v , value: %v\n", k, v)
	}
	//as well as maps
	u := map[string]int{
		"option-A": 1,
		"option-B": 2,
		"option-C": 3,
	}
	fmt.Println("****Map*****")
	for k, v := range u {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

	//You range over strings too
	v := "Faris"
	fmt.Println("*****String*****")
	for k, v := range v {
		fmt.Printf("index: %v, value: %v\n", k, v)
	}
	/*
		what if you don't want the key but just the value?
		If you try to print just one you'll get an error.
		You have to use a write only variable in this case
		so you're not forced to use something you don't need.
	*/
	w := "Faris"
	fmt.Println("*****String(vals only)*****")
	for _, v := range w {
		fmt.Printf("value is: %v\n", v)
	}
}
