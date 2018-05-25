package main

/*
******************************Understanding defer and LIFO

Note: order of exeuction is labeled
*/
import (
	"fmt"
)

func main() {
	fmt.Println("start")
	//1

	defer fmt.Println("middle")
	//6

	fmt.Println("end")
	//2

	/*
		Defer is used to postpone something
		until the entire func main() has executed but before it has returned.
	*/

	again()
	// Notice how the deferred func() in main prints after the LIFO printing of the again() is complete
	//LIFO is "last in first out"....see output below on lines 48-55 as well as labeled order
}

func again() {

	defer fmt.Println("second start")
	//5

	defer fmt.Println("second middle")
	//4

	defer fmt.Println("second  end")
	//3

}

/*

output is ....

start
end
second  end
second middle
second start
middle


*/
