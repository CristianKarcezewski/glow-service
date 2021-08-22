package dao

type (
	User struct {
		tableName   struct{} `sql:"users"`
		UserId      int64    `json:"userId,omitempty" pg:"id,pk"`
		UserGroupId int64    `json:"userGroupId,omitempty" pg:"user_group_id"`
		UserName    string   `json:"userName,omitempty" pg:"user_name"`
		LastLogin   string   `json:"lastLogin,omitempty" pg:"last_login"`
		Email       string   `json:"email,omitempty" pg:"email"`
		Phone       string   `json:"phone,omitempty" pg:"phone"`
		ImageUrl    string   `json:"imageUrl,omitempty" pg:"image_url"`
		CreatedAt   string   `json:"createdAt,omitempty" pg:"created_at"`
		Active      bool     `json:"active" pg:"active"`
	}
)
