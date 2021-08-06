package dao

type (
	UserDao struct {
		UserId      int64
		UserGroupId int64
		UserName    string
		LastLogin   string
		Email       string
		Phone       string
		ImageUrl    string
		CreatedAt   string
		Active      bool
	}
)
