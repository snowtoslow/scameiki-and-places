package benches

import "errors"

var (
	ErrNotExist = errors.New("User does not exist")
	ErrInvalidId = errors.New("Invalid ID")
	ErrInvalidBody = errors.New("Invalid request body")
	ErrInsertionFailed = errors.New("Error in  user insertion")
	ErrUpdateFailed = errors.New("Error in user update")
	ErrDeleteFailed = errors.New("Error in user delete")
)
