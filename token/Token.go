package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GetSessionToken(signingKey string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(signingKey))

	if err != nil {
		return ""
	}

	return tokenString
}

func Validate(tokenString string, signingKey string) bool {
	token, err := jwt.Parse(tokenString, jwt.Keyfunc(func(t *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	}))
	// Error examples in: https://godoc.org/github.com/dgrijalva/jwt-go#Parse
	return err == nil && token != nil && token.Valid
}
