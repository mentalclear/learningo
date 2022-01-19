package main

import "fmt"

func main() {
	forRangeLoop()
	forRangeNoKey()
	onlyKeyNoValues()
}

// You can only use a for-range loop to iterate over the built-in compound
// types and user-defined types that are based on them.

func forRangeLoop() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for index, value := range evenVals { // To make it more visual for me i, v renamed
		fmt.Println(index, value)
	}
}

// i (index) used for indexes in arrays, slices or strings
// k (key) used for maps
// v (value) but sometimes given a name based on the type

// Avoid using the key:
func forRangeNoKey() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals { // Ignoring key(or index) with _
		fmt.Println(v)
	}
}

// The most common reason for iterating over the key is when a map is being used as a set.
// In those situations, the value is unimportant.
func onlyKeyNoValues() {
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}
}
