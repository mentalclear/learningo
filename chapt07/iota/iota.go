package main

type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

type BitField int

const (
	Field1 BitField = 1 << iota // assigned 1
	Field2                      // assigned 2
	Field3                      // assigned 4
	Field4                      // assigned 8
)

func main() {

}
