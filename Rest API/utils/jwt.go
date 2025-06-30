package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "supersecretkey"

func GenerateToken(id int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": id,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
