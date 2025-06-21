package main

import "fmt"

func main() {
	var accountBalance float64 = 1_000

	fmt.Println("Welcome to Go Bank!")
	for {
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
			if depositAmount < 0 {
				fmt.Println("Invalid amount")
				continue
			}
			accountBalance += depositAmount
		} else if choice == 3 {
			fmt.Print("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount < 0 {
				fmt.Println("Invalid amount")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated: ", accountBalance)
		} else if choice == 4 {
			fmt.Println("Goodbye!")
			break
		} else {
			fmt.Println("Invalid choice")
		}
	}
}
