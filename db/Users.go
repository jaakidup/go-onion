package db

import (
	"fmt"
	"log"

	"github.com/jaakit/go-onion/models"
)

var users []models.User

var jenny = models.User{
	Name:    "jenny",
	Surname: "Rissveds",
	Age:     25,
}

var amy = models.User{
	Name:    "Amy",
	Surname: "Du Plessis",
	Age:     22,
}

func init() {
	log.Println("Init Users DB")
	users = append(users, amy)
	users = append(users, jenny)
}

type Users struct{}

func NewUserDB() *Users {
	userDB := Users{}
	return &userDB
}

func (u *Users) AddUser(user models.User) {
	fmt.Println("Add user to Users", user)
	users = append(users, user)
}

func (u *Users) GetAllUsers() []models.User {
	return users
}
