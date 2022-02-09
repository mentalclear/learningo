package main

import "fmt"

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func GenerateError(flag bool) error {
	// var genErr StatusErr - creates unexpected output
	// Fix for this
	var genErr error

	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}

	// return genErr
	// return nil // but in this case genErr is unused
	return genErr
}

func main() {
	err := GenerateError(true)
	fmt.Println(err != nil)
	err = GenerateError(false)
	fmt.Println(err != nil)
}
