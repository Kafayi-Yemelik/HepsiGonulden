package authentication

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func JwtGenerator(id string, firstname string, lastname string) (string, error) {

	claims := jwt.MapClaims{
		"Id":        id,
		"FirstName": firstname,
		"LastName":  lastname,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))

	return t, err
}
