package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("He said, \"I am great\"") // backslash is used to escape the double quotes
	fmt.Println(`He said, "I am great"`)   // backticks are used to escape the double quotes

	// compile a regex pattern to match email address
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Test some strings
	email1 := "user@email.com"
	email2 := "invalid_email"

	// match the strings
	fmt.Println("Email1: ", re.MatchString(email1))
	fmt.Println("Email2: ", re.MatchString(email2))

	// capturing the groups
	// compile the regex pattern to capture the groups (year-month-day)
	re2 := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	date := "2025-09-25"

	//find all submatches
	submatches := re2.FindStringSubmatch(date)
	fmt.Println("Submatches: ", submatches)
	fmt.Println("Full date: ", submatches[0])
	fmt.Println("Year: ", submatches[1])
	fmt.Println("Month: ", submatches[2])
	fmt.Println("Day: ", submatches[3])

	//Target string

	str := "hello world"

	re3 := regexp.MustCompile(`[aeiou]`)

	// Replace all vowels with *
	result := re3.ReplaceAllString(str, "*")
	fmt.Println(result)

	//i - case insensitive, starts with question mark
	//m - multiline model
	//s - dot matches all

	re4 := regexp.MustCompile(`(?i)go`)

	// Test string
	text := "Go is a programming language"

	// Match the string
	fmt.Println(re4.MatchString(text))

}
