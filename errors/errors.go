package errors

import (
	"errors"
	"fmt"
)

func CreateError(message string, err error) error {
	var errorMessage string

	if err == nil {
		errorMessage = ""
	} else {
		errorMessage = err.Error()
	}

	return errors.New(fmt.Sprintf("%s %s", message, errorMessage))
}
