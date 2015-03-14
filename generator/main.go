package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// We use target [1] to grab the argument
	// since [0] in this slice is the path to the binary

	if len(os.Args) < 2 {
		log.Fatalln("Not enough arguments")
	}

	filename := os.Args[1]
	fmt.Println(filename)
}
