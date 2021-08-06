package dto

import "glow-service/models"

type (
	UserDto struct {
		UserId   int64  `json:"userId,omitempty"`
		UserName string `json:"userName,omitempty" validate:"required"`
		Email    string `json:"email,omitempty" validate:"required,email"`
		Phone    string `json:"phone,omitempty"`
		Password string `json:"password,omitempty" validate:"required,min=6"`
	}
)

func (u *UserDto) ToModel() *models.User {
	return &models.User{
		UserId:   u.UserId,
		UserName: u.UserName,
		Email:    u.Email,
		Phone:    u.Phone,
		Password: u.Password,
	}
}
