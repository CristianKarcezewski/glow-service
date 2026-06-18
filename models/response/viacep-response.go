package response

type (
	Viacep struct {
		Cep      string `json:"cep,omitempty"`
		Ibge     string `json:"ibge,omitempty"`
		Uf       string `json:"uf,omitempty"`
		City     string `json:"localidade,omitempty"`
		District string `json:"bairro,omitempty"`
		Street   string `json:"logradouro,omitempty"`
	}
)
