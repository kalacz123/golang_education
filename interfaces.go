package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) perimeter() float64 {
	return 2 * c.radius * math.Pi
}

func (c Circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}


func main() {
	r := Rectangle{width: 10, height: 5}
	c := Circle{radius: 5}

	measure(r)
	measure(c)

	myPrinter(1, "Hello", true)

}

func myPrinter( i ... interface{}) {
	for _, v := range i {
		s := fmt.Sprintf("Type: %T, Value: %v", v, v)
		fmt.Println(s)
	}
}