package repositories

import "gorm.io/gorm"

type IProjectRepository interface {
}

type ProjectRepository struct {
	base
}

func NewProjectRepository(db *gorm.DB) IProjectRepository {
	return &ProjectRepository{base: base{db: db}}
}
