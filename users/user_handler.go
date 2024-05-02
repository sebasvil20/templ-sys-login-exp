package users

import (
	"database/sql"
	"fmt"

	"github.com/sebasvil20/templ-sys-login-exp/models"
)

type IUser interface {
	AddUser(user models.User) error
	GetUsers() ([]models.User, error)
	Authenticate(credentials models.UserCredentials) error
}

type UserHandler struct {
	// users []models.User
	dbInyectado *sql.DB
}

func InitUserHandler(dbAInyectar *sql.DB) UserHandler {
	/*userHandler := UserHandler{}
	userHandler.users = []models.User{
		{Username: "Roger", Email: "roger@roger.com", Password: "password"},
		{Username: "Sebas", Email: "sebas@2.com", Password: "123"},
	}*/

	userHandler := UserHandler{
		dbInyectado: dbAInyectar,
	}

	return userHandler
}

func (userHandlerArg UserHandler) AddUser(user models.User) error {
	// userHandlerArg.users = append(userHandlerArg.users, user)
	fmt.Printf("Adding user: %v\n", user)
	result, err := userHandlerArg.dbInyectado.
		Exec("INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, 'admin')", user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Printf("User added with id: %d\n", id)

	return nil
}

func (userHandlerArg UserHandler) GetUsers() ([]models.User, error) {
	// return userHandlerArg.users, nil

	return nil, nil
}

func (userHandlerArg UserHandler) Authenticate(user models.UserCredentials) error {
	/*for _, u := range userHandlerArg.users {
		if u.Username == user.Username && u.Password == user.Password {
			return nil
		}
	}
	return errors.New("invalid credentials")*/
	return nil
}
