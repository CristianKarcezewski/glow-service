package dao

import (
	"errors"
	"glow-service/models"

	"golang.org/x/crypto/bcrypt"
)

type (
	User struct {
		tableName   struct{} `pg:"users"`
		UserId      int64    `pg:"id,pk"`
		UserGroupId int64    `pg:"user_group_id"`
		UserName    string   `pg:"user_name"`
		LastLogin   string   `pg:"last_login"`
		Email       string   `pg:"unique:email"`
		Phone       string   `pg:"phone"`
		ImageUrl    string   `pg:"image_url"`
		CreatedAt   string   `pg:"created_at"`
		Active      bool     `pg:"active"`
	}

	Hash struct {
		tableName struct{} `pg:"hashs"`
		HashId    int64    `pg:"id,pk"`
		UserId    int64    `pg:"user_id"`
		Hash      string   `pg:"hash"`
	}
)

func NewDAOUser(u *models.User) (*User, *Hash, error) {
	newUser := User{
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

	bcrypt, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, errors.New("Error encrypting hash")
	}

	hash := Hash{
		Hash: string(bcrypt),
	}

	return &newUser, &hash, nil
}

func (u *User) ToModel() *models.User {
	return &models.User{
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
