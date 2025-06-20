package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	// Earnings Before Tax
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate)
	ratio := ebt / profit

	fmt.Println("Earnings before tax: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Earnings to profit ratio: ", ratio)
}
