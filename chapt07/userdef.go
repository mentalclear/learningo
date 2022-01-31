package main

// Declaring a user defined type:
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Concrete types
type Score int
type Converter func(string) Score
type TeamScores map[string]Score

// Abstract type - specifies what a type should do but not how
// Concrete type - specifies what and how.

func main() {

}
