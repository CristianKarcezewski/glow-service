package models

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
