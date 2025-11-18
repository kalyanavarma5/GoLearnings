package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["route"] = 66
	m["route2"] = 666
	m["route3"] = 66666
	fmt.Println(m)

	m2 := make(map[string]string)
	m2["route"] = "66"
	m2["route2"] = "666"
	m2["route3"] = "66666"
	fmt.Println(m2)

	v := m["route"]
	fmt.Println(v)

	delete(m, "route")
	fmt.Println(v)

	ages := make(map[string]int)

	// Add key-value pairs
	ages["Alice"] = 30
	ages["Bob"] = 25

	// Update a value
	ages["Alice"] = 31

	// Access and print values
	fmt.Println("Alice's age:", ages["Alice"])

	// Check if a key exists
	if age, exists := ages["Charlie"]; exists {
		fmt.Println("Charlie's age:", age)
	} else {
		fmt.Println("Charlie not found")
	}

	// Iterate over the map
	fmt.Println("All ages:")
	for name, age := range ages {
		fmt.Printf("%s: %d\n", name, age)
	}
	// Delete a key
	delete(ages, "Bob")
	fmt.Println("After deletion:", ages)

	m23 := map[string]int{"a": 1, "b": 2}
	v, ok := m23["a"]
	if ok {
		fmt.Println("Key exists with value:", v)
	} else {
		fmt.Println("Key not found")
	}
}
