package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func hello2(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func hello4(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {

	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)

	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")

	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")

	done3 := make(chan bool)
	go hello2(done3)
	<-done3
	fmt.Println("main is called")
	* /
		done5 := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello4(done5)
	<-done5
	fmt.Println("Main received data")
}
