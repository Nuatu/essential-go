package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func NewPoint(x, y int) point {
	return point{x, y}
}

type circle struct {
	radius float64
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

type rect struct {
	pos    point
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

func main() {
	p := point{20, 40}
	fmt.Println(p.x)

	r := rect{
		pos:    NewPoint(20, 40),
		width:  200,
		height: 400,
	}

	c := circle{
		radius: 2.5,
	}
	fmt.Println("rect height: ", r.height)
	fmt.Println("rect width: ", r.width)
	fmt.Println("rect area: ", r.area())
	fmt.Println("circle radius: ", c.radius)
	fmt.Println("circle diameter: ", c.diameter())

}
