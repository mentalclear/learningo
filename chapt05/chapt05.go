package main

import "fmt"

func main() {
	result := div(5, 0)
	fmt.Println(result)
}

// func div(numerator int, denominator int) int {
func div(numerator, denominator int) int { // same type params
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}
