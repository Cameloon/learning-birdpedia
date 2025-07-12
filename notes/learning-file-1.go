//Thiss is the name of our package
//everything with this apckage name can see everything
// else inside the same package, regardless of the file ethey are in 
// package main -> have to comment this out otherwise the compiler throws an error because of dupliacte package main (in main.go) 
package notes

//these are the libraries we are going to use
// both "fmt" and "net" are part of the GO standard library

import (
	//"fmt" has method for formatted I/O operations (like printing to the console)
	"fmt"
	//the "net/http" library has methods to implement HTTP clients and servers "net/http"
	"net/http"
)

func main() {
	//the "HandleFunc" method accepts a path and a function as arguments
	//(yes, we can pass functions as arguments, and even treat them like variables in Go)
	//However, the handler function has to have the appropriate signature (as described by the "handler" func below)
	http.HandleFunc("/", handler)

	//After defining our server, we finally "listen and serve" on port 8080
	//the second argument is the handler, which we will come to later on, but for now it is left as nil,
	//and the handler defined above (in "HandleFunc") is used
	http.ListenAndServe(":8080", nil)
}

//"handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request
//tye as the arguments
func handler(w http.ResponseWriter, r *http.Request) {
	//for this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World!")
}