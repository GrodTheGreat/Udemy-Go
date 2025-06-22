package user

import (
	"errors"
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
}

func (u *User) OutputUserData() {
	fmt.Println(u.FirstName, u.LastName, u.Birthdate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}

func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name, and birth date are required")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
	}, nil
}
