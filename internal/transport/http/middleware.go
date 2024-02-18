package http

import (
	"net/http"
	"time"

	"github.com/anil1226/go-employee/config"
	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte(config.GetEnvKey("JWTSECRET"))

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {

	// Create JWT token
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	// Send token in response
	return tokenString, nil
}

func verifyJWT(endpointHandler func(w http.ResponseWriter, req *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return sampleSecretKey, nil
		})
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		endpointHandler(w, r)
	}
}
