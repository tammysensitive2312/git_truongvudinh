package usecase

import (
	"testing"
	"time"

	"git_truongvudinh/go_web/internal/entity"
	"git_truongvudinh/go_web/internal/repository"
)

func TestCreateNewUser(t *testing.T) {
	mockRepo := repository.NewMockUserRepository()
	userService := NewUserService(mockRepo)

	user := &entity.User{
		FirstName: "truong",
		LastName:  "vu",
		Email:     "dinhtruong1234lhp@gmail.com",
		Password:  "yeutuyen",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser := userService.CreateNewUser(user)

	if createdUser.ID == 0 {
		t.Errorf("Expected non-zero user ID")
	}
	if createdUser.Email != user.Email {
		t.Errorf("Expected email to be %v, got %v", user.Email, createdUser.Email)
	}
	if createdUser.FirstName != user.FirstName {
		t.Errorf("Expected first name to be %v, got %v", user.FirstName, createdUser.FirstName)
	}
	if createdUser.LastName != user.LastName {
		t.Errorf("Expected last name to be %v, got %v", user.LastName, createdUser.LastName)
	}
	if createdUser.Password == user.Password {
		t.Errorf("Expected hashed password to be different from original")
	}
}
