package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[1])
	prices = append(prices, 5.99)
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
// 	productNames[2] = "A carpet"
// 	fmt.Println(productNames)
// 	fmt.Println(prices)
// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(featuredPrices)
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
