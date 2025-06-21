package main

import (
	"fmt"

	"bank/fileops"
)

const accountBalanceFile string = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
	}

	fmt.Println("Welcome to Go Bank!")
Loop:
	for {
		presentOptions()

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var message string = fmt.Sprintf("Your balance is %.2f", accountBalance)
			fmt.Println(message)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount < 0 {
				fmt.Println("Invalid amount")
				continue
			}
			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
			fmt.Println("Balance updated: ", accountBalance)
		case 3:
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
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			break Loop
		default:
			fmt.Println("Invalid choice")
		}
		fmt.Println("")
	}
}
