package dao

import "glow-service/models"

type (
	State struct {
		tableName struct{} `json:"-" pg:"states"`
		StateId   int64    `json:"stateId,omitempty" pg:"id,pk"`
		Uf        string   `json:"uf,omitempty" pg:"uf"`
		Name      string   `json:"name,omitempty" pg:"name"`
	}
)

func (state *State) ToModel() *models.State {
	return &models.State{
		StateId: state.StateId,
		Uf:      state.Uf,
		Name:    state.Name,
	}
}
