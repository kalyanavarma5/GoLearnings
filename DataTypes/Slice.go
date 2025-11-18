package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "orange"}
	fmt.Println(fruits)
	fruits = append(fruits, "banana")
	fruits = append(fruits, "BlueMoon")
	fmt.Println(fruits)

	numbers := make([]int, 5)
	fmt.Println(numbers)

	numbers2 := make([]int, len(fruits), len(fruits))
	fmt.Println(numbers2)

	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:]
	fmt.Println(slice)

	arr2 := [5]int{10, 20, 30, 40, 50}
	// Create a slice from the array
	slice2 := arr2[1:4]
	// Print the array and the slice
	fmt.Println("Array:", arr2)
	fmt.Println("Slice:", slice2)

}
