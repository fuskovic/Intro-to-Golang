package main

import (
	"fmt"
)

//*********************************** Anonymous funcs and funcs as types

func main() {
	func() {
		for i := 0; i < 5; i++ {
			func(i int) {
				/*
					this is an anonymous func nested inside another anonymous func
					be default, the i variable would only be available to the scope
					of the anonymous function and not to i.
				*/
				fmt.Println(i)
			}(i)
			//however, by passing i into its invoke call, it becomes available to its parent func
		}
	}()

	//funcs as types

	func() {
		var f = func() {
			/*
				Notice f is of type func.
				Also notice short var declaration is not used even tho f
				is nested inside another func.
			*/
			fmt.Println("Hello go")
		}
		f()
		//f is invoked
	}()
	//func f is neted in is invoked

	func() {
		var divide func(float64, float64) (float64, error)
		//var divide is declared with type func, required args and return variables
		divide = func(a, b float64) (float64, error) {

			if b == 0.0 {
				//little bit of error catching logic
				return 0.0, fmt.Errorf("Cannot divide by zero")
			} else {
				return a / b, nil
				//if there is no error return the quotient and nil for return error value
			}
		}

		d, err := divide(5.0, 3.0)

		//because divide returns an error value, we need to handle it
		if err != nil {
			fmt.Println(err)
			return
			//handle the error here if there is one
		}
		fmt.Println(d)

	}()
}
/*      NOTE:
		If you're using variables with func types make sure they are defined before using them.
		*/
