package usecase

import (
	_ "crypto/md5"
	"git_truongvudinh/go_web/internal/common"
	"git_truongvudinh/go_web/internal/entity"
	"git_truongvudinh/go_web/internal/repository"
)

type IUserService interface {
	CreateNewUser(user *entity.User) (u *entity.User)
}

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u UserService) CreateNewUser(user *entity.User) (r *entity.User) {
	hashPassword := common.HashPassword(user.Password)
	newUser := &entity.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashPassword,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	_ = u.userRepository.CreateUser(newUser)

	return newUser
}
