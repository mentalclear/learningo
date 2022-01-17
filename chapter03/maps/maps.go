package main

import "fmt"

func main() {

	// nil map example
	//var nilMap map[string]int // String keys and int values here

	/*
		Attempting to read a nil map always returns the zero value for the map's
		value type. However, attempting to write to a nil map variable causes a panic.
	*/

	// Empty map literal example, can be red and written to
	// totalWins := map[string]int{}

	// non-empty map example

	teams := map[string][]string{
		// Originally:
		/*
			"Orcas": []string{"Fred", "Ralph", "Bijou"},
			"Lions": []string{"Sarah", "Peter", "Billie"},
			"Kittens": []string{"Waldo", "Raul", "Ze"},
		*/
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams["Orcas"])

	ages := make(map[int][]string, 10) // 10 is the size of the map
	fmt.Println(ages, ages[0] == nil)
	ages[0] = []string{"Fred", "Ralph", "Bijou"}
	ages[1] = []string{"Sarah", "Peter", "Billie"}
	ages[2] = []string{"Waldo", "Raul", "Ze"}
	fmt.Println(ages[2])
	fmt.Println(len(ages))

	fmt.Println("\n************************")
	readingMap()
	fmt.Println("\n************************")
	commaOkIdiom()
	deletingFromMap()
}

func readingMap() {
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])
}

// Used to learn if a key is in the map
func commaOkIdiom() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)
	if !ok {
		fmt.Println("goodbye is not in the map")
	}
}

// Deleting from a map
func deletingFromMap() {
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	delete(m, "hello")
	v, ok := m["hello"]
	if !ok {
		fmt.Println("\"hello\" is not in the map anymore:", v)
	}
}
