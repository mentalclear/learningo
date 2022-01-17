package main

import "fmt"

func main() {
	indexExpression()
	fmt.Println("\n*************************")
	unicodeChars()
	fmt.Println("\n*************************")
	runeConvToString()
	fmt.Println("\n*************************")
	buggyConversion()
	fmt.Println("\n*************************")
	convertToBytes()

}

func indexExpression() {
	var s string = "Hello there"
	var b byte = s[6]
	fmt.Println(b)

	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s2, s3, s4)
}

func unicodeChars() {
	var s string = "Hello 🌞"
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s2, s3, s4)
	fmt.Println(len(s))
}

func runeConvToString() {
	var a rune = 'x'
	var s string = string(a)
	var b byte = 'y'
	var s2 string = string(b)
	var c rune = 'й'
	var s3 string = string(c)

	fmt.Println(s, s2, s3)
}

func buggyConversion() {
	// var x int = 65
	// var y string = string(x)
	// fmt.Println(y)
	fmt.Println("This is buggy!")
	/*
		The result will be A for y. Since Go 1.15 any conversion to string from any
		interger type other than rune or byte will be blocked.
	*/
}

func convertToBytes() {
	var s string = "Hello 🌞"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}
