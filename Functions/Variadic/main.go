package main

import "fmt"

func main() {
	sum := sumup(1, 10, 15, 40, -15)
	fmt.Println(sum)
}

func sumup(startingVal int, numbers ...int) int {
	sum := startingVal

	for _, val := range numbers {
		sum += val
	}

	return sum
}
