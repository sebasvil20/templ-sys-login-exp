package users

import (
	"errors"

	"github.com/sebasvil20/templ-sys-login-exp/models"
)

type IUser interface {
	AddUser(user models.User) error
	GetUsers() ([]models.User, error)
	Authenticate(credentials models.UserCredentials) error
}

type UserHandler struct {
	users []models.User
}

func InitUserHandler() UserHandler {
	userHandler := UserHandler{}
	userHandler.users = []models.User{
		{Username: "Roger", Email: "roger@roger.com", Password: "password"},
		{Username: "Sebas", Email: "sebas@2.com", Password: "123"},
	}

	return userHandler
}

func (userHandlerArg UserHandler) AddUser(user models.User) error {
	userHandlerArg.users = append(userHandlerArg.users, user)

	return nil
}

func (userHandlerArg UserHandler) GetUsers() ([]models.User, error) {
	return userHandlerArg.users, nil
}

func (userHandlerArg UserHandler) Authenticate(user models.UserCredentials) error {
	for _, u := range userHandlerArg.users {
		if u.Username == user.Username && u.Password == user.Password {
			return nil
		}
	}
	return errors.New("invalid credentials")
}
