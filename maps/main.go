package main

import "fmt"

func main() {

	age := make(map[string]int)

	age["allie"] = 27
	age["jj"] = 28
	age["liz"] = 29
	fmt.Println(age)

	fmt.Println("Allie's age:", age["allie"])

	delete(age, "jj")
	fmt.Println(age)

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(m)

	for pos, num := range m {
		fmt.Printf("%v is associated with %v", pos, num)
	}
}
