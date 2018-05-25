package main

import (
	"fmt"
	"strconv"
)

//********************Variables- declaring and converting them

//this is a block of variables at package level
var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sara Jane Smith"
	doctorNumner int    = 3
	season       int    = 11
)

var i int = 27

/*
Understanding scopes- 3 types in go
-lower case at package level are consumed in the package(any file in same package can access)
-Upper case at package level is global(accessible outside of package)
-block scope(inside function) only visible within the function
*/

/*
Rules on declaring variables
-length of variable name should represent life of variable(camelCase)
-Acronyms (like HTTP) should be all caps
-short as possible
*/

func main() {
	//ways to declare a variable
	//set value later
	var foo int
	var bar int = 42
	day := 42

	fmt.Printf("%d, %d, %d\n", foo, bar, day)

	//Understanding type conversion through variables
	//destinationType(variable) like( int(i) where i="42" originally)
	//strconv is for strings
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)
	//output is "42, int"
	var j string
	j = string(i)
	//output is "*, string"
	fmt.Printf("%v, %T\n", j, j)
	//converting int to string
	var k string
	k = strconv.Itoa(i)
	fmt.Printf("%v, %T", i, k)
	//output will be "42, string"
}
