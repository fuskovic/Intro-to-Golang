package main

import (
	"fmt"
)

//********************************************* Appending and deleting key/valu pairs from Maps

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{

		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	/*
		appending to map is easy since it does not have
		a fixed length or capacity.
		Simpley add the key and set its value with syntax below
	*/
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations["Georgia"])
	/*
		NOTE:eturn orders of maps are not guaranteed
		could be appended anywhere
	*/
	fmt.Println(statePopulations)
	//output will list the above values in different orders on each call

	//deteting key value pairs
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Georgia"])
	//trying to print key value pair that doesn't exist returns 0
}
