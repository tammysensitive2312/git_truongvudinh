package usecases

import (
	"context"
	_ "crypto/md5"
	"git_truongvudinh/go_web/internal/common"
	"git_truongvudinh/go_web/internal/domain/dto"
	"git_truongvudinh/go_web/internal/domain/entity"
	"git_truongvudinh/go_web/internal/repositories"
)

type IUserService interface {
	CreateNewUser(ctx context.Context, request *dto.CreateUserRequest) (*entity.User, error)
}

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) IUserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u UserService) CreateNewUser(ctx context.Context, request *dto.CreateUserRequest) (*entity.User, error) {
	hashPassword := common.HashPassword(request.Password)
	newUser := request.ToUserEntity()
	newUser.Password = hashPassword
	userRsp, err := u.userRepository.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}
	return userRsp, nil
}
