package main

import "fmt"

func main() {
	doWhileAnalog()
	confusingFizzBuzz()
	cleanFizzBuzz()
}

func doWhileAnalog() {
	i := 100
	condition := true
	for {
		// Things to do in the loop
		i--
		if i == 0 {
			condition = false
		}

		// Breaking out of the loop
		if !condition {
			fmt.Println("Breaking out!...")
			break
		}
	}
}

func confusingFizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func cleanFizzBuzz() {
	fmt.Println("\n**************")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}
