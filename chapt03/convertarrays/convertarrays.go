package main

import "fmt"

func main() {
	arrayToSlice()
	fmt.Println("\n*************************")
	makingCopy()
}

func arrayToSlice() {
	x := [4]int{5, 6, 7, 8}
	y := x[:2]
	z := x[2:]
	x[0] = 10
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

// To create an independent slice - make a copy
func makingCopy() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num)
	y[0] = 10
	fmt.Println(x, y)
	x[0] = 100
	fmt.Println(x, y)

	// Copy from the middle:
	// x := []int{1, 2, 3, 4}
	// y := make([]int, 2)
	// copy(y, x[2:])

	// Copy overalpping sections
	// x := []int{1, 2, 3, 4}
	// num = copy(x[:3], x[1:])
	// fmt.Println(x, num)

	// Copy form a slice of an array
	// x := []int{1, 2, 3, 4}
	// d := [4]int{5, 6, 7, 8}
	// y := make([]int, 2)
	// copy(y, d[:])
	// fmt.Println(y)
	// copy(d[:], x)
	// fmt.Println(d)

}
