package main

import "fmt"


func main(){

	fmt.Println(fibonacci(10))
}

func fibonacci(n int) int {
	//Base case
	if n <= 1{
		return n
	}
	//Recursive case
	return fibonacci(n-1) + fibonacci(n-2)
}