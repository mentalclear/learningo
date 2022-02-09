package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	result, err := doubleEven(2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)

	second_result, err := doubleEvenTwo(3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(second_result)

}

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("only even numbers are processed")
	}
	return i * 2, nil
}

func doubleEvenTwo(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
}
