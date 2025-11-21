package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"time"
)

func main() {
	// flag package
	name := flag.String("name", "World", "name to greet")
	phone := flag.String("phone", "123456", "phone number to greet")
	flag.Parse()
	fmt.Printf("Hello, %s!\n", *name)
	fmt.Printf("Hello, %s!\n", *phone)

	// time package
	now := time.Now()
	fmt.Println("Current time:", now.Format(time.RFC1123))

	duration := 2 * time.Second
	fmt.Println("Sleeping for", duration)
	time.Sleep(duration)
	fmt.Println("Awake now")

	// Encoding/JSON package

	u := User{Name: "Alice", Email: "alice@example.com"}
	jsonData, _ := json.Marshal(u)
	fmt.Println(string(jsonData))

	var u2 User
	json.Unmarshal(jsonData, &u2)
	fmt.Println(u2.Name, u2.Email)

	// OS Package
	file, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("Hello from Go os package")

	// Bufio Package
	file2, _ := os.Open("example.txt")
	defer file2.Close()

	scanner := bufio.NewScanner(file2)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// SLOG PACKAGE
	logger := slog.Default()
	logger.Info("Starting application", "version", "1.0")
	logger.Debug("version", "version")

	// regexp

	re := regexp.MustCompile(`hello`)
	fmt.Println(re.MatchString("hello world"))            // true
	fmt.Println(re.ReplaceAllString("hello world", "hi")) // hi world
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
