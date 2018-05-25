package main

//*********************************Constant fundamentals
import (
	"fmt"
)

const (
	e int16 = 27
	g       = iota
	/*
		iota is cool as it recursively adds the type to each var in the const block and increments their values by 1
		see lines 56-58...... NOTE: DEFAULT IOTA VALUE IS 1
	*/
	h
	i
)
const (
	a2 = iota
)

func main() {
	//below is example of a typed constant
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)
	//constants can not be made up of values that are determined at run time

	//const anotherConst float64 = math.Sin(1.57)
	//fmt.Printf("%v, %T\n", anotherConst, anotherConst)
	//error thrown in debugger saying its not constant

	//constants can be made up of all primitive types
	const a int = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	//constants can have their values and types changed..... weird.....thought they were constant???????????

	const e = 14
	fmt.Printf("%v, %T\n", e, e)
	var f int16 = 27
	fmt.Printf("%v, %T\n", e+f, e+f)
	/*
		compiler gives benefit of the doubt to e and implicitly types
		it as int16 to execute the arithmetic
	*/

	fmt.Printf("%v\n", g)
	fmt.Printf("%v\n", h)
	fmt.Printf("%v\n", i)
	/*
	   notice that h and i do not have their values set yet still output 2 and 3
	   iota h and i do not have their value or types set but because they reside
	   in the same block under g = iota, the compiler implicitly types them and increments them accordingly
	*/
	fmt.Printf("%v\n", a2)
	//iota instance resets to 0 in output because it resides in a seperate constant block
}

//NOTE: inner block constants take priority over package level ones
