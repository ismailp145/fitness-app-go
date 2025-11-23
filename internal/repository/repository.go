// internal/repository/repository.go
package repository

import (
	"context"
	"fitness-app-go/internal/domain"
)

// UserRepository defines the interface for user data operations
// This interface is defined in the inner layer but implemented in outer layer
type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*domain.User, error)
}
