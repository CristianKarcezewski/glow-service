package dto

type (
	AuthData struct {
		Email    string `json:"email,omitempty" validate:"required,email"`
		Password string `json:"password,omitempty" validate:"required,min=6"`
	}
)
