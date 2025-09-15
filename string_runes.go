package main

import (
	"fmt"
)

// String is a sequence of bytes that represents a sequence of characters
// Rune is an integer value that represents a character

func main() {

	message := "Hello, \nGo!"    // normal string, it's used to store a string with the escape characters
	message2 := "Hello, \tGo!"   // normal string, it's used to store a string with the tab character
	message3 := "Hello, \rGo!"   // normal string, it's used to store a string with the carriage return character
	rawMessage := `Hello, \nGo!` // raw string, it's used to store a string without the escape characters

	fmt.Println(message)
	fmt.Println(message2)
	fmt.Println(message3) // Go!lo - takes the Hello, and replaces it with Go!
	fmt.Println(rawMessage)

	fmt.Println(len(message))    // \n is 1 character, so it's 11 characters
	fmt.Println(len(rawMessage)) // because it's a raw string, it's 12 characters

	fmt.Println(message[0]) // ASCII value of H is 72

	greeting := "Hello"
	name := "Bohdan"

	fmt.Println(greeting + ", " + name) // concatenation

	result := fmt.Sprintf("%s, %s", greeting, name) // formatting - it's a better way to concatenate strings
	fmt.Println(result)

	fmt.Println(getSum(greeting))
	fmt.Println(getSum(name))
	// comparing strings
	fmt.Println(greeting == name)
	fmt.Println(greeting != name)
	fmt.Println(greeting < name) // ASCII value of H is 72, ASCII value of B is 66, so 72 < 66 is false
	fmt.Println(greeting > name) // ASCII value of H is 72, ASCII value of B is 66, so 72 > 66 is true
	fmt.Println(greeting > name)
	fmt.Println(greeting <= name)
	fmt.Println(greeting >= name)

    // itaration over strings
	for i, v := range greeting {
		fmt.Printf("Index: %d, Rune: %c\n", i, v) //%d - number, %c - character, %s - string, %v - value, %t - boolean, %f - float, %q - quoted string, %x - hexadecimal, %X - hexadecimal uppercase, %o - octal, %O - octal uppercase, %b - binary, %B - binary uppercase, %e - exponential notation, %E - exponential notation uppercase, %g - general, %G - general uppercase, %p - pointer, %T - type, %v - value, %#v - value with type, %#x - hexadecimal with type, %#X - hexadecimal uppercase with type, %#o - octal with type, %#O - octal uppercase with type, %#b - binary with type, %#B - binary uppercase with type, %#e - exponential notation with type, %#E - exponential notation uppercase with type, %#g - general with type, %#G - general uppercase with type, %#p - pointer with type, %#T - type with value
	}

	const NIHONGO = "日本語" // Japanese
	fmt.Println(NIHONGO)
	for i, v := range NIHONGO {
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}

}
func getSum(message string) int {
	sum := 0
	for i := 0; i < len(message); i++ {
		sum += int(message[i])
	}
	return sum
}