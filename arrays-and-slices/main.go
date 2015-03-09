package main

import "fmt"

func main() {
	var nums [5]int
	fmt.Println("empty:", nums)

	nums[4] = 100
	fmt.Println("set:", nums)
	fmt.Println("get:", nums[4])
}
