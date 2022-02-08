package main

import (
	"fmt"
	"io"
)

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)

	// Produces:
	// panic: interface conversion: interface {} is main.MyInt, not string
	// i3 := i.(string)
	// fmt.Println(i3)

	// as well as this one
	// i4 := i.(int)
	// fmt.Println(i4 + 1)

	// Use comma Ok idiom
	// i5, ok := i.(int)
	// if !ok {
	// 	return fmt.Errorf("Unexpected type for %v", i)
	// }
	// fmt.Println(i5 + 1)

	doThings(i)
}

// When an interface could be one of multiple possible types,
// use a type switch instead:
func doThings(i interface{}) {
	switch j := i.(type) {
	case nil:
		// i is nil, type of j is interface{}
	case int:
		// j is of type int
	case MyInt:
		// j is of type MyInt
	case io.Reader:
		// j is of type io.Reader
	case string:
		// j is a string
	case bool, rune:
		// i is either a bool or rune, so j is of type interface{}
	default:
		// no idea what i is, so j is of type interface{}
		fmt.Println(j)
	}

}
