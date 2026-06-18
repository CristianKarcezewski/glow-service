package models

type (
	State struct {
		StateId int64  `json:"stateId,omitempty"`
		Uf      string `json:"uf,omitempty"`
		Name    string `json:"name,omitempty"`
	}
)
