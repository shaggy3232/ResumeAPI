package models

type Project struct {
	Name     string   `json:"name"`
	Duration string   `json:"duration"`
	Stack    []string `json:"stack"`
	Purpose  string   `json:"purpose"`
	RepoLink string   `json:"github-link"`
}
