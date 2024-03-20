package models

import "fmt"

type User struct {
	username string
	password string
}

// this function is a method of User struct and (u *User) is a receiver to get User fields

func (u *User) Greeting() {
	fmt.Println("Hello my name is", u.username)
}

func (u *User) SetUserName(username string) {
	u.username = username
}
