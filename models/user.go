package models

type (
	User struct {
		UserId      int64  `json:"userId,omitempty"`
		Uid         string `json:"-"`
		UserGroupId int64  `json:"userGroupId,omitempty"`
		UserName    string `json:"name,omitempty"`
		LastLogin   string `json:"-"`
		Email       string `json:"email,omitempty"`
		Phone       string `json:"phone,omitempty"`
		FileUrl     string `json:"FileUrl,omitempty"`
		CreatedAt   string `json:"-"`
		Password    string `json:"-"`
		Active      bool   `json:"-"`
		DaysLeft    int64  `json:"daysLeft,omitempty"`
	}

	Auth struct {
		UserId        int64  `json:"userId,omitempty"`
		Uid           string `json:"uid,omitempty"`
		UserGroupId   int64  `json:"userGroupId,omitempty"`
		Authorization string `json:"authorization,omitempty"`

		UserName string `json:"name,omitempty"`
		Email    string `json:"email,omitempty"`
		Phone    string `json:"phone,omitempty"`
		FileUrl  string `json:"FileUrl,omitempty"`
		DaysLeft int64  `json:"daysLeft,omitempty"`
	}
)
