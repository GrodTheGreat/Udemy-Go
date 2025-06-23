package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}

	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)

	trnsFn1 := getTransformer(numbers[0])
	trnsFn2 := getTransformer(numbers[2])

	transformedNumbers := transformNumbers(&numbers, trnsFn1)
	transformedNumbers2 := transformNumbers(&numbers, trnsFn2)

	fmt.Println(transformedNumbers, transformedNumbers2)

	quadrupled := transformNumbers(&numbers, func(num int) int { return num * 4 })
	fmt.Println(quadrupled)

	quintuple := createTransformer(5)
	quintupled := transformNumbers(&numbers, quintuple)
	fmt.Println(quintupled)
}

type transformFn = func(int) int

func transformNumbers(numbers *[]int, transform transformFn) []int {
	tNumbers := []int{}
	for _, value := range *numbers {
		tNumbers = append(tNumbers, transform(value))
	}

	return tNumbers
}

func getTransformer(number int) transformFn {
	if number == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func createTransformer(factor int) transformFn {
	return func(number int) int {
		return number * factor
	}
}
