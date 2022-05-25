package response

type (
	MapsResponse struct {
		Results []addressComponents `json:"results,omitempty"`
	}

	addressComponents struct {
		Components []components `json:"address_components,omitempty"`
		Geometry   geonetry     `json:"geometry,omitempty"`
	}

	geonetry struct {
		Location location `json:"location,omitempty"`
	}

	components struct {
		LongName  string   `json:"long_name,omitempty"`
		ShortName string   `json:"short_name,omitempty"`
		Types     []string `json:"types,omitempty"`
	}

	location struct {
		Lat float64 `json:"lat,omitempty"`
		Lng float64 `json:"lng,omitempty"`
	}
)
