package models

// User is a representation of a user
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	user   []*User // Slice with pointers to User objects
	nextID = 1
)
