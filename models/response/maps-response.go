package response

type (
	MapsResponse struct {
		Results []addressComponents `json:"results,omitempty"`
	}

	addressComponents struct {
		Components []components `json:"address_components,omitempty"`
		Geometry   location     `json:"geometry,omitempty"`
	}

	components struct {
		LongName  string   `json:"long_name,omitempty"`
		ShortName string   `json:"short_name,omitempty"`
		Types     []string `json:"types,omitempty"`
	}

	location struct {
		Lat string `json:"lat,omitempty"`
		Lng string `json:"lng,omitempty"`
	}
)
