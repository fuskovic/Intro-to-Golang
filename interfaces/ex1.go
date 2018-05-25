package main

import (
	
	"fmt"
)

//*********************************************** Basic Interface

func main() {

	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go"))
	/*
		Notice the w variable is of type Writer or, in other words,
		is implementing the Writer interface.
		
		It also an instance is an instance of the ConsoleWriter struct. 
		
		It could have just as easily created an instance of whatever struct it wanted  
		because it implements the Writer interface, whatever struct would then
		access that interfaces methods or behaviors. 
		
		Kind of like polymorphism in Java but not really. We'll see why later.

	*/
	fmt.Println("Hello, playground!")
}
type ConsoleWriter struct {
}

type Writer interface {
	/*
		interfaces store method definitions.

		It is also convention to name single method interfaces with the
		suffix "er", which are very common in Go.
	*/

	Write([]byte) (int, error)
	/*
		so basically anthing that implements
		this interface will contain the above method that
		takes a []byte and returns (int,error)
	*/
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	/*
		This func creates and passes an instance of the
		ConsoleWriter struct as its receiver and also implements
		the writer interface by borrowing its "Write()" method
		or behavior
	*/
	n, err := fmt.Println(string(data))
	return n, err
}