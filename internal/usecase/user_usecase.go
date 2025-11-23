// internal/usecase/user_usecase.go
package usecase

import (
	"context"
	"fitness-app-go/internal/domain"
	"fitness-app-go/internal/repository"
)

// UserUseCase defines the interface for user business logic
type UserUseCase interface {
	Register(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, id int64) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id int64) error
	ListUsers(ctx context.Context, limit, offset int) ([]*domain.User, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
}

// NewUserUseCase creates a new user use case
func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u *userUseCase) Register(ctx context.Context, user *domain.User) error {
	// Business logic: validate user
	if err := user.Validate(); err != nil {
		return err
	}

	// Check if email already exists
	existing, err := u.userRepo.GetByEmail(ctx, user.Email)
	if err == nil && existing != nil {
		return domain.ErrDuplicateEmail
	}

	// Hash password (you'd use bcrypt here)
	// user.Password = hash(user.Password)

	// Create user
	return u.userRepo.Create(ctx, user)
}

func (u *userUseCase) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	return u.userRepo.GetByID(ctx, id)
}

func (u *userUseCase) UpdateUser(ctx context.Context, user *domain.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	return u.userRepo.Update(ctx, user)
}

func (u *userUseCase) DeleteUser(ctx context.Context, id int64) error {
	return u.userRepo.Delete(ctx, id)
}

func (u *userUseCase) ListUsers(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	return u.userRepo.List(ctx, limit, offset)
}
