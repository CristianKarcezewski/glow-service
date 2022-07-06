package dao

import (
	"glow-service/models"
)

type (
	UserDao struct {
		tableName   struct{} `json:"-" pg:"users"`
		UserId      int64    `json:"userId,omitempty" pg:"id,pk"`
		Uid         string   `json:"-" pg:"uid"`
		UserGroupId int64    `json:"userGroupId,omitempty" pg:"user_group_id"`
		UserName    string   `json:"name,omitempty" pg:"user_name"`
		LastLogin   string   `json:"-" pg:"last_login"`
		Email       string   `json:"email,omitempty" pg:"unique:email"`
		Phone       string   `json:"phone,omitempty" pg:"phone"`
		ImageUrl    string   `json:"imageUrl,omitempty" pg:"image_url"`
		CreatedAt   string   `json:"-" pg:"created_at"`
		Active      bool     `json:"active,omitempty" pg:"active"`
	}

	Hash struct {
		tableName struct{} `pg:"hashs"`
		HashId    int64    `pg:"id,pk"`
		UserId    int64    `pg:"user_id"`
		Hash      string   `pg:"hash"`
	}
)

func NewDAOUser(u *models.User) *UserDao {
	newUser := UserDao{
		UserId:      u.UserId,
		Uid:         u.Uid,
		UserGroupId: u.UserGroupId,
		UserName:    u.UserName,
		LastLogin:   u.LastLogin,
		Email:       u.Email,
		Phone:       u.Phone,
		ImageUrl:    u.ImageUrl,
		CreatedAt:   u.CreatedAt,
		Active:      u.Active,
	}

	return &newUser
}

func (u *UserDao) ToModel() *models.User {
	return &models.User{
		UserId:      u.UserId,
		Uid:         u.Uid,
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
