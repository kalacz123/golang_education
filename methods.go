package main

import "fmt"

type Rectangle struct {
	length float64
	width float64
}

type Shape struct{
	Rectangle
}

// Method with value receiver, when we want something to calculate but not change the value of the struct
func (r Rectangle) area() float64 {
	return r.length * r.width
}

// Method with pointer receiver, when we want to change the value of the struct
func (r *Rectangle) scale(factor float64) {
	r.length *= factor
	r.width *= factor

}
type MyInt int

//Method on a user-defined type
func (m MyInt) isPositive() bool {
	return m > 0
}

func (MyInt) welcome() string {
	return "Welcome to MyInt type"
}

func main() {
	rect := Rectangle{length: 10, width: 5}
	fmt.Println(rect.area())
	rect.scale(2)
	fmt.Println(rect.area())

	num := MyInt(-5)
	num1 := MyInt(5)
	fmt.Println(num.isPositive())
	fmt.Println(num1.isPositive())
	fmt.Println(num.welcome())

	shape := Shape{Rectangle{length: 10, width: 5}}
	fmt.Println(shape.area())


}