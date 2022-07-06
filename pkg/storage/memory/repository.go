package memory

import (
	"errors"
	"sync"

	"github.com/nspforever/app-metadata-service/pkg/filtering/app"
	"github.com/nspforever/app-metadata-service/pkg/models"
	"github.com/nspforever/app-metadata-service/pkg/storage/memory/query"
)

// ErrAppNotFound is returned by GetApps when no app metadata is found
var ErrAppNotFound = errors.New("No app found")

type appFiltersApplier interface {
	Apply(filters *app.Filters, app models.AppMetadata) bool
}

// Storage keeps data in memory
type Storage struct {
	sync.RWMutex
	apps              map[string]models.AppMetadata
	appFiltersApplier appFiltersApplier
}

// New creates new storage instance
func New() *Storage {
	return &Storage{
		apps:              make(map[string]models.AppMetadata),
		appFiltersApplier: &query.SimpleAppFiltersApplier{},
	}
}

// UpsertApp persists the given app metadata to storage
func (s *Storage) UpsertApp(app *models.AppMetadata) (err error) {
	s.Lock()
	defer s.Unlock()
	s.apps[app.Title+"@"+app.Version] = *app
	return
}

// GetApps retrieves apps metadata
func (s *Storage) GetApps(filters *app.Filters) (apps []models.AppMetadata, err error) {
	s.RLock()
	defer s.RUnlock()

	for _, app := range s.apps {
		if !s.appFiltersApplier.Apply(filters, app) {
			continue
		}
		apps = append(apps, app)
	}

	if len(apps) == 0 {
		err = ErrAppNotFound
	}

	return
}
