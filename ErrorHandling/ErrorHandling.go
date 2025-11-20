package main

import (
	"fmt"
	"runtime/debug"
)

type MyError struct {
	Msg string
}

func (e MyError) Error() string {
	return e.Msg
}

func main() {
	//data, err := ioutil.ReadFile("file.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(data))
	//err2 := errors.New("something went wrong")
	//fmt.Println(err2)

	////err3 := fmt.Errorf("failed to open file: %v", "King")
	////fmt.Println(err3)
	//
	//fmt.Errorf("connection to server %s failed: %w", "", err)
	//
	//var ErrNotFound = errors.New("not found")
	//if errors.Is(err, ErrNotFound) {
	//	// handle
	//}

	causePanic()

}

func mayPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("something bad happened")
}

func causePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			debug.PrintStack()
		}
	}()
	panic("unexpected error")
}
