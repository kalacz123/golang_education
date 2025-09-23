package main

import (
	"fmt"
	// "regexp"
	// "strconv"
	"strings"
	// "unicode/utf8"
)

func main() {
	// str := "Hello Go!"

	// fmt.Println(len(str)) // length of the string

	// str1 := "hello "
	// str2 := "world"
	// result := str1 + str2 // concatenation of two strings
	// fmt.Println(result)

	// fmt.Println(str[0])   // index of the string
	// fmt.Println(str[1:5]) // substring of the string

	// //string conversion
	// num := 18
	// str3 := strconv.Itoa(num)
	// fmt.Println(len(str3))

	// //split string and store in an array
	// fruits := "apple, banana, orange"
	// fruits1 := "apple-banana-orange"
	// parts := strings.Split(fruits, ", ")
	// parts1 := strings.Split(fruits1, "-")
	// fmt.Println(parts)
	// fmt.Println(parts1)

	// //join array into a string
	// countries := []string{"Germany", "France", "Italy"}
	// joined := strings.Join(countries, ", ")
	// fmt.Println(joined)

	// // check if a string contains a substring
	// fmt.Println(strings.Contains(str, "Go"))

	// //replace a substring with another substring
	// replaced := strings.Replace(str, "Go", "Everybody", 1)
	// fmt.Println(replaced)

	// //trim spaces from a string at the beginning and the end
	// strWithSpaces := " Hello Everyone ! "
	// fmt.Println(strWithSpaces)
	// fmt.Println(strings.TrimSpace(strWithSpaces))

	// //convert a string to lowercase and uppercase
	// fmt.Println(strings.ToLower(strWithSpaces))
	// fmt.Println(strings.ToUpper(strWithSpaces))

	// //repeat a string
	// fmt.Println(strings.Repeat("Ha", 3))

	// //count the number of occurrences of a substring in a string
	// fmt.Println(strings.Count(str, "o"))

	// //check if a string has a prefix and a suffix
	// fmt.Println(strings.HasPrefix("Hello", "he"))
	// fmt.Println(strings.HasSuffix("Hello", "lo"))

	// //Regular expressions. It's used to search for a pattern in a string.
	// // . - any character
	// // \d - any digit
	// // \w - any word character
	// // \s - any whitespace character
	// // [abc] - any of the characters in the brackets
	// // [0-9] - any digit from 0 to 9
	// // [a-z] - any lowercase letter from a to z

	// // +     - one or more repetitions
	// // *     - zero or more repetitions
	// // ?     - zero or one repetition
	// // {3}   - exactly 3 repetitions
	// // {2,5} - from 2 to 5 repetitions

	// //^     - start of the string
	// //$     - end of the string

	// str5 := "Hello, 123 Go!"
	// re := regexp.MustCompile(`\d+`)
	// result1 := re.FindAllString(str5, -1) // returns all the matches in an array
	// fmt.Println(result1)

	// str6 := "Hello, 你好"
	// fmt.Println(utf8.RuneCountInString(str6)) // returns the number of runes in the string

	// STRING BUILDER
	var builder strings.Builder

	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")
	// convert builder to string
	result := builder.String()
	fmt.Println(result)

	// Using WriteRune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you?")
	result = builder.String()
	fmt.Println(result)

	// Reset the builder
	builder.Reset()
	builder.WriteString("Hello, starting fresh!")
	fmt.Println(builder.String())

}
