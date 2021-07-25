package repository

import (
	"github.com/jaakit/go-onion/db"
	"github.com/jaakit/go-onion/models"
)

type Users struct {
	Name string
	DB   *db.Users
}

func NewUserRepository() *Users {
	userDB := db.NewUserDB()

	userRepository := Users{
		Name: "User Repository",
		DB:   userDB,
	}
	return &userRepository
}

func (u *Users) AddUser(user models.User) {
	u.DB.AddUser(user)
}

func (u *Users) ListUsers() []models.User {
	return u.DB.GetAllUsers()
}
