package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, expectedReturnRate, years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years Invested: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+(expectedReturnRate/100), years)
	realFutureValue := futureValue / math.Pow(1+(inflationRate/100), years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", realFutureValue)
	// fmt.Println("Future Value: ", futureValue)
	// fmt.Printf("Future Value: %.2f\n", futureValue)
	// fmt.Println("Future Value (adjusted for inflation): ", realFutureValue)
	// fmt.Printf("Future Value (adjusted for inflation): %.2f\n", realFutureValue)
	fmt.Print(formattedFV, formattedRFV)
}
