package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name was empty")
	}

	// Return a greeting that embeds the name in a message
	return fmt.Sprintf("Hi, %v. Welcome!", name), nil
}
