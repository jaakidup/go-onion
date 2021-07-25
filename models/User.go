package models

type User struct {
	Name    string
	Surname string
	Age     int
}

func NewUser(name, surname string, age int) *User {
	user := User{}
	user.Name = name
	user.Surname = surname
	user.Age = age
	return &user
}
