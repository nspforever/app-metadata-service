package searching

import (
	"github.com/nspforever/app-metadata-service/pkg/filtering/app"
	"github.com/nspforever/app-metadata-service/pkg/models"
)

// Service provides app listing operations.
type Service interface {
	SearchApps(filters *app.Filters) ([]models.AppMetadata, error)
}

// Repository provides access to app metadata repository.
type Repository interface {
	// GetApps retrieve apps metadata from the repository.
	GetApps(filters *app.Filters) ([]models.AppMetadata, error)
}

type service struct {
	repo Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetApps retrieves apps metadata
func (s *service) SearchApps(filters *app.Filters) (apps []models.AppMetadata, err error) {
	return s.repo.GetApps(filters)
}
