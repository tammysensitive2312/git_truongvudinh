package repositories

import (
	"context"
	"git_truongvudinh/go_web/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Tạo bảng cho các entity cần thiết
	err = db.AutoMigrate(&entities.User{}, &entities.Project{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestCreateUser(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("failed to setup test db: %v", err)
	}

	repo := NewUserRepository(db)

	// Tạo user để test
	user := &entities.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "hashedpassword",
	}

	createdUser, err := repo.CreateUser(context.Background(), user)
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, "John", createdUser.FirstName)
	assert.Equal(t, "Doe", createdUser.LastName)
	assert.Equal(t, "john.doe@example.com", createdUser.Email)
}

func TestGetUserById(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("failed to setup test db: %v", err)
	}

	repo := NewUserRepository(db)

	// Tạo user để test
	user := &entities.User{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane.doe@example.com",
		Password:  "hashedpassword",
	}
	createdUser, err := repo.CreateUser(context.Background(), user)
	assert.NoError(t, err)

	// Lấy user không preload
	fetchedUser, err := repo.GetUserById(context.Background(), createdUser.ID, false)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedUser)
	assert.Equal(t, "Jane", fetchedUser.FirstName)
	assert.Empty(t, fetchedUser.Projects)

	// Tạo project và liên kết với user
	project := &entities.Project{
		Name:             "Project 1",
		ProjectStartedAt: createdUser.CreatedAt,
		UserID:           int64(createdUser.ID),
	}
	db.Create(project)

	// Lấy user với preload
	fetchedUserWithProjects, err := repo.GetUserById(context.Background(), createdUser.ID, true)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedUserWithProjects)
	assert.NotEmpty(t, fetchedUserWithProjects.Projects)
	assert.Equal(t, "Project 1", fetchedUserWithProjects.Projects[0].Name)
}
