package http

import (
	"time"

	"github.com/anil1226/go-employee/config"
	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte(config.GetEnvKey("JWTSECRET"))

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "username"
	return "", nil
}
