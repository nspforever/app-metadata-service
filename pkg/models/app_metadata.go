package models

type AppMetadata struct {
	Title       string       `json:"title" yaml:"title" binding:"required"`
	Version     string       `json:"version" yaml:"version" binding:"required"`
	Maintainer  []Maintainer `json:"maintainer" yaml:"maintainer" binding:"required,dive,required"`
	Company     string       `json:"company" yaml:"company" binding:"required"`
	Website     string       `json:"website" yaml:"website" binding:"required,url"`
	Source      string       `json:"source" yaml:"source" binding:"required,url"`
	License     string       `json:"license" yaml:"license" binding:"required"`
	Description string       `json:"description" yaml:"description" binding:"required"`
}
