package listing

import (
	"github.com/nspforever/app-metadata-service/pkg/models"
)

// Service provides app listing operations.
type Service interface {
	GetApps() ([]models.AppMetadata, error)
}

// Repository provides access to app metadata repository.
type Repository interface {
	// GetApps retrieve apps metadata from the repository.
	GetApps() ([]models.AppMetadata, error)
}

type service struct {
	repo Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetApps retrieves apps metadata
func (s *service) GetApps() (apps []models.AppMetadata, err error) {
	// TODO add validation
	return s.repo.GetApps()
}
