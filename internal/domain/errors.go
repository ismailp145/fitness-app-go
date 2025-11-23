// internal/domain/errors.go
package domain

import "errors"

var (
    ErrUserNotFound      = errors.New("user not found")
    ErrInvalidEmail      = errors.New("invalid email")
    ErrPasswordTooShort  = errors.New("password must be at least 8 characters")
    ErrDuplicateEmail    = errors.New("email already exists")
)

