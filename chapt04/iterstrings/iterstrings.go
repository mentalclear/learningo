package main

import "fmt"

func main() {
	overStrings()
}

// behavior here - iterate over runes not bytes.
// for-range loop does that
func overStrings() {
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r)) // index, rune, string
		}
		fmt.Println()
	}
}
