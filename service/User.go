package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jaakit/go-onion/repository"
)

type User struct {
	Name       string
	Repository repository.Users
}

func (s *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	users := s.Repository.ListUsers()
	fmt.Println(users)
	jenc := json.NewEncoder(w)
	jenc.Encode(users)

}

func NewUserService() *User {

	userRepository := repository.Users{}
	userService := User{
		Name:       "User Service",
		Repository: userRepository,
	}

	return &userService
}
