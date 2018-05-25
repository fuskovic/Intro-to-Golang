package main

import (
	"fmt"
)

//****************** Switch statements- tag and tagless syntax and fallthrough

func main() {

	//switch statements take a value and find a case it falls under

	/*
		Tag syntax
		You can also set a tag/variable here, and then specify exactly what you're testing.
		
		Second parameter is what actually gets tested.
	*/
	switch i := 2 + 3; i {
	
	case 1, 5, 10:
		//notice a case uses colon to start and nothing to close as next case is treated as the delimiter
		fmt.Println("one")
		//cases must be unique- for example. Each of these cases could not contain 5
	case 2, 4, 6:
		fmt.Println("two")

	default:
		fmt.Println("not one or two")
	}

	/*
		Tagless syntax
	*/
	j := 10

	switch {

	case j <= 10:
		fmt.Println("less than or equal to ten")
		/*
			Notice how these two cases over lap as 10 is less than 10 and 20.
			This is done to illustrate tagless syntax switch statement short circuiting
			however, you can continue testing the other cases using keyword "fallthrough"
		*/
		fallthrough
		//Output executes for both cases now
	case j <= 20:
		fmt.Println("less then or equal to twenty")

	default:
		fmt.Println("Greater than twenty")
	}
	}