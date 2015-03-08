package main

import "fmt"

func main() {

	// Exercise:
	// Print out the numbers 0-29
	// with 0-9 on line one,
	// 10-19 on line two,
	// and 20-29 on line three.

	for i := 0; i <= 29; i++ {
		if (i+1)%10 == 0 {
			fmt.Println(i)
		} else {
			fmt.Print(i, " ")
		}
	}

	// if true == true {
	// 	fmt.Println("true is true")
	// } else {
	// 	fmt.Println("true is false")
	// }

	// for i := 0; i < 20; i++ {
	// 	fmt.Println("Go!")
	// }

	// for {
	// 	fmt.Println("This will go forever")
	// 	break
	// }

	// j := true
	// for j {
	// 	fmt.Println("This should happen once")
	// 	j = false
	// }
}
