package main

import "fmt"

func main() {

	var i = 10
	fmt.Println(i)

	// Define an array
	var x [3]int
	fmt.Println(x)

	// Define an array with initial values.
	var x2 = [3]int{10, 20, 30}
	fmt.Println(x2)

	var x3 [3]int = [3]int{10, 20, 30}
	fmt.Println(x3)

	// Sparse Array => An array where most elements are set to zero value and
	// specify only the indices with values in the array literal.
	var x4 [15]int = [15]int{1, 5: 4, 6: 200, 10: 100, 15}
	fmt.Println(x4)

	var x5 = [...]int{10, 20, 30}
	fmt.Println(x5)

	fmt.Println(x3 == x5) // Prints true.

	// Multi-dimensional arrays
	var x6 [2][3]int
	fmt.Println(x6)

	x5[2] = 50
	fmt.Println(x5)

	fmt.Println("Length of array x5 is", len(x5))
}
