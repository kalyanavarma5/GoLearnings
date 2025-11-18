package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	simpleMessage := "Hello World Gotham"
	fmt.Println("This is simpleMessage  text", simpleMessage)
	rawMessage := `This is a raw string literal.
    It can span multiple lines
    without escape sequences like \n.
	Helooo
	polo`
	fmt.Println("This is raw string text", rawMessage)

	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName
	fullName2 := fullName + lastName + firstName
	fmt.Println(fullName2)
	fmt.Println(fullName)

	// String Conversions
	age := 30
	ageStr := fmt.Sprintf("Age: %d", age)
	fmt.Println(ageStr)

	score := 95.5
	scoreStr := strconv.FormatFloat(score, 'f', 2, 64)
	fmt.Println("Score: " + scoreStr)

	// String length

	text := "Hello, 你好"
	length := len(text)                        // Incorrect length
	runeLength := utf8.RuneCountInString(text) // Correct length
	fmt.Printf("Incorrect Length: %d\n", length)
	fmt.Printf("Correct Length: %d\n", runeLength)

}
