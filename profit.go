package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	revenue = getUserInput("Revenue: ")
	expenses = getUserInput("Expenses: ")
	taxRate = getUserInput("Tax Rate: ")

	// Earnings Before Tax
	ebt, profit, ratio := calculateProfits(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Earnings to profit ratio: %.2f\n", ratio)
}

func getUserInput(text string) float64 {
	fmt.Print(text)
	var input float64
	fmt.Scan(&input)

	return input
}

func calculateProfits(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
