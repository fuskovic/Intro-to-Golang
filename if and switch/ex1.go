package main

import (
	"fmt"
)

//******************************Logical operators and short-circuiting

func main() {

	statePopulations := map[string]int{
		//dont forget your comma after each map index
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	if pop, ok := statePopulations["Florida"]; ok {
		/*
	first if line is called initializer
	sets pop == value paired with the key "Florida" and ok to its error value
	if Florida is in the map, print its value pair
	*/
		fmt.Println(pop)
		//pop variable is only defined the scope of the if statement.
		//Cant be used outside of if statement block
	} else {
		fmt.Println(ok)
		//would output false if a key was used that did not exist or was misspelled
	}

	number := 50
	//guess := 70
	//guess := 30
	guess := -1
	//operators
	if guess < number {
		fmt.Println("Too low")
		//greater than
	}
	if guess > number {
		fmt.Println("Too high")
		//less than
	}
	if guess == number {
		fmt.Println("You got it!")
		//equal to
	}
	fmt.Println(number <= guess, number >= guess, number != guess)
	/*
		<=   greater than or equal to
		>=	 less than or equal to
		!=   not equal to
	*/

	// && is "and" logical operator and || is "or"

	if guess < 1 || returnTrue() || guess > 100 {
		//only one argument has to be true

		/*
			Even though second argument always returns true it will
			still print below Println because once one argument is true
			it drops testing the others in an or statement. 
			
			This is known as short circuiting
		*/

		fmt.Println("Guess must be between 1-100")
	}
	if guess >= 1 && guess <= 100 {
		/*
		both arguments have to be equal to true
		if one of these returns false it short circuits
		*/

		if guess < number {
			fmt.Println("too low")
		}
		if guess > number {
			fmt.Println("too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
	}
}

func returnTrue() bool {
	fmt.Println("Returning true")
	return true
}
