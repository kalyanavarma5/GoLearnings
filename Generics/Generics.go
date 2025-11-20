package main

import "fmt"

func main() {
	print("Hello World")
	print(21312)
	print(221.3423423)
}

func print[T any](x T) {
	fmt.Println(x)
}
