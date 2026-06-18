package models

type (
	Package struct {
		PackageId   int64  `json:"packageId,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		Icon        string `json:"icon,omitemty"`
		Days        int64  `json:"days,omitempty"`
		Value       string `json:"value,omitempty"`
	}
)
