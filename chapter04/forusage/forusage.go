package main

import "fmt"

func main() {
	useRangeForLoop()
	fmt.Println("\n***")
	contrastFullFor()
}

func useRangeForLoop() {
	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		if i == 0 {
			continue
		}
		if i == len(evenVals)-2 {
			break
		}
		fmt.Println(i, v)
	}
}

func contrastFullFor() {
	evenVals := []int{2, 4, 6, 8, 10}
	for i := 1; i < len(evenVals)-1; i++ {
		fmt.Println(i, evenVals[i])
	}
}
