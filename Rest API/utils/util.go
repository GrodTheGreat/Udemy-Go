package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
)

func CheckClose(c io.Closer, err *error) {
	if closeErr := c.Close(); closeErr != nil {
		if *err != nil {
			fmt.Println(*err)
		} else {
			*err = closeErr
		}
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}
