package main

import "fmt"

type person struct {
	age  int
	name string
}

func main() {
	p := person{}
	i := 2
	s1 := "Hello"
	modifyFails(i, s1, p)
	fmt.Println(i, s1, p)

	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	modSlice(s)
	fmt.Println(s)
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}
