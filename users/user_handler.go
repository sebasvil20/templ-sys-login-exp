package users

import "errors"

type User struct {
	Username string `schema:"username" validate:"required"`
	Email    string `schema:"email" validate:"required"`
	Password string `schema:"password" validate:"required"`
}

type UserCredentials struct {
	Username string `schema:"username" validate:"required"`
	Password string `schema:"password" validate:"required"`
}

type IUser interface {
	AddUser(user User) error
	GetUsers() ([]User, error)
	Authenticate(credentials UserCredentials) error
}

type UserHandler struct {
	users []User
}

func InitUserHandler() UserHandler {
	userHandler := UserHandler{}
	userHandler.users = []User{
		{Username: "Roger", Email: "roger@roger.com", Password: "password"},
		{Username: "Sebas", Email: "sebas@2.com", Password: "123"},
	}

	return userHandler
}

func (userHandlerArg UserHandler) AddUser(user User) error {
	userHandlerArg.users = append(userHandlerArg.users, user)

	return nil
}

func (userHandlerArg UserHandler) GetUsers() ([]User, error) {
	return userHandlerArg.users, nil
}

func (userHandlerArg UserHandler) Authenticate(user UserCredentials) error {
	for _, u := range userHandlerArg.users {
		if u.Username == user.Username && u.Password == user.Password {
			return nil
		}
	}
	return errors.New("invalid credentials")
}
