package strings

import (
	"errors"
	"regexp"
)

// ValidateEmail will return an error if the email is invalid
func ValidateEmail(email string) error {
	pattern, err := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if err != nil {
		return errors.New("Failed to compile regex")
	}

	if valid := pattern.MatchString(email); !valid {
		return errors.New("Please provide a valid email address")
	}

	return nil
}
