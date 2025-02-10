package keeper

import (
	"errors"
)

var ErrDuplicate = errors.New("duplicate")

var ErrNotExist = errors.New("no data")

// // AddURLError is duplicating error.
// type ErrDuplicate struct {
// 	DuplicateValue string
// 	Err            error
// }

// // Error is provided to implement Error interface.
// func (ae *AddURLError) Error() string {
// 	return fmt.Sprintf("%v %v", ae.DuplicateValue, ae.Err)
// }

// // NewDuplicateError return new AddURLError.
// func NewDuplicateError(err error, shortenedURL string) error {
// 	return &AddURLError{
// 		DuplicateValue: shortenedURL,
// 		Err:            err,
// 	}
// }

// var ErrUnexpected = errors.New()
