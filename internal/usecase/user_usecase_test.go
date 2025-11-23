// internal/usecase/user_usecase_test.go
package usecase_test

import (
    "context"
    "testing"
    "your-project/internal/domain"
    "your-project/internal/usecase"
    
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// Mock repository
type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *domain.User) error {
    args := m.Called(ctx, user)
    return args.Error(0)
}

func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
    args := m.Called(ctx, email)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*domain.User), args.Error(1)
}

// Test
func TestRegister_Success(t *testing.T) {
    mockRepo := new(MockUserRepository)
    uc := usecase.NewUserUseCase(mockRepo)
    
    user := &domain.User{
        Email:    "test@example.com",
        Name:     "Test User",
        Password: "password123",
    }
    
    mockRepo.On("GetByEmail", mock.Anything, user.Email).Return(nil, domain.ErrUserNotFound)
    mockRepo.On("Create", mock.Anything, user).Return(nil)
    
    err := uc.Register(context.Background(), user)
    
    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
}

