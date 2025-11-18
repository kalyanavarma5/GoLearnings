package main

import "fmt"

type name struct {
	UserName string
	PassWord string
	Email    string
	IsActive bool
}

func main() {

	user1 := name{
		UserName: "Kalyan",
		PassWord: "123456",
		Email:    "Kalyan",
		IsActive: true,
	}
	fmt.Println(user1)
}
