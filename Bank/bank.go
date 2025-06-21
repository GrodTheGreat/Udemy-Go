package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile string = "balance.txt"
const defaultBalance float64 = 1_000

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return defaultBalance, errors.New("failed to find balance.txt")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return defaultBalance, errors.New("failed to parse stored balance")
	}

	return balance, nil
}

func main() {
	accountBalance, err := getBalanceFromFile()
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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			break Loop
		default:
			fmt.Println("Invalid choice")
		}
		fmt.Println("")
	}
}
