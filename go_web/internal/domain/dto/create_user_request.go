package dto

import (
	"git_truongvudinh/go_web/internal/domain/entity"
	"time"
)

type UserCreatable interface {
	ToUserEntity() *entity.User
}

// CreateUserRequest
/*
các quy tắc validation trong gin framework :
- required : không được để trống
- email : phải đúng format email
- url : phải là URL hợp lệ
- numeric : phải là giá trị số
- alphanum : chỉ chứa chữ cái và số
- min=X, max=X, oneof=X Y Z
*/

type CreateUserRequest struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,gte=8,lte=32"`
}

func (req *CreateUserRequest) ToUserEntity() *entity.User {
	return &entity.User{
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
	}
}

type CreateUserProjectRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Projects  []struct {
		Name             string     `json:"name" binding:"required"`
		ProjectStartedAt *time.Time `json:"project_started_at"`
		ProjectEndedAt   *time.Time `json:"project_ended_at"`
	} `json:"projects" binding:"required"`
}

func (req *CreateUserProjectRequest) ToUserEntity() *entity.User {
	user := &entity.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	for _, projectReq := range req.Projects {
		startTime := time.Now()
		if projectReq.ProjectStartedAt != nil {
			startTime = *projectReq.ProjectStartedAt
		}

		var endTime *time.Time
		if projectReq.ProjectEndedAt != nil {
			endTime = projectReq.ProjectEndedAt
		}

		project := entity.Project{
			Name:             projectReq.Name,
			ProjectStartedAt: startTime,
			ProjectEndedAt:   endTime, // Nếu ProjectEndedAt là nil, endTime sẽ là nil
		}
		user.Projects = append(user.Projects, project)
	}

	return user
}
