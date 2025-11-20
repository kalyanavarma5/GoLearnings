package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u User) String() string {
	return fmt.Sprintf("%s (%d)", u.name, u.age)
}

func main() {
	user := User{"Kalyan", 35}
	fmt.Println(user)

	email := EmailNotifier{Email: "user@example.com"}
	sms := SMSNotifier{SMS: "+1234567890"}

	sendNotification(email, "Hello via Email")
	sendNotification(sms, "Hello via SMS")
}

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) Notify(message string) {
	fmt.Println("Email:", e.Email, message)
}

type SMSNotifier struct {
	SMS string
}

func (s SMSNotifier) Notify(message string) {
	fmt.Println("SMS:", s.SMS, message)
}

func sendNotification(n Notifier, msg string) {
	n.Notify(msg)
}
