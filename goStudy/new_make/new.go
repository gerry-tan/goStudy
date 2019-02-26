package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type UserInfo interface {
	GetUser() *User
}

func (user *User) GetUser() *User {
	return &User{
		Name: "Tomy",
		Age:  23,
	}
}

func main() {
	user := new(User)
	user.Name = "Jame"

	fmt.Println(user)

	var userInfo UserInfo
	userInfo = new(User)
	user = userInfo.GetUser()

	fmt.Println(user)
}
