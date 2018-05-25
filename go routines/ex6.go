package main

import (
	"fmt"
	"runtime"
)

//****************************************** Manipulating runtime threads

func main() {
	runtime.GOMAXPROCS(1)
	//returns number of threads to use
	fmt.Printf("threads: %v\n", runtime.GOMAXPROCS(-1))
	//-1 param doesn't change anything just means give us what we set
	/*
		Go gives you the number of os threads equal to number of cores on the machine
	*/
}

/*
******************Best practices for threading

	1 os thread/core is minimum

	Application get faster by increasing the gomaxprocs

	This can create additional overhead if you create too many threads

	Good to run your app with higher gomax procs than 1 just to expose
	any race conditions.

	Also a good idea to use a performance test suite to test the performance
	of the software using different gomaxprocs settings

	Dont create goroutines in libraries- let users control concurrency

	When creating goroutines, understand how it will end to avoid
	subtle memory leaks

	Check for race conditions at compile time(before running the app) if
	you are running any concurrency at all

		-How? use go run -race {target file}
		will print WARNING: DATA RACE
		will also print line which data race occurred and for which var


*/
