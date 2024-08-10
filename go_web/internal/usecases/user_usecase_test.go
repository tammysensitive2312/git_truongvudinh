package usecases

import (
	"context"
	"git_truongvudinh/go_web/internal/domain/dto"
	"git_truongvudinh/go_web/internal/domain/entity"
	"git_truongvudinh/go_web/internal/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateNewUser(t *testing.T) {
	// Khởi tạo mock repository
	mockRepo := new(repositories.MockUserRepository)

	// Khởi tạo service với mock repository
	userService := NewUserService(mockRepo)

	// Tạo request DTO
	request := &dto.CreateUserRequest{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "password123",
	}

	// Tạo một user entity dự kiến trả về
	user := &entity.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "hashedpassword",
	}

	// Thiết lập mock cho CreateUser
	mockRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*entity.User")).Return(user, nil)

	// Gọi phương thức cần test
	result, err := userService.CreateNewUser(context.Background(), request)

	// Kiểm tra kết quả
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.FirstName, result.FirstName)
	assert.Equal(t, user.LastName, result.LastName)
	assert.Equal(t, user.Email, result.Email)

	// Kiểm tra xem mock có được gọi đúng không
	mockRepo.AssertExpectations(t)
}
