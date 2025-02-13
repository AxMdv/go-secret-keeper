package keeper

import (
	"errors"
	"fmt"
)

var ErrDuplicate = errors.New("duplicate")

var ErrNotExist = errors.New("no data")

// ErrUnexpected is unexpected error.
type UnexpectedError struct {
	Message string
	Err     error
}

// Error is provided to implement Error interface.
func (u *UnexpectedError) Error() string {
	return fmt.Sprintf("%v: %v", u.Message, u.Err)
}

// NewDuplicateError return new AddURLError.
func NewUnexpectedError(err error, message string) error {
	return &UnexpectedError{
		Message: "unexpected error",
		Err:     err,
	}
}
