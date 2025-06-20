package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, expectedReturnRate, years float64

	investmentAmount = outputText("Investment Amount: ")
	expectedReturnRate = outputText("Expected Return Rate: ")
	years = outputText("Years Invested: ")

	futureValue, realFutureValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", realFutureValue)
	// fmt.Println("Future Value: ", futureValue)
	// fmt.Printf("Future Value: %.2f\n", futureValue)
	// fmt.Println("Future Value (adjusted for inflation): ", realFutureValue)
	// fmt.Printf("Future Value (adjusted for inflation): %.2f\n", realFutureValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) float64 {
	fmt.Print(text)
	var value float64
	fmt.Scan(&value)

	return value
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+(expectedReturnRate/100), years)
	rfv := fv / math.Pow(1+(inflationRate/100), years)

	return fv, rfv
}
