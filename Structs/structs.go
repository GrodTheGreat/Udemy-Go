package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	user, err := user.NewUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		return
	}

	user.OutputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)

	return value
}
