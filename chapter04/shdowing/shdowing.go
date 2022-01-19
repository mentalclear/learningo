package main

import "fmt"

func main() {
	singleShadowing()
	multiShadowing()

	x := 10
	fmt.Println(x)
	// fmt := "oops" // shadows fmt import here
	// fmt.Println(fmt) // errors due to shadowing,
	// local var knows nothing about Println

	fmt.Println("\n************")
	shadowingTrue()
}

func singleShadowing() {

	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5 // shadowing x here
		fmt.Println(x)
	}
	fmt.Println(x)
}

func multiShadowing() {
	x := 10
	if x > 5 {
		x, y := 5, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)
}

func shadowingTrue() {
	fmt.Println(true)
	true := 10
	fmt.Println(true)
}
