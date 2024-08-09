package dto

import (
	"git_truongvudinh/go_web/internal/domain/entity"
)

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
