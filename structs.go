package main


import "fmt"


type Person struct {
	firstName string
	lastName  string
	age       int
	address Address
	PhoneHomeCell
}
type Address struct {
	city string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) updateAge(newAge int) {
	p.age = newAge
}



func main() {
	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       20,
		address: Address{
			city: "New York",
			country: "USA",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "1234567890",
			cell: "0987654321",
		},
	}

	p2 := Person{
		firstName: "Jane",
		age:       25,
	}

	p3 := Person{
		firstName: "Jane",
		age:       25,
	}


	fmt.Println(p1)
	fmt.Println(p1.home)
	fmt.Println(p2)

	fmt.Println("are p3 and p2 equal?", p3 == p2)

	// Anonymous struct
	user := struct {
		userName string
		email string
	}{
		userName: "john123",
		email: "john123@gmail.com",
	}

	fmt.Println(user)

	// functions with structs
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())

	p1.updateAge(21)
	fmt.Println(p1.age)
	fmt.Println(p1)
}
