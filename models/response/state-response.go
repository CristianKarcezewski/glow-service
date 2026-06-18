package response

import "glow-service/models"

type (
	State struct {
		Id   int64  `json:"id,omitempty"`
		Uf   string `json:"sigla,omitempty"`
		Name string `json:"nome,omitempty"`
	}
)

func (state *State) ToModel() *models.State {
	return &models.State{
		StateId: state.Id,
		Uf:      state.Uf,
		Name:    state.Name,
	}
}
