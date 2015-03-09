package main

import "fmt"

func main() {
	num := 5
	x, y := 20, 40

	swap(&x, &y)
	fmt.Println(x, y)

	double(num)
	fmt.Println(num)

	triple(&num)
	fmt.Println(num)
}

func swap(x, y *int) {
	temp_x := *x
	*x = *y
	*y = temp_x
}

func double(n int) {
	n = n * 2
}

func triple(n *int) {
	*n = *n * 3
}
