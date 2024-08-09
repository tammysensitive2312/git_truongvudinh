package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"git_truongvudinh/go_web/internal/domain/entity"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

type UserRepository struct {
	db *sql.DB
}

// NewUserRepository constructor
func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	query := `INSERT INTO users (firstname, lastname, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`

	// Sử dụng ExecContext thay vì Exec
	result, err := u.db.ExecContext(ctx, query, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	user.ID = id
	return user, nil
}
