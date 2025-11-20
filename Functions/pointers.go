package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

// Method with pointer receiver to modify the account balance
func (acc *BankAccount) Deposit(amount float64) {
	acc.Balance += amount
}

func main() {
	var x int = 10 // normal variable
	var p *int     // pointer variable

	p = &x // assign address of x to p

	fmt.Println("Value of x:", x)             // 10
	fmt.Println("Address of x:", &x)          // memory address of x
	fmt.Println("Value of p (address):", p)   // same address as &x
	fmt.Println("Value pointed to by p:", *p) // 10 (dereferenced)

	*p = 20                           // change value at address pointed by p
	fmt.Println("New value of x:", x) // 20 (changed via pointer)

	account := BankAccount{Owner: "Anna", Balance: 100}
	fmt.Println("Before deposit:", account.Balance) // 100

	// Call method with pointer receiver; modifies original account
	account.Deposit(50)

	fmt.Println("After deposit:", account.Balance) // 150

}
