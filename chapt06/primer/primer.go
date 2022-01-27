package main

import "fmt"

func main() {
	// x := "Hello"
	// pointerToX := &x

	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)  // prints a memory address
	fmt.Println(*pointerToX) // prints the value it points to

	z := 5 + *pointerToX
	fmt.Println(z)

	// Before dereferencing, pointer shouldn't point to nil
	var x1 *int
	fmt.Println(x1 == nil)
	// fmt.Println(*x1)
	/* Cusing
		panic: runtime error: invalid memory address or nil pointer dereference
	[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x47e3b1]
	*/

	// Pointer type
	x2 := 10
	var pointerToX2 *int
	pointerToX2 = &x2

	fmt.Println(*pointerToX2)

	var x3 = new(int)
	fmt.Println(x3 == nil) // false
	fmt.Println(*x3)       // 0

	/*
		You can't use an & before a primitive literal (numbers, booleans, and strings)
		or a constant because they don't have memory addresses; they exist only at compile
		time. When you need a pointer to a primitive type, declare a variable and point
		to it:
	*/
	// x4 := &Foo{}
	// var y string
	// z := &y

	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	p := person{
		FirstName: "Pat",
		//MiddleName: "Perry", // This line won't compile
		MiddleName: stringp("Perry"), // This works now
		LastName:   "Peterson",
	}
	fmt.Println(p)
}

// Helper function to resolve issue with pointer in the struct
func stringp(s string) *string {
	return &s
}
