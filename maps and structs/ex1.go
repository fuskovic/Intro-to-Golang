package main

import (
	"fmt"
)

//********************************************** Making maps and returning values by key

func main() {
	/*
		maps match key value pairs
		type have to be consistent
		almost any types can be used
	*/

	//making maps
	statePopulations := make(map[string]int)
	//setting map key/value pairs
	statePopulations = map[string]int{
		//dont forget your comma after each map index
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)
	/*
	slices can not be used as key types, only arrays
	m := map[[]int]string{}
	fmt.Println(m)
	output is "invalid key type []int"
	*/
	m := map[[3]int]string{}
	fmt.Println(m)
	//output is map[] because no key value pairs

	//pullout map values by key
	fmt.Println(statePopulations["Ohio"])
	//will output 11614373
}