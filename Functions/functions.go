package main

import "fmt"

func addition(a int, b int) int {
	return a + b
}
func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	var addition2 = addition(1, 33)
	fmt.Println("addition2", addition2)
}
