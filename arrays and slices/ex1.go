package main

import (
	"fmt"
)

//*********************************** Arrays in Go

func main() {
	//arrays must have their size specified in order to be considered arrays
	grades := [3]int{97, 85, 93}
	// you can also use [...]as size as if to say "hold whatever i store"
	fmt.Printf("Grades: %v", grades)

	//still an array as it holds 3 indices even though value of those indices is not set
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	//output is empty for now. No values set yet
	students[0] = "Lisa"
	students[2] = "Ahmed"
	students[1] = "Arnold"
	fmt.Printf("%v\n", students[0])
	//outputs the designated index in students
	fmt.Printf("Student: %v has a grade: %v in the class\n", students[0], grades[2])
	//prints first index of []students and third index of []grades as
	fmt.Printf(" there are %v students in the class\n", len(students))
	//pass []students array to len func to output length of array
}
