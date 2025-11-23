// internal/domain/user.go
package domain

import "time"

// User represents the core user entity
type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate performs business rule validation
func (u *User) Validate() error {
	if u.Email == "" {
		return ErrInvalidEmail
	}
	if len(u.Password) < 8 {
		return ErrPasswordTooShort
	}
	return nil
}
