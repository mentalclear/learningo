package main

import "fmt"

func main() {
	createArray()
	fmt.Println("------------------")
	createSparseArray()
	fmt.Println("------------------")
	compareArrays()
	fmt.Println("------------------")
	creatingSlices()
	appendingToSlice()
	fmt.Println("\n*****************")
	understandCapacity()
	fmt.Println("\n*****************")
	createSlicesWithMake()
}

func createArray() {
	// var x [3]int   declare array all values are zero
	var x = [3]int{10, 20, 30}

	for i := range x {
		fmt.Println(x[i])
	}
}

func createSparseArray() {
	var x = [12]int{1, 5: 4, 6, 10: 100, 15}

	for i := range x {
		fmt.Println(x[i])
	}

	// Array length here:
	fmt.Printf("Array's length: %d\n", len(x))
}

func compareArrays() {
	var x = [...]int{10, 20, 30}
	var y = [3]int{10, 20, 30}
	fmt.Println(x == y)
}

func creatingSlices() {
	var x = []int{10, 20, 30} // This is slice [...] makes array
	var y = []int{1, 5: 4, 6, 10: 100, 15}
	var z [][]int // Multidem slice

	fmt.Println(x, y, z)

	// Comparing slice with nil
	fmt.Println(x == nil)
}

func appendingToSlice() {
	var x []int
	x = append(x, 10)
	fmt.Println("Appended:", x)

	var y = []int{1, 2, 3}

	y = append(y, 4)
	fmt.Println("Appended one:", y)

	var z = []int{1, 2, 3}
	z = append(z, 4, 5, 6)
	fmt.Println("Appended multiple:", z)

	// Apending with variadic operator
	y1 := []int{20, 30, 40}
	x = append(x, y1...)
	fmt.Println("Appended with variadic:", x)
}

func understandCapacity() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
}

func createSlicesWithMake() {
	x := make([]int, 5)
	fmt.Println(x)
	x = append(x, 10)
	fmt.Println(x)

	y := make([]int, 5, 10) // 5 is length, 10 is capacity
	fmt.Printf("The slice: %d len: %d cap: %d\n", y, len(y), cap(y))

	z := make([]int, 0, 10)
	z = append(z, 5, 6, 7, 8, 9)
	fmt.Printf("The slice: %d len: %d cap: %d\n", z, len(z), cap(z))
}
