package dtos

type HealthResponseDto struct {
	App         string `json:"app"`
	Environment string `json:"environment"`
	Database    string `json:"database"`
	Redis       string `json:"redis"`
}
