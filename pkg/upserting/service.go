package upserting

import (
	"github.com/nspforever/app-metadata-service/pkg/models"
)

// Service provides app upsert operations.
type Service interface {
	UpsertApp(*models.AppMetadata) error
}

// Repository provides access to app metadata repository.
type Repository interface {
	// UpsertApp add a given app metadata to the repository.
	UpsertApp(*models.AppMetadata) error
}

type service struct {
	repo Repository
}

// NewService creates an upserting service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// UpsertApp persists the given app metadata to storage
func (s *service) UpsertApp(app *models.AppMetadata) error {
	// TODO add validation
	return s.repo.UpsertApp(app)
}
