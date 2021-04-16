package main

import (
	"api-nos-golang/src/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := routes.CreateRoutes()

	fmt.Printf("Running api in port %s", "5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
