package memory

import (
	"errors"
	"sync"

	"github.com/nspforever/app-metadata-service/pkg/filtering/app"
	"github.com/nspforever/app-metadata-service/pkg/models"
)

var errAppNotFound = errors.New("No app found")

// Memory storage keeps data in memory
type Storage struct {
	sync.RWMutex
	apps map[string]models.AppMetadata
}

func New() *Storage {
	return &Storage{apps: make(map[string]models.AppMetadata)}
}

// UpsertApp persists the given app metadata to storage
func (s *Storage) UpsertApp(app *models.AppMetadata) (err error) {
	s.Lock()
	defer s.Unlock()
	s.apps[app.Title+"@"+app.Version] = *app
	return err
}

// GetApps retrieves apps metadata
func (s *Storage) GetApps(filters *app.Filters) (apps []models.AppMetadata, err error) {
	s.RLock()
	defer s.RUnlock()

	for _, app := range s.apps {
		if !filters.Apply(app) {
			continue
		}
		apps = append(apps, app)
	}
	if len(apps) == 0 {
		err = errAppNotFound
	}
	return apps, err
}
