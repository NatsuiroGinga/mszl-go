package main

import "fmt"

type User struct {
	id   int
	name string
}

type Manager struct {
	User
}

func (u *User) ToString() string {
	return fmt.Sprintf("User: %p %v", u, u)
}

func main() {
	manager := Manager{User{
		id:   0,
		name: "zjj",
	}}
	fmt.Println("Manager: ", manager)
	fmt.Println("User: ", manager.ToString())
}
