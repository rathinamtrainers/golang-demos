package main

import "fmt"

func main() {
	var x1 = []int{10, 20, 30}
	fmt.Println(x1)

	var x2 = []int{1, 5: 4, 10, 11: 100, 15}
	fmt.Println(x2)

	// Multi Dimensional Slices
	var xm [][]int
	fmt.Println(xm)

	x1[0] = 11
	fmt.Println(x1)

	var xa [5]int
	fmt.Println(xa)

	var xs []int
	fmt.Println(xs) // nil

	fmt.Println(xs == nil)
	fmt.Println("Size of slice is ", len(xs))

	xs = append(xs, 10)
	fmt.Println(xs)

	var xs2 = []int{1, 2, 3}
	xs2 = append(xs2, 99)
	fmt.Println(xs2)
	xs2 = append(xs2, 50, 63, 72, 86, 94)
	fmt.Println(xs2)
	fmt.Println("Length of slice xs2 is ", len(xs2))

	fmt.Println("*** Capacity demonstration against len function ***")
	var xs3 []int
	fmt.Println(xs3, len(xs3), cap(xs3))
	xs3 = append(xs3, 10)
	fmt.Println(xs3, len(xs3), cap(xs3))
	xs3 = append(xs3, 20)
	fmt.Println(xs3, len(xs3), cap(xs3))
	xs3 = append(xs3, 30)
	fmt.Println(xs3, len(xs3), cap(xs3))
	xs3 = append(xs3, 40)
	fmt.Println(xs3, len(xs3), cap(xs3))
	xs3 = append(xs3, 50)
	fmt.Println(xs3, len(xs3), cap(xs3))

	// Next session: make(), Map, String, Structures
}
