package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Checking if the user that be making the request be authentication
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("called function to valid token and then execute function that come in parameters")
		//function "next" go execute function that be coming of parameters of function
		//this case is the function "route.Function"
		next(w, r)
	}
}

// Logger write logs in terminal and use "next" to execute function that come in parameters
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Logger...")
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}
