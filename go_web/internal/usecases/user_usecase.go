package usecases

import (
	"context"
	"git_truongvudinh/go_web/internal/common"
	"git_truongvudinh/go_web/internal/domain/dto"
	"git_truongvudinh/go_web/internal/domain/entity"
	"git_truongvudinh/go_web/internal/repositories"
	"time"

	log "github.com/sirupsen/logrus"
)

type IUserService interface {
	Create(ctx context.Context, request dto.UserCreatable) (*entity.User, error)
	GetUserByID(ctx context.Context, ID int) (*entity.User, error)
}

type UserService struct {
	userRepository repositories.IUserRepository
}

func (u UserService) Create(ctx context.Context, request dto.UserCreatable) (*entity.User, error) {
	newUser := request.ToUserEntity()

	hashedPassword := common.HashPassword(newUser.Password)
	newUser.Password = hashedPassword

	now := time.Now()
	newUser.CreatedAt = now
	newUser.UpdatedAt = now

	data, err := u.userRepository.CreateUser(ctx, newUser)
	if err != nil {
		log.Error("Failed to create user:", err)
		return nil, err
	}

	log.Info("User created successfully")
	return data, nil
}

func (u UserService) GetUserByID(ctx context.Context, ID int) (*entity.User, error) {
	user, err := u.userRepository.GetUserById(ctx, ID)
	if err != nil {
		log.Error("Failed to get user by ID:", err)
		return nil, err
	}

	return user, nil
}

func NewUserService(userRepository repositories.IUserRepository) IUserService {
	return &UserService{
		userRepository: userRepository,
	}
}
