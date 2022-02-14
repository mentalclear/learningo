package main

import (
	"fmt"

	print "github.com/mentalclear/mytestmod/formatter"
	"github.com/mentalclear/mytestmod/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println("Output:", output)
}
