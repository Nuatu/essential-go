package main

import (
	"log"
	"os"

	"github.com/nuatu/essential-go/generator/pages"
)

func main() {

	// We use target [1] to grab the argument
	// since [0] in this slice is the path to the binary

	if len(os.Args) < 2 {
		log.Fatalln("Not enough arguments")
	}

	filename := os.Args[1]
	p, err := pages.NewPage(filename)
	if err != nil {
		log.Fatalln(err)
	}

	err = p.Render("layout.html", os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}
