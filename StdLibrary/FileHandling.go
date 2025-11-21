package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create text file
	content := "Hello, this is some sample text in the file."
	// Create or overwrite the file
	err := os.WriteFile("example.txt", []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
	// Read From a Text File
	data, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File content:", string(data))
	// append data to existing text file
	f, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString("\nAppended line."); err != nil {
		log.Fatal(err)
	}
	println("Data appended to file")
}
