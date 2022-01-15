package main

import "fmt"

func main() {
	slicingSlices()
	fmt.Println("\n*****************")
	overalapSlices()
	fmt.Println("\n*****************")
	slicingAppend()
	fmt.Println("\n*****************")
	confusingSlices()
}

func slicingSlices() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}

func overalapSlices() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	x[1] = 20
	y[0] = 10
	z[1] = 30
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func slicingAppend() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	fmt.Println(cap(x), cap(y))
	y = append(y, 30)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}

func confusingSlices() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]                          // 1, 2
	z := x[2:]                          // 3, 4
	fmt.Println(cap(x), cap(y), cap(z)) // 5, 2, 2
	y = append(y, 30, 40, 50)           // 1, 2, 30, 40, 50
	x = append(x, 60)                   // 1, 2, 3, 4, 60
	z = append(z, 70)                   // 1, 2, 3, 4, 60, 70
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
