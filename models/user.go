package models

import (
	"errors"
	"fmt"
)

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
	if u.ID != 0 {
		return User{}, errors.New("new user must not include id")
	}

	u.ID = nextID
	nextID++
	users = append(users, &u)

	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, current_user := range users {
		if current_user.ID == id {
			return *current_user, nil
		}
	}

	return User{}, fmt.Errorf("user with ID '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, current_user := range users {
		if current_user.ID == u.ID {
			users[i] = &u

			return u, nil
		}
	}

	return User{}, fmt.Errorf("user with id '%v' not found", u.ID)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("user with id '%v' not found", id)
}
