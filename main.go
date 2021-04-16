package main

import (
	"api-nos-golang/src/config"
	"api-nos-golang/src/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadingEnviroment()

	// fmt.Println(config.ConnectionDataBase)

	r := routes.CreateRoutes()

	fmt.Printf("Running api in port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
