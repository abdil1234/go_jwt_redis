package helpers

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	jwt.StandardClaims
	Username string
}

type M map[string]interface{}

var APPLICATION_NAME = "My Simple JWT App"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")
