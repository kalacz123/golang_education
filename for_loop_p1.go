package main

import "fmt"

func main() {
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// numbers := []string{"one", "two", "three", "four", "five", "six"}
	// for index, number := range numbers {
	// 	fmt.Println(index+1, number)
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Even number: ", i)
	// 	if i == 8 {
	// 		break
	// 	}
	// }

	// rows := 5

	// for i := 1; i <= rows; i++ {
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	for i:=range 10 {
		fmt.Println(10 - i)
	}
	fmt.Println("Done")
}
