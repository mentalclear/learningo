package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	var fred person
	fmt.Println("Fred has all values defaulted:", fred)
	bob := person{} // different way to do the same thing
	fmt.Println("Bob has all values defaulted too:", bob)

	// Simplier way
	julia := person{
		"Julia",
		40,
		"cat",
	}

	// With keys provided
	beth := person{
		age:  30,
		name: "Beth",
	}

	fmt.Println(julia)
	fmt.Println("beth's name is", beth.name)

	bob.name = "Bob"
	fmt.Println("bob's name is", bob.name)

	fmt.Println("\n*****")
	anonStructs()
}

// Anonymous structs
func anonStructs() {

	// person.name = "bob"
	// person.age = 50
	// person.pet = "dog"

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println("pet:", pet)

}

/*
	There are two common situations where anonymous structs are handy.
	The first is when you translate external data into a struct or a struct
	into external data (like JSON or protocol buffers).
	This is called unmarshaling and marshaling data.

	Writing tests is another place where anonymous structs pop up.
*/
