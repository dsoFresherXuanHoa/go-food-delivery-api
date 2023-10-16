package exceptions

import "errors"

var (
	ErrUserNotFound = errors.New("user don't exist: check your username or password and try again")
)
