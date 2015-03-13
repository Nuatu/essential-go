package main

import (
	"fmt"
)

func main() {

	age := make(map[string]int)

	age["allie"] = 27
	age["jj"] = 28
	age["liz"] = 29
	fmt.Println(age)

	delete(age, "jj")
	fmt.Println(age)
	fmt.Println("Allie's age:", age["allie"])

	people := map[string]int{
		"jeremy": 24,
		"jordie": 22,
		"josh":   27,
	}
	fmt.Println(people)
	fmt.Println("Jeremy's age:", people["jeremy"])

	k := keys(people)
	fmt.Println(k)
}

func keys(m map[string]int) []string {
	names := []string{}

	for name, _ := range m {
		names = append(names, name)
	}
	return names
}
