package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//************* Closing response resources when making an API call

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	/*
	   You may not want to defer closing here unless you need all those resources to stay open during
	   the call.

	   This means the resource( page, text file, etc...) will close
	   only after the function is done.

	   May not be suitable for large Requests where each
	   resource needs to be closed as soon as you're done using it.

	   Don't wan't to keep a bunch of resources open at one time and waste memory.

	   Try grouping open and close functions together.

	*/
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	// here we added a close to the API response resources so they don't stay open longer than necessary
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
