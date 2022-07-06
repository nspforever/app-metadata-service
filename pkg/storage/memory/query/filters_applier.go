package query

import (
	"github.com/nspforever/app-metadata-service/pkg/filtering/app"
	"github.com/nspforever/app-metadata-service/pkg/models"
)

// SimpleAppFiltersApplier implements a simple way to apply app filters
type SimpleAppFiltersApplier struct {
}

// Apply checks if the app metadata meets the criteria by applying the filters
func (fa SimpleAppFiltersApplier) Apply(f *app.Filters, app models.AppMetadata) bool {
	if f == nil {
		return true
	}

	if !f.Title.Apply(app) {
		return false
	}

	if !f.Version.Apply(app) {
		return false
	}

	if !f.Maintainer.Apply(app) {
		return false
	}

	if !f.Company.Apply(app) {
		return false
	}

	if !f.Website.Apply(app) {
		return false
	}

	if !f.Source.Apply(app) {
		return false
	}

	if !f.License.Apply(app) {
		return false
	}

	if !f.Description.Apply(app) {
		return false
	}

	return true
}
