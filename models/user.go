package models

type (
	User struct {
		UserId      int64  `json:"userId,omitempty"`
		UserGroupId int64  `json:"userGroupId,omitempty"`
		UserName    string `json:"name,omitempty"`
		LastLogin   string `json:"-"`
		Email       string `json:"email,omitempty"`
		Phone       string `json:"phone,omitempty"`
		ImageUrl    string `json:"imageUrl,omitempty"`
		CreatedAt   string `json:"-"`
		Password    string `json:"-"`
		Active      bool   `json:"-"`
	}

	Auth struct {
		Authorization string `json:"authorization,omitempty"`
	}
)
