package repository

import (
	"database/sql"
	"fmt"
	"git_truongvudinh/go_web/internal/entity"
)

type IUserRepository interface {
	CreateUser(user *entity.User) *entity.User
}

type UserRepository struct {
	db *sql.DB
}

// NewUserRepository constructor
func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) CreateUser(user *entity.User) *entity.User {
	//TODO implement me
	query := `INSERT INTO users (firstname, lastname, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := u.db.Exec(query, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		fmt.Println(err)
	}
	user.ID, err = result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	return user
}
