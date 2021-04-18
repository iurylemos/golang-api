package main

import (
	"api-nos-golang/src/config"
	"api-nos-golang/src/routes"
	"fmt"
	"log"
	"net/http"
)

// function used to generate key base 64 equal be in .env
// func init() {
// 	key := make([]byte, 64)

// 	// filling this key that have 64 positions 0
// 	if _, erro := rand.Read(key); erro != nil {
// 		log.Fatal(erro)
// 	}

// 	// salving this slice to save in .env

// 	stringBase64 := base64.StdEncoding.EncodeToString(key)

// 	fmt.Println(stringBase64)
// }

func main() {
	config.LoadingEnviroment()

	r := routes.CreateRoutes()

	fmt.Printf("Listening in port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
