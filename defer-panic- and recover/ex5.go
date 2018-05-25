package main

import (
	"net/http"
)

//********************************** Handling your own errors

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	//returns an error value
	if err != nil {
		panic(err.Error())
		/*
			Go gives you a lot of flexibility on how to handle errors and expects you,the dev,
			to understand how to handle it. Also, deferred functions used to close resources will
			still happen even if there are panics due to the order of execution.
		*/
	}
}

/*
	ORDER OF EXECUTION

		the func
		defers
		panics
		return values.
*/
