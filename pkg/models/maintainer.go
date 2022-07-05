package models

// Maintainer represents app maintainer
type Maintainer struct {
	Name  string `json:"name"  yaml:"name" binding:"required"`
	Email string `json:"email"  yaml:"email" binding:"required,email"`
}
