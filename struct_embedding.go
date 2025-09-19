package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	person     // Anonymous Embedded struct like a field, it can be also named
	employeeID string
	salary     float64
}

func (p person) introduce(){
	fmt.Printf("Hello, I'm %s and I'm %d years old.\n", p.name, p.age)
}

func (e employee) introduce(){
	fmt.Printf("Hello, I'm %s and I'm %d years old. My employee ID is %s and my salary is %f.\n", e.person.name, e.person.age, e.employeeID, e.salary)
}

func main() {
	emp := employee{
		person: person{name: "John", age: 30},
		employeeID: "123456",
		salary: 100000,
	}

	fmt.Println(emp.name) // embedded struct fields are promoted to the outer struct
	fmt.Println(emp.age)
	fmt.Println(emp.employeeID)
	fmt.Println(emp.salary)
	emp.introduce()
	emp.person.introduce()
}
