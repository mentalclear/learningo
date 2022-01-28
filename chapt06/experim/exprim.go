package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)

	y := multip(x)
	fmt.Println(y)

}

func multip(item int) int {
	return item * 2
}
