package models

type AppMetadata struct {
	Title       string       `json:"title"`
	Version     string       `json:"version"`
	Maintainer  []Maintainer `json:"maintainer"`
	Company     string       `json:"company"`
	Website     string       `json:"website"`
	Source      string       `json:"source"`
	License     string       `json:"license"`
	Description string       `json:"description"`
	hash        string
}
