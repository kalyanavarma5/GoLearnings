package main

import "fmt"

func addition(a int, b int) int {
	return a + b
}
func swap(x string, y string) (string, string) {
	return y, x
}

func greet(name string) string {
	return name
}

func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0.0, "Error by divide by 0"
	}
	return a / b, "Success"
}

func multiplication(a int, b int) int {
	return a * b
}

func getUserInfo() (name string, age int) {
	name = "John Doe"
	age = 30
	return // No need to explicitly return variables
}

// Variadic function to sum numbers
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// swap tries to swap two integers (call by value)
func swap2(x, y int) {
	temp := x
	x = y
	y = temp
	fmt.Println("Inside swap:", x, y) // Values are swapped here
}

func main() {
	var addition2 = addition(1, 33)
	fmt.Println("addition2", addition2)
	fmt.Println(greet("Gopher"))
	fmt.Println(divide(100, 20))
	fmt.Println(multiplication(10, 20))
	fmt.Println(getUserInfo())
	fmt.Println(sum(1, 2, 3, 4, 5))      // Output: 15
	fmt.Println(swap("Kalyan", "Varma")) // Output: 15

	// Anonymous function assigned to a variable
	sayHello := func(name string) {
		fmt.Println("Hello,", name)
	}
	sayHello("Alice") // Output: Hello, Alice

	increment := counter()
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2

	a, b := 5, 10
	fmt.Println("Before swap:", a, b) // 5 10
	swap2(a, b)
	fmt.Println("After swap:", a, b) // 5 10 (no change)
}
