package cmdmanager

import "fmt"

type CMDManager struct{}

func (cm *CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER")
	fmt.Println("Enter 0 to quit the loop")

	var prices []string
	for {
		fmt.Print("Price: ")
		var price string
		fmt.Scan(&price)
		prices = append(prices, price)

		if price == "0" {
			break
		}
	}

	return prices, nil
}

func (cm *CMDManager) WriteResult(data any) error {
	fmt.Println(data)

	return nil
}

func New() *CMDManager {
	return &CMDManager{}
}
