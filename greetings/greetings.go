package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name was empty")
	}

	// Return a greeting that embeds the name in a message
	return fmt.Sprintf(randomFormat(), name), nil
}

// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	formats := []string{
		"Great to see you %v!",
		"Hi, %v. Welcome, yo!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
