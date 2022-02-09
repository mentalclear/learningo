package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(calcRemainderAndMod(4, 2))
	fmt.Println(calcRemainderAndMod(4, 0))
}

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		// Setting other return values to 0
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}
