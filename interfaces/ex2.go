package main

import (
	"fmt"
)

//************************** Pointer Receiver VS Value Receiver

type myWriterCloser struct {
}

func (mwc myWriterCloser) Write(data []byte) (int, error) {
	/*
	 methods need to have receivers, in the case the receiver is
	 a value receiver.

	 A pointer reciever would look like func(mwc *myWriterCloser)
	*/
	return 0, nil
}
func (mwc myWriterCloser) Close() error {
	return nil
}

type Writer interface {
	/*
		interfaces describe behaviors by storing method definitions.
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

type Closer interface {
	Close() error
}

type WriterCloser interface {
	/*
	 composed of both writer and closer interfaces

	 Kind of like embedding with structs but for interfaces
	 and its called composing instead
	*/
	Writer
	Closer

	/*
	   So a struct would implement WriterCloser if it has the
	   Write([]byte) (int,error) and Close() error
	   methods on it.
	*/
}

func main() {
	var wc WriterCloser = myWriterCloser{}
	fmt.Println(wc)
}

	/*
		Notice pointer receiver is not used (*myWriterCloser{}).

		When you define types and assign methods to them. Each one of those types has a method set.

		When working with types directly, the method set is all of the methods associated with that type,
		regardless of the receiver type.

		Interfaces are a little different.

		When talking in the context of an interface, is any method that has a value receiver.

		In this case, wc holds the value of myWriterCloser{}

		So if Write() had a pointer receiver, it would not be implemented because its method set
		would be incomplete......unless, var wc WriterCloser = &myWriterCloser

		The reason for that is the method set for a pointer is the sum of all of the value and
		pointer receiver methods.

		When you implement an interface with a value type, the methods that implement the interface,
		have to all have value receivers. If you implement an interface with a pointer, then you just
		have to have the methods there, regardless of receiver type.

		So basically var wc WriterCloser = &myWriterCloser{} would be more flexible to work with
		because if you had func(mwc *myWriterCloser) Write...... and func( mwc myWriterCloser) Close
		.....it could still implement the interface even though the methods contain different types
		of receivers. 
		
		Where as in its current state( var wv WriterCloser=myWriterCloser{}) needs
		the Write() and Close() methods to each contain value receivers like they currently do.

		You want to use pointer receivers when you need access to the underlying data
	*/
