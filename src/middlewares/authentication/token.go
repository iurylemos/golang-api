package authentication

import (
	"api-nos-golang/src/config"
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
