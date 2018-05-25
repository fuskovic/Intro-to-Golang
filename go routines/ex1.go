package main

import (
	"fmt"
	"time"
)

//**************************Using a go routine to invoke a named function
func main() {

	go sayHello()
	/*
	pre fixing functions with go will cause them to run in seperate threads
	like multi-threading in Java
	*/
	time.Sleep(100 * time.Millisecond)

	/* 
	The delay is necessary so thet the go routine executes first before main(). 
	
	Using time.Sleep() is considered bad practice to acheive this effect.

	Had to use it for this simple example
	
	*/

}

func sayHello() {
	fmt.Println("Hello")
}
