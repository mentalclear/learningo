package main

import "fmt"

// Go does allow you to perform a type conversion from one struct type
// to another if the fields of both structs have the same names, order, and types.

type firstPerson struct {
	name string
	age  int
}

type secondPerson struct {
	name string
	age  int
}

/*
	We can use a type conversion to convert an instance of firstPerson
	to secondPerson, but we can't use == to compare an instance of
	firstPerson and an instance of secondPerson, because they are different types
*/

// We can't convert an instance of firstPerson to thirdPerson,
// because the fields are in a different order:
type thirdPerson struct {
	age  int
	name string
}

// We can't convert an instance of firstPerson to fourthPerson because
// the field names don't match:

type fourthPerson struct {
	firstName string
	age       int
}

// Finally, we can't convert an instance of firstPerson to fifthPerson
// because there's an additional field:

type fifthPerson struct {
	name          string
	age           int
	favoriteColor string
}

func main() {
	man1 := firstPerson{"James", 30}
	man2 := secondPerson{"Josh", 40}
	woman1 := thirdPerson{25, "Julia"}
	woman2 := fourthPerson{"Anna", 35}
	child := fifthPerson{"Bob", 5, "blue"}

	fmt.Println(man1, man2, woman1, woman2, child)
	fmt.Println("\n***")
	anonStructTwist()
}

func anonStructTwist() {
	f := firstPerson{"Bob", 50}

	// var g struct {
	// 	name string
	// 	age  int
	// }

	// // Compiles
	// g = f

	// Same thing but with no warning
	var g struct {
		name string
		age  int
	} = f
	fmt.Println(f == g)
}
