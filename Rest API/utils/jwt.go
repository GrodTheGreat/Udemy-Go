package utils

import (
	"errors"
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

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return errors.New("invalid token")
	}

	//claims, ok := parsedToken.Claims.(jwt.MapClaims)
	//if !ok {
	//	return errors.New("invalid token claims")
	//}

	//email := claims["email"].(string)
	//userId := claims["userId"].(int64)

	return nil
}
