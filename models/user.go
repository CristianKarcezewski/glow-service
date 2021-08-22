package models

import "glow-service/models/dao"

type (
	User struct {
		UserId      int64
		UserGroupId int64
		UserName    string
		LastLogin   string
		Email       string
		Phone       string
		ImageUrl    string
		CreatedAt   string
		Password    string
		Active      bool
	}
)

func (u *User) AdaptToDAO() *dao.User {
	return &dao.User{
		UserId:      u.UserId,
		UserGroupId: u.UserGroupId,
		UserName:    u.UserName,
		LastLogin:   u.LastLogin,
		Email:       u.Email,
		Phone:       u.Phone,
		ImageUrl:    u.ImageUrl,
		CreatedAt:   u.CreatedAt,
		Active:      u.Active,
	}
}
