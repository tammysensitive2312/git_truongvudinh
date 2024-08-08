package repository

import (
	"git_truongvudinh/go_web/internal/entity"
)

/*
initialize a new virtual repository
database = map[id:int64]entity:User
*/

type MockUserRepository struct {
	users map[int64]*entity.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[int64]*entity.User),
	}
}

func (m *MockUserRepository) CreateUser(user *entity.User) *entity.User {
	user.ID = int64(len(m.users) + 1)
	m.users[user.ID] = user
	return user
}
