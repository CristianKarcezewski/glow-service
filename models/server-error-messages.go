package models

type (
	ServerErrorMessages struct {
		Header headerErrorMessages `json:"header,omitempty"`
	}

	headerErrorMessages struct {
		PlatformNotFound string `json:"platformNotFound,omitempty"`
		NotAuthorized    string `json:"notAuthorized,omitempty"`
	}
)
