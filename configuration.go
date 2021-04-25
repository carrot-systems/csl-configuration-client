package csl_configuration_client

type Status struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ConfigEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ConfigList []*ConfigEntry

type ConfigResponse struct {
	Status Status `json:"status"`

	Values ConfigList `json:"values"`
}
