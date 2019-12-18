package errors

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Message string
	Line   int
}

func (err *CustomError) Error() string {
	return err.Message
}

func New(message string, line int) error {
	return &CustomError{message, line}
}

func CreateError(message string, err error) error {
	var errorMessage string

	if err == nil {
		errorMessage = ""
	} else {
		errorMessage = err.Error()
	}

	return errors.New(fmt.Sprintf("%s %s", message, errorMessage))
}
