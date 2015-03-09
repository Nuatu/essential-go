package main

import "fmt"

func main() {
	var nums [5]int
	fmt.Println("new empty array:", nums)

	nums[4] = 100
	fmt.Println("set:", nums)
	fmt.Println("get:", nums[4])

	ints := []int{1, 2, 3, 4, 5}
	fmt.Println("new slice:", ints)

	ints = append(ints, 6)
	fmt.Println("appended slice:", ints)

	fmt.Println("0-2:", ints[:2])
	fmt.Println("2-4:", ints[2:4])
	fmt.Println("4-6:", ints[4:])

	for key, val := range ints {
		fmt.Println(key, val)
	}
}
