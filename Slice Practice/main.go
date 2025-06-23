package main

import "fmt"

func main() {
	// 1
	hobbies := [3]string{"Programming", "Sleeping", "Eating"}
	fmt.Println(hobbies)
	// 2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])
	// 3
	h := hobbies[:2]
	h2 := []string{hobbies[0], hobbies[1]}
	fmt.Println(h2)
	// 4
	h = h[1:3]
	// 5
	goals := []string{}
	goals = append(goals, "Learn Go")
	goals = append(goals, "Get Money")
	// 6
	goals[1] = "In depth!"
	goals = append(goals, "Fast")
	// 7
	type Product struct {
		id    string
		title string
		price float64
	}
	products := []Product{Product{"1", "Product 1", 10.99}, Product{"2", "Product 2", 13.99}}
	products = append(products, Product{"3", "Product 3", 11.99})
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
