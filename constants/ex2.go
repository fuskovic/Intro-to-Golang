package main

//*******************************Using consts to check for equality
import (
	"fmt"
)

const (
	errorSpecialist = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {

	var specialistType int = catSpecialist
	fmt.Printf("%v\n", specialistType == catSpecialist)
	//checks const type int for value equality
	var anotherSpecialistType int
	//int defaults to 0
	fmt.Printf("%v\n", anotherSpecialistType == dogSpecialist)
	//false because anotherSpecialistType = 0 and dogSpecialist = 3
}
