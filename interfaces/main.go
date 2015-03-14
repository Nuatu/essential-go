package main

import "fmt"

type Cat struct {
	name string
}

type Dog struct {
	name string
}

type Fish struct {
	name string
}

func (c Cat) Name() string {
	return c.name
}

func (d Dog) Name() string {
	return d.name
}

func (f Fish) Name() string {
	return f.name
}

func (c Cat) Pet() {
	fmt.Println(" 'cool cat corner store' ")
}

func (d Dog) Pet() {
	fmt.Println(" 'bigg dogg entertainment' ")
}

func (f Fish) Pet() {
	fmt.Println(" 'you trippin, ain't nobody got time for that' ")
}

func (c Cat) Eat() {
	fmt.Println("   'love me some catnip'   ")
}

func (c Dog) Eat() {
	fmt.Println("   'throw a dogg a bone'   ")
}

func (c Fish) Eat() {
	fmt.Println("   'who got that seaweed yo'   ")
}

type Animal interface {
	Name() string
	Pet()
	Eat()
}

func Compliment(a Animal) {
	fmt.Println("Great Job:", a.Name())
	a.Pet()
}

func Feed(a Animal) {
	a.Eat()
}

func main() {
	c := Cat{"ally cat"}
	Compliment(c)
	Feed(c)

	d := Dog{"snoop dogg"}
	Compliment(d)
	Feed(d)

	f := Fish{"slick rick"}
	Compliment(f)
	Feed(f)
}
