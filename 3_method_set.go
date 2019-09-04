package main

import "fmt"

type User struct {
	Name, Email string
}

func (u *User) Notify() { // Notify is pointer
	fmt.Println("User notify na ka")
}

func (u User) SendEmail() { // SendEmail is by value
	fmt.Println("User send email")
}

type Notifier interface {
	Emailer
	Notify()
}

type Emailer interface {
	SendEmail()
}

func main() {
	u := User{"OAT", "mail"}
	SendNotify(&u)   // put with '&' will make interface can support both (u User) and (u *User) method
	SendToEmailNa(u) // put with normal opject will make interface can support only (u User)
}

func SendNotify(n Notifier) {
	n.Notify()
}

func SendToEmailNa(e Emailer) {
	e.SendEmail()
}
