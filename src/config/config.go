package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// string connection with dabase and port that api is running
// SecretKey is the key used to auto self sign token

var (
	ConnectionDataBase = ""
	Port               = 0
	SecretKey          []byte
)

// loading going begin the variables of enviroment
func LoadingEnviroment() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		//stop execution
		log.Fatal(erro)

	}

	// the information already staying inside in package "os"

	// set value in PORT

	Porta, erro := strconv.Atoi(os.Getenv("API_PORT"))

	if erro != nil {
		Port = 9000
	} else {
		Port = Porta
	}

	// string to connection with database

	ConnectionDataBase = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
