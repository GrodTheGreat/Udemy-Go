package main

import (
	"fmt"

	usr "example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	user, err := usr.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		return
	}

	user.OutputUserData()

	admin, err := usr.NewAdmin("mail@example.com", "password")
	if err != nil {
		return
	}
	admin.OutputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)

	return value
}
