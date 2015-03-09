package main

import "fmt"

func main() {
	num := 5

	double(num)
	fmt.Println(num)

	triple(&num)
	fmt.Println(num)
}

func double(n int) {
	n = n * 2
}

func triple(n *int) {
	*n = *n * 3
}
