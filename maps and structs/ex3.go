package main

import (
	"fmt"
)

//********************************************* Using bools to check if keys are present in a map

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

	//You can add an ok var to check if a key value pair exists in map
	pop, ok := statePopulations["oho"]
	fmt.Println(pop, ok)

	/*
		misspelled ohio on purpose. Output is 0, false
		if you only want to return the bool value you
		can use a write only var becuase we know false
		will be 0 everytime
	*/

	//ok var name is not special just typical convention
	_, ok := statePopulations["oho"]
	fmt.Println(ok)
	//output is false.

	//find amount of map elements
	fmt.Println(len(statePopulations))
	//outputs 7

	//maps pass data by reference just like slices
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)
	//both sp and statePopulations no longer contain ohio

}
