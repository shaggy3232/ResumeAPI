package models

type Work struct {
	Company      string   `json:"company"`
	Duration     string   `json:"duration"`
	Tasks        []string `json:"task"`
	Title        string   `json:"title"`
	technologies []string `json:"technologies"`
}
