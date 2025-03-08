package models

type Resume struct {
	Contact  Contact   `json:"contact"`
	Projects []Project `json:"projects`
	Work     []Work    `json:"work"`
}
