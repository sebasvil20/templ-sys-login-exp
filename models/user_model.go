package models

type User struct {
	Username string `schema:"username" validate:"required"`
	Email    string `schema:"email" validate:"required"`
	Password string `schema:"password" validate:"required"`
}

type UserCredentials struct {
	Username string `schema:"username" validate:"required"`
	Password string `schema:"password" validate:"required"`
}
