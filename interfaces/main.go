package main

import "fmt"

type Cat struct {
	name string
}

type Dog struct {
	name string
}

func (c Cat) Name() string {
	return c.name
}

func (d Dog) Name() string {
	return d.name
}

func (c Cat) Pet() {
	fmt.Println(" 'cool cat corner store' ")
}

func (d Dog) Pet() {
	fmt.Println(" 'bigg dogg entertainment' ")
}

type Animal interface {
	Pet()
	Name() string
}

func Compliment(a Animal) {
	fmt.Println("Great Job:", a.Name())
	a.Pet()
}

func Feed(a Animal) {
	fmt.Println("Eat it all:", a.Name)
}

func main() {
	c := Cat{"ally cat"}
	Compliment(c)

	d := Dog{"snoop dogg"}
	Compliment(d)
}
