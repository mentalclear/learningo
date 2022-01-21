package main

import "fmt"

func main() {
	modTheCopy()
}

// Each time the range loop iterates over the compound type
// it copies the value and modifies that copy not the source.
func modTheCopy() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	new_eventVals := []int{}
	for _, v := range evenVals {
		v *= 2
		new_eventVals = append(new_eventVals, v)
	}
	fmt.Println(evenVals)
	fmt.Println(new_eventVals)
}
