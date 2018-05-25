package main

import (
	"bytes"
	"fmt"
)

//*************************** Composing Interfaces

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

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
	/*
		Field "buffer" points to Buffer struct from bytes library containing the following fields

					buf        []byte
				    off        int
				    bootstrap    [64]byte
				    lastRead    readOp

	*/

}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {

	n, err := bwc.buffer.Write(data)
	/*
		buffer struct from inside bwc calls Write and uses the data passed and returns
		n, which is an int, and err, which is = to error and, stores that in buffer
	*/
	if err != nil {

		return 0, nil
		// has to return int and error so we return 0 and nil if there was an error writing
	}

	v := make([]byte, 8)
	//take n int and turns it into a []byte with length of 8 stored in v
	for bwc.buffer.Len() > 8 {
		//	As long as v has more than 8 characters.....
		_, err := bwc.buffer.Read(v)
		//read the  the []byte and return ( n int, err)
		if err != nil {
			return 0, err
			//if there was an error return (0, err)
		}
		_, err = fmt.Println(string(v))
		//print string version of n
		if err != nil {
			return 0, err
		}
	}
	return n, nil

}

func (bwc *BufferedWriterCloser) Close() error {

	/*
		Have to use all the methods of an interface to implement it.

		Method for BufferedWriterCloser now uses both the Close and Write
		of the WriterCloser interface.
	*/

	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		//pullout next 8 characters
		_, err := fmt.Println(string(data))
		//keep flushing the buffer until its empty
		if err != nil {
			return err
		}
	}
	return nil
	// return error value of nothing
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	/*
		this is used as a constructor func to initialize
		a new BufferedWriterCloser object
	*/
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
		//initializes new empty buffer for the new bufferwritercloser object
	}
}

func main() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello everybody, this is a test "))
	//convert string to []byte
	wc.Close()

	//demonstrating type conversion
	r, ok := wc.(*BufferedWriterCloser)
	/*

		Now you can access the fields of BufferWriterCloser, which is buffer
		previous type,

		WriterCloser was unaware of that field

		Type conversion would only work between types that implement the same interface

		Wouldn't work if *BufferedWriterCloser didn't have both a Write() and Close()

	*/
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}
	/*
		the reason we set the conversion to r,ok on line 134 is to catch
		and handle any errors that could result from the conversion
		in order to avoid panic
	*/
}
