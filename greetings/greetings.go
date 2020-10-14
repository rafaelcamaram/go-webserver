package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func Hello(name string) (string, error) {
	// If no name was given, return error with a message
	if name == "" {
		return "", errors.New("Empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil;
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
