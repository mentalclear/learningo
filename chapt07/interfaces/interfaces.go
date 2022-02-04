package main

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	// business logic
	return "Result"
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	// get data from somewhere
	data := "Dah"
	c.L.Process(data)
}

func main() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
