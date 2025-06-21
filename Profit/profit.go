package main

import (
	"errors"
	"fmt"
	"os"
)

const profitFile string = "profits.txt"

func main() {
	var revenue, expenses, taxRate float64
	var err error

	revenue, err = getUserInput("Revenue: ")
	if err != nil {
		panic(err)
	}

	expenses, err = getUserInput("Expenses: ")
	if err != nil {
		panic(err)
	}

	taxRate, err = getUserInput("Tax Rate: ")
	if err != nil {
		panic(err)
	}

	// Earnings Before Tax
	ebt, profit, ratio := calculateProfits(revenue, expenses, taxRate)
	profitData := fmt.Sprintf("%.2f\n%.2f\n%.2f\n", ebt, profit, ratio)
	writeData := fmt.Sprintf(profitData, ebt, profit, ratio)
	os.WriteFile(profitFile, []byte(writeData), 0644)

	fmt.Printf("Earnings before tax: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Earnings to profit ratio: %.2f\n", ratio)
}

func getUserInput(text string) (float64, error) {
	fmt.Print(text)
	var input float64
	fmt.Scan(&input)
	if input < 0 {
		return 0, errors.New("input cannot be negative")
	}
	if input == 0 {
		return 0, errors.New("input cannot be 0")
	}

	return input, nil
}

func calculateProfits(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
