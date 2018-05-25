package main

import (
	"fmt"
)

//************************* Return values and passing slices

func main() {

	s := sum("The sum is : ", 1, 2, 3, 4, 5)
	//s is int type due to return type of sum()
	fmt.Println(*s)
	//dereference the underlying data that sum returned....see lines 46-48
	d, err := divide(3.0, 5.0)
	if err != nil {
		fmt.Println(err)
		//divide returns an error value that must be handled
		return
	}
	fmt.Println(d)
}

func sum(msg string, values ...int) *int {

	/*

		"values...int)" Means any ints entered for values will be placed into a []int

		Notice return value is unnamed. Named return value might be (result int)

		funcs can also have multiple return values.
		An example might be (result int, error)

		pointer on return value of int type.....see lines 46-48

	*/

	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
	/*
		Return address of result. This is done to point to underlying data
		as opposed to creating a copy of it at execution time.

		Remember that a functions execution stack is cleared after execution.

		This way if you use pointers, address of, and dereference you can promote that return
		value to the heap(shared memory of the program)
	*/

	//fmt.Println("The sum is: ", result)
}

func divide(a, b float64) (float64, error) {
	if a == 0.0 || b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
		/*
			error checking done here. Think of this as a make shift try/catch
			we execute code based on value of error
		*/
	} else {
		return a / b, nil
		//error returned as nil
	}
}
