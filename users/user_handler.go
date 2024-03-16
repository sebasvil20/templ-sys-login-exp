package users

type User struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserCredentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
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
