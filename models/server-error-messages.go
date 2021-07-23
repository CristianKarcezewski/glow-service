package models

type (
	ServerErrorMessages struct {
		Header headerErrorMessages `json:"header,omitempty"`
	}

	headerErrorMessages struct {
		PlatformNotFound      string `json:"platformNotFound,omitempty"`
		AuthorizationNotFound string `json:"authorizationNotFound,omitempty"`
	}
)
