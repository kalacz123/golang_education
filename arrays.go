package main

import "fmt"

func main() {

	// var arrayName [size] elementType

	// var numbers [5]int blank initialization
	// fmt.Println(numbers)

	// numbers[0] = 1
	// numbers[2] = 1
	// numbers[4] = 1
	// fmt.Println(numbers)

	numbers := [5]string{"one", "two", "three", "four", "five"}
	fmt.Println(len(numbers))

	// we can't change the size of the array, but we can change the values
	fmt.Println(numbers)
	copiedNumbers := numbers
	copiedNumbers[0] = "six"
	fmt.Println(numbers)
	fmt.Println(copiedNumbers)

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	for _, value := range numbers {
		fmt.Printf("Value: %s\n", value)
	}

	_, b := someFunc() // _ - blank identifier, it's used to ignore the first return value
	fmt.Println(b)

	//Compare two arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	fmt.Println(array1 == array2)

	//multi-dimensional array
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])
		}
	}

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 4

	fmt.Println("Original Array:", originalArray)
	fmt.Println("Copied Array:", copiedArray)

}

func someFunc() (int, int) {
	return 1, 2
}
