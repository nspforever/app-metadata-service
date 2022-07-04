package memory

import (
	"sync"

	"github.com/nspforever/app-metadata-service/pkg/models"
)

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
func (s *Storage) GetApps() (apps []models.AppMetadata, err error) {
	s.RLock()
	defer s.RUnlock()

	for _, app := range s.apps {
		apps = append(apps, app)
	}

	return apps, err
}
