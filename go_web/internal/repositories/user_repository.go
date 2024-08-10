package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"git_truongvudinh/go_web/internal/domain/entity"
	"time"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserById(ctx context.Context, ID int64) (*entity.User, error)
}

type UserRepository struct {
	db *sql.DB
}

// NewUserRepository constructor
func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) GetUserById(ctx context.Context, ID int64) (*entity.User, error) {
	query := `SELECT id, firstname, lastname, email, password, created_at, updated_at FROM users WHERE id = ?`
	row := u.db.QueryRowContext(ctx, query, ID)

	// Tạo một biến User để lưu kết quả
	var user entity.User
	var createdAtStr, updatedAtStr string

	// Scan dữ liệu từ hàng trả về vào struct User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &createdAtStr, &updatedAtStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with ID %d not found", ID)
		}
		return nil, fmt.Errorf("error querying user: %v", err)
	}

	// Chuyển đổi từ chuỗi sang time.Time
	user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing created_at: %v", err)
	}

	user.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAtStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing updated_at: %v", err)
	}

	return &user, nil
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
