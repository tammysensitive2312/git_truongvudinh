package repositories

import (
	"context"
	"git_truongvudinh/go_web/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository là một mock của IUserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*entity.User), args.Error(1)
}
