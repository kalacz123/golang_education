package main

import "fmt"

func main() {

	num := 42
	fmt.Printf("%05d\n", num) // 00042

	message := "o"
	fmt.Printf("|%10s|\n", message) // |         o|
	fmt.Printf("|%-10s|\n", message) // |o         |

	message1 := "Hello \nWorld!" // prints Hello and World! in separate lines
	message2 := `Hello \nWorld!` // prints Hello \nWorld! as it is
	fmt.Println(message1)
	fmt.Println(message2)

	sqlQuery := `SELECT * FROM users WHERE age > 30`
	fmt.Println(sqlQuery)
}
