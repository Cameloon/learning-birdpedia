package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
)


//only tests the handler and not the routing to our handler! meaing it ensures that the req is coming in
//will get served correctly provided that it's delivered to the correct handler
func TestHandler(t *testing.T) {
	//here we form a new HTTP request - this is the request that going to be passed to our handler
	//the firs argument is the method, the second argument is the route (which we leave blank for now buit also get back to soon)
	//and the thid is the request body, which we dont have in this case
	req, err := http.NewRequest("GET", "", nil)

	//ion case there is an error in forming the request, we fail and stop the test
	if err != nil {
		t.Fatal(err)
	}

	//we use Go`s httptest library to create an http recorder - this recorder will act as the target of
	//our http request, (think of it as a mini-browser which will accept the result of the http req that we make)
	recorder := httptest.NewRecorder()

	//create an HTTP handler from our handler function, "handler" is the handler
	//function definded in our main.go file that we want to test
	hf := http.HandlerFunc(handler)

	//serve the HTTP request to our recorder, this is the line that actually executes our the handler
	// that we want to test 
	hf.ServeHTTP(recorder, req)

	//check the status code is what we expect
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//check the response body is what we expect
	expected := "Hello World!"
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}



func TestRouter(t *testing.T) {
	//instantiate the router using the constructor function that we defined previously
	r := newRouter()

	//create a new server using the "httptest" libraries "NewServer" method
	//documentation: https://golang.org/pkg/net/http/httptest/#NewServer
	mockServer := httptest.NewServer(r)

	//the mock server (object/variable??) that we created above runs a server and exposes 
	//its location in the URL attribute we make a GET request to the "hello" route we difned in the router
	resp, err := http.Get(mockServer.URL + "/hello")

	//handle unexpected errosL:
	if err != nil {
		t.Fatal(err)
	}

	//we want our status toi be 200(OK)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	//in the next few lines, the response body is read, and converted to a string
	defer resp.Body.Close()
	//read the bopdy into a bunch of bytes (b)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	//convert the bytes to string
	respString := string(b)
	expected := "Hello World!"

	//we want our response to mathc the on defined in our handler
	//if it does hapopoen to be "Hello World!", the it confirms, that the route is correct
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestRouterForNonExistentRoute(t *testing.T)   {

	r := newRouter()
	mockServer := httptest.NewServer()
	//Most of rthe code is similar, the only differences is that now we make a request to a route we know
	//we dindt define like the "POST /hello" route
	resp, err := http.Post(mockServer.URL + "/hello", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	//we want our status to be 405 (method not allowed)
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405", got %d, resp.StatusCode)
	}


	//the codef to test the body is also mostly the same,  except this time we expect an empty body
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respStrings)
	}

}