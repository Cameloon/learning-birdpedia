//Thiss is the name of our package
//everything with this apckage name can see everything
// else inside the same package, regardless of the file ethey are in 
package main
//these are the libraries we are going to use
// both "fmt" and "net" are part of the GO standard library

import (
	//"fmt" has method for formatted I/O operations (like printing to the console)
	"fmt"
	//the "net/http" library has methods to implement HTTP clients and servers "net/http"
	"net/http"
	"github.com/gorilla/mux"
)


//the new router function creates the router and returns it to us, we can now use this function
//to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	//declare a new router
	r := mux.NewRouter()
	//this is where the  router is useful, it allows us to declare methods that this poath will be valid for
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}

func main() {
	//the router is now formed by calling the `newRouter` constructor function
	//that we defined above. The rest of the code stays the same
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//for this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World!")
}