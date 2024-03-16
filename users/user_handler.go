package users

type User struct {
	Username string
	Password string
}

var users []User

func AddUser(user User) {
	users = append(users, user)
}

func GetUsers() []User {
	return users
}

func Authenticate(user User) bool {
	for _, u := range users {
		if u.Username == user.Username && u.Password == user.Password {
			return true
		}
	}
	return false
}
