package main

import "fmt"

var (
	author  = "@nuatu"
	Version = "0.0.1"
)

const (
	CCVisa       = "Visa"
	CCMasterCard = "MasterCard"
	CCAMex       = "American Express"
)

func main() {
	var greeting string = "Hello World!"
	var a, b, c float64 = 1.1, 2, 3
	fmt.Println(greeting, a, b, c)

	// var name = "Nuatu Tseggai"
	// var d, e, f = 1, 2.0, "Three"
	// fmt.Println(name, d, e, f)

	// course := "Essential Go"
	// x, y, z := 1, 2, 3
	// fmt.Println(course, x, y, z)
}
