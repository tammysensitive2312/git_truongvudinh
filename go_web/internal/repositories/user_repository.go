package repositories

import (
	"context"
	"errors"
	"fmt"
	"git_truongvudinh/go_web/internal/domain/entities"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	GetUserById(ctx context.Context, ID int, preload bool) (*entities.User, error)
}

type UserRepository struct {
	base
}

func (u UserRepository) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	if err := u.db.WithContext(ctx).Create(user).Error; err != nil {
		log.Error("Cannot create user with err:", err.Error())
		return nil, err
	}
	return user, nil
}

func (u UserRepository) GetUserById(ctx context.Context, ID int, preload bool) (*entities.User, error) {
	var user entities.User
	query := u.db.WithContext(ctx)

	if preload {
		query = query.Preload("Projects")
	}

	if err := query.First(&user, ID).Error; err != nil {
		log.Error("Can not find user with ID: ", ID, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user %v not found", user)
		}
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}
	return &user, nil
}

// NewUserRepository constructor
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{base: base{db: db}}
}
