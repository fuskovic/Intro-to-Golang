package main

import (
	"fmt"
)

//******************************** Scope and passing multiple parameters

func main() {
	//upper case is public scope from package and lowercase is internal to the package
	greeting := "hello"
	name := "stacy"

	for i := 0; i < 5; i++ {
		sayMessage("Hello Go", i)
	}

	sayGreeting(&greeting, &name)
	fmt.Println(name)
}
func sayMessage(msg string, idx int) {
	//takes string an int. See them passed on line 11.
	fmt.Printf(" message is: %v and index is: %v\n", msg, idx)
}

func sayGreeting(greeting, name *string) {
	//can use the above syntax to let the compiler infer greeting and name are both strings.
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}
