package http

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/anil1226/go-employee/config"
	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte(config.GetEnvKey("JWTSECRET"))

func GenerateJWT() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Now().Add(10 * time.Minute),
	})

	tokenString, err := token.SignedString(sampleSecretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyJWT(endpointHandler func(w http.ResponseWriter, req *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		tokenParts := strings.Split(authToken, " ")
		var token *jwt.Token
		var err error
		if len(tokenParts) > 0 {
			token, err = jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
				return sampleSecretKey, nil
			})
		} else if len(tokenParts) == 0 {
			token, err = jwt.Parse(tokenParts[0], func(token *jwt.Token) (interface{}, error) {
				return sampleSecretKey, nil
			})
		}

		if token.Valid {
			fmt.Println("You look nice today")
			endpointHandler(w, r)
		} else if ve, ok := err.(*jwt.ValidationError); ok {

			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("That's not even a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				fmt.Println("Timing is everything")
			} else {
				fmt.Println("Couldn't handle this token:", err)
			}
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		} else {
			fmt.Println("Couldn't handle this token:", err)
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
	}
}
