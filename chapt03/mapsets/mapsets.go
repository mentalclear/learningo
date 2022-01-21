package main

import "fmt"

func main() {
	setSimulation()
	fmt.Println("\n***")
	setSimulWithStruct()
}

// using map to simulate some of set features
func setSimulation() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	// len returns 8 due to no duplicates allowed in the map
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	intSet[100] = true
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}

func setSimulWithStruct() {
	intSet := map[int]struct{}{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = struct{}{}
	}
	if _, ok := intSet[5]; ok {
		fmt.Println("5 is in the set")
	}
}
