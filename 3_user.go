package main

import "fmt"

type Notifier interface { // interface normally has postfix as "er"
	Emailer // interface can do embeded also
	Notify()
}
type Emailer interface {
	SendEmail(cc string) (result string)
}
type User struct {
	Name, Email string
}
type Admin struct {
	User
	Level string
} // this is example of embeded "User" to "Admin"
// Embeded sustain class inheritant

func (u User) Notify() {
	fmt.Println("Func User Notify() " + u.Name)
}
func (a Admin) Notify() {
	fmt.Println("Func Admin Notify() " + a.Name)
}

func (u User) SendEmail(cc string) (result string) {
	return cc + " Email User"
}
func (a Admin) SendEmail(cc string) (result string) {
	return cc + " Email Admin"
}

func SendNotify(a Notifier) {
	a.Notify() // this is polymorphism
	r := a.SendEmail("Test")
	fmt.Println("Result from SendNotify", r)
}
func SendEmail(e Emailer) {
	r := e.SendEmail("Test")
	fmt.Println("Result from SendEmail", r)
}

func main() {
	u := User{"A1", "a1@mail.com"}
	fmt.Println(u)
	a := Admin{User{"Admin", "Admin@mail.com"}, "super"} // literal declaration
	fmt.Println(a)
	fmt.Println(a.Name)
	fmt.Println(a.Email)

	fmt.Println("\n   +++Block 1++++ print embeded struct")
	aa := Admin{}
	aa.User.Name = "OAT" // it is optional to put "User"
	aa.Email = "mail"
	aa.Level = "Lv1"

	fmt.Printf("User: % #v\n", u)
	fmt.Printf("Admin: % #v\n", aa)

	fmt.Println("\n   +++Block 2++++ call method and override method func")
	u.Notify()
	User.Notify(u) // Optional form to be called

	// Embeded get the method inheritant from User also, not just struct
	a.Notify()      // use override method of "Admin"
	a.User.Notify() // call original method of "User"

	User{"a", "b"}.Notify()

	fmt.Println("\n   +++Block 3++++ call interface")
	SendNotify(u)
	SendNotify(a)

	fmt.Println()
	SendEmail(u)
	SendEmail(a)
}
