package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sampleSwitch()
	breakSwitchInForBad()
	fmt.Println("\n**********")
	breakSwitchInForGood()
	fmt.Println("\n**********")
	blankSwitches()
	fmt.Println("\n**********")
	rewriteIfElse()

}

func sampleSwitch() {
	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		// Empty cases are just ignored!
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}
}

func breakSwitchInForBad() {
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break
		default:
			fmt.Println(i, "is boring")
		}
	}
}

func breakSwitchInForGood() {
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}
}

// Any boolean comaprison can be used in case here.
func blankSwitches() {
	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}
}

func rewriteIfElse() {
	switch n := rand.Intn(10); {
	case n == 0:
		fmt.Println("That's too low")
	case n > 5:
		fmt.Println("That's too big:", n)
	default:
		fmt.Println("That's a good number:", n)
	}
}
