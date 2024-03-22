package users

type User struct {
	Username string `schema:"username" validate:"required"`
	Email    string `schema:"email" validate:"required"`
	Password string `schema:"password" validate:"required"`
}

type UserCredentials struct {
	Username string `schema:"username" validate:"required"`
	Password string `schema:"password" validate:"required"`
}

var users []User

func AddUser(user User) {
	users = append(users, user)
}

func GetUsers() []User {
	return users
}

func Authenticate(user UserCredentials) bool {
	for _, u := range users {
		if u.Username == user.Username && u.Password == user.Password {
			return true
		}
	}
	return false
}
