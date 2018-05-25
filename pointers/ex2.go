package main

import (
	"fmt"
)

//************* Using pointers, &a, and dereferencers with structs

type myStruct struct {
	foo int
}

func main() {
	//new instance of myStruct
	var ms myStruct
	fmt.Printf(" 'var ms myStruct' is equal to : %v\n", ms)
	///output will be field values of myStruct so {0}

	//instance of myStruct that holds an int with a value of 10
	ms = myStruct{foo: 10}
	fmt.Printf(" 'ms=myStruct{foo: 10}' will print %v\n", ms)
	//output: field values of ms so {10}

	fmt.Printf("  printing '&ms' will output: %v\n", &ms)
	//output will be address of field values so &{10}

	fmt.Println("*************************************") // placed to seperate example types

	var ds *myStruct
	fmt.Printf(" 'var ds *myStruct' outputs: %v\n", ds)
	//output will point to default field values of myStruct so <nil>

	ds = &myStruct{foo: 20}
	fmt.Printf(" ' ds = &myStruct{foo: 20}' will output: %v\n", ds)
	//output will be address of field value of struct or &{20}

	fmt.Printf(" '*ds' will output: %v\n", *ds)
	//output will dereference address of myStruct so {20}

	fmt.Printf(" 'ds.foo' will output: %v\n", ds.foo)
	//output will completely isolate the specific field value of the address of myStruct so 20

	fmt.Println("*************************************") //placed to seperate example types

	var es *myStruct
	fmt.Printf(" 'var es *myStruct' will output: %v\n", es)
	//output will point to default values of myStruct so <nil>

	es = new(myStruct)
	fmt.Printf(" 'es = new(myStruct)' will output: %v\n", es)
	//output will be field values of address of myStruct so &{0}

	es = &myStruct{foo: 50}
	fmt.Printf(" 'es = &myStruct{foo: 50}' will output: %v\n", es)
	//output will be = to the address of the field values of myStruct so &{50}

	fmt.Printf(" '*es' will output: %v\n", *es)
	//output will dereference the address of myStruct so {50}

	fmt.Printf(" 'es.foo' will output: %v\n", es.foo)
	//output will completely isolate the specific field value at the address of myStruct so 50
}
