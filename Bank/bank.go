package main

import "fmt"

func main() {
	var accountBalance float64 = 1_000

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		var message string = fmt.Sprintf("Your balance is %.2f", accountBalance)
		fmt.Println(message)
	} else if choice == 2 {
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		if depositAmount >= 0 {
			accountBalance += depositAmount
		} else {
			fmt.Println("Invalid amount")
		}
	} else if choice == 3 {
		fmt.Print("Withdraw amount: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		if withdrawAmount < 0 {
			fmt.Println("Invalid amount")
			return
		}
		if withdrawAmount <= accountBalance {
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated: ", accountBalance)
		} else {
			fmt.Println("Insufficient funds")
		}
	} else if choice == 4 {
		fmt.Println("Goodbye!")
	} else {
		fmt.Println("Invalid choice")
	}

}
