package main

import "fmt"

func main() {

	// Printing functions
	// fmt.Print("Hello")
	// fmt.Println("World!")
	// fmt.Printf("Name: %s, Age: %d\n", "John", 20)

	// Formatting functions
	// s := fmt.Sprint("Hello", "World!", 123, 456)
	// fmt.Println(s)
	// s = fmt.Sprintln("Hello", "World!", 123, 456)
	// fmt.Println(s)
	// s = fmt.Sprintf("Name: %s, Age: %d\n", "John", 20)
	// fmt.Println(s)

	// Scanning functions
	// var name string
	// var age int
	// fmt.Scan(&name, &age)
	// fmt.Scanln(&name, &age)
	// fmt.Scanf("%s %d", &name, &age)

	// Error formatting functions
	err := checkAge(19)
	if err != nil {
		fmt.Println(err)
	}

}
func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age: %d is less than 18", age)
	}
	return nil
}
