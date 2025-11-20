package main

import "fmt"

type Order struct {
	Amount  float64
	TaxRate float64
}

// Method with receiver 'o' of type Order
func (o Order) TotalPrice() float64 {
	return o.Amount + (o.Amount * o.TaxRate / 100)
}

type User struct {
	Name     string
	IsActive bool
}

// Method to activate user account
func (u *User) Activate() {
	u.IsActive = true
}

// Method to deactivate user account
func (u *User) Deactivate() {
	u.IsActive = false
}

func main() {
	order := Order{Amount: 100, TaxRate: 10}
	fmt.Println("Total Price:", order.TotalPrice()) // Output: 110

	user := User{Name: "Emma", IsActive: false}
	fmt.Println("Account active?", user.IsActive) // false

	user.Activate()
	fmt.Println("Account active?", user.IsActive) // true
}
