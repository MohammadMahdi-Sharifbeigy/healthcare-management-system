package entities

import "fmt"

type Error struct {
	Code    string
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func ErrInvalidInput(msg string) Error {
	return Error{
		Code:    "INVALID_INPUT",
		Message: msg,
	}
}

func ErrNotFound(msg string) Error {
	return Error{
		Code:    "NOT_FOUND",
		Message: msg,
	}
}

func ErrConflict(msg string) Error {
	return Error{
		Code:    "CONFLICT",
		Message: msg,
	}
}
