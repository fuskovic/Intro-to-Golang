package main

import (
	"fmt"
)

//*******************************************Building structs and returning their values

//think of structs as custom types

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	//create instance of Doctor type and set field values
	aDoctor := Doctor{
		//notice comma syntax after each field value
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor)
	// outputs all field values within a pair of {}

	//returning only one value of the struct
	fmt.Println(aDoctor.companions[1])

	/*
	 it should be noted...the above aDoctor instance of struct,
	 lists the field values its about to set like "number:"
	 and "actorName" but doesnt necessarily have to
	*/

	doctorB := Doctor{
		4,
		"generic actor name",
		[]string{
			"companion1",
			"companion2",
			"companion3",
			//include a comma after closing the []string
		},
	}
	/*
		the above is correct golang syntax but is not encouraged
		because it could throw an error or cause confusion
		 if another field is added to the struct.

		Setting the field names when creating structs
		is best practice because you don't even have to do it in order
		and the compiler will accurately understand where each value
		goes.

		When naming structs, if they have a capital letter, they are
		external to the package. If they have a lowercase letter,
		they are internal to the package. You can also do the same
		with the field names to export them outsidethe package.
	*/

	fmt.Println(doctorB)

}
