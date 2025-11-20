package main

import (
	"fmt"
	"reflect"
)

type name struct {
	UserName string
	PassWord string
	Email    string
	IsActive bool
}

// Base struct
type Person struct {
	Name string
	Age  int
}

// Method on Person
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

// Embedding Person into Employee
type Employee struct {
	Person     // Embedded struct
	EmployeeID string
}

func main() {

	user1 := name{
		UserName: "Kalyan",
		PassWord: "123456",
		Email:    "Kalyan",
		IsActive: true,
	}

	type Programmer struct {
		FirstName string
		LastName  string
		Language  string `help:"The primary programming language used by the programmer"`
	}

	programmer := Programmer{
		FirstName: "Tomas",
		LastName:  "Svojanovsky",
		Language:  "Go",
	}

	// Use reflection to retrieve metadata about the 'language' field.
	if member, ok := reflect.TypeOf(programmer).FieldByName("Language"); ok {
		// If the 'language' field exists, print the value of the 'help' struct tag.
		fmt.Printf("The field language means: %s", member.Tag.Get("help"))
	} else {
		// If the 'language' field is not found, print a fallback message.
		fmt.Println("Unknown language")
	}

	fmt.Println(user1)

	e := Employee{
		Person:     Person{Name: "Alice", Age: 30},
		EmployeeID: "E12345",
	}

	// Access embedded struct fields directly
	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)

	// Call method from embedded struct
	e.Greet()

	// Access Employee's own fields
	fmt.Println("Employee ID:", e.EmployeeID)
}
