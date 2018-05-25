package main

import (
	"fmt"
)

//************************************ Anonymous structs and using pointers with structs

func main() {

	/*
		structs also do not have to have names.
		Referred to as anonymous structs.

		Also, structs are kind of like arrays.

		when one struct variable holds tanother like line 22 it creates a copy of the original.
		Modifying the original or the copy afterwards does not change the other.
	*/

	someThing := struct{ name string }{name: "john pertwee"}
	anotherSomething := someThing
	//copy of someThing created and changed field value name
	anotherSomething.name = "TomBaker"
	fmt.Println(someThing)
	fmt.Println(anotherSomething)
	/*
		something still outputs john pertwee
		anotherSomething outputs tomBaker
	*/

	someThing.name = "Faris Huskovic"
	fmt.Println(someThing)
	fmt.Println(anotherSomething)
	//something.name now outputs Faris Huskovic
	//anotherSomerhing still outputs Tom Baker

	//in order to tie their values, we must point to the address of underlying data

	person := struct{ name string }{name: "Faris"}
	person2 := &person
	person.name = "bowser"
	fmt.Println(person2)
	/*
		now person2 will reflect any changes to person1
		output is &{bowser}
		to isolate, just dereference
	*/
	fmt.Println(*person2)
	//output is {bowser}
}
