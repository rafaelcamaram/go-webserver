package main

import (
	"fmt"
)

type user struct {
	ID        int    // Default 0
	FirstName string // Default ''
	LastName  string // Default ''
}

func main() {
	var u user
	u.ID = 1
	u.FirstName = "Rafael"
	u.LastName = "Camara"
	fmt.Println(u)

	u2 := user{
		ID:        2,
		FirstName: "Flavia",
		LastName:  "Gonzaga",
	}
	fmt.Println(u2)
}
