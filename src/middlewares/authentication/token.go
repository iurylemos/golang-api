package authentication

import (
	"api-nos-golang/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// how it works a json web token
// it simply take a series of information that are pass
// example: id user and permissions that it have
// it take this and transform in a string that called of token
// this token take all informations that this user have and too who user that be me refer
// and can be used as a authentication in our api
// this data stay hidden inside that string

func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID

	//secret to do signature this token and ensure authentication

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)

	token, erro := jwt.Parse(tokenString, returnKeyVerification)
	if erro != nil {
		return erro
	}

	// return all the claims that was settings in function create token
	// token valid verify if this token is valid or not
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("token invalid")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	// Bearer eyashuhkekaaaaaE02y

	//verify if lenght this string have space between two words
	if len(strings.Split(token, " ")) == 2 {
		// if exists two words I do the return the two word
		// that contain a token
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnKeyVerification(token *jwt.Token) (interface{}, error) {
	// verify method signin
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("method signin unexpected %v", token.Header["alg"])
	}

	// if not happened no one error then i was this that i wanted

	return config.SecretKey, nil
}
