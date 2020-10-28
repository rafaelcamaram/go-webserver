package models

// User is a representation of a user
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User // Slice with pointers to User objects
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)

	return u, nil
}
