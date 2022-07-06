package app

import (
	"strings"

	"github.com/nspforever/app-metadata-service/pkg/models"
)

// Filters represents the filters
type Filters struct {
	Title       *Title
	Version     *Version
	Maintainer  *Maintainer
	Company     *Company
	Website     *Website
	Source      *Source
	License     *License
	Description *Description
}

// NewFilters is a constructor function for *Filters
func NewFilters(opts ...FilterOption) *Filters {
	filters := &Filters{}
	for _, opt := range opts {
		opt(filters)
	}
	return filters
}

// Title filter
type Title struct {
	Equal string
}

// Apply returns true if app metadata meets the criteria
func (t *Title) Apply(app models.AppMetadata) bool {
	if t == nil {
		return true
	}
	if t.Equal != "" && t.Equal != app.Title {
		return false
	}

	return true
}

// Version filter
type Version struct {
	Equal string
}

// Apply returns true if app metadata meets the criteria
func (v *Version) Apply(app models.AppMetadata) bool {
	if v == nil {
		return true
	}

	if v.Equal != "" && v.Equal != app.Version {
		return false
	}

	return true
}

// Maintainer filter
type Maintainer struct {
	HasName  string
	HasEmail string
}

// Apply returns true if app metadata meets the criteria
func (m *Maintainer) Apply(app models.AppMetadata) bool {
	if m == nil {
		return true
	}

	for _, am := range app.Maintainer {
		if am.Name == m.HasName || am.Email == m.HasEmail {
			return true
		}
	}

	return false
}

// Company filter
type Company struct {
	Equal string
}

// Apply returns true if app metadata meets the criteria
func (t *Company) Apply(app models.AppMetadata) bool {
	if t == nil {
		return true
	}
	if t.Equal != "" && t.Equal != app.Company {
		return false
	}

	return true
}

// Website filter
type Website struct {
	Equal string
}

// Apply returns true if app metadata meets the criteria
func (t *Website) Apply(app models.AppMetadata) bool {
	if t == nil {
		return true
	}
	if t.Equal != "" && t.Equal != app.Website {
		return false
	}

	return true
}

// Source filter
type Source struct {
	Equal string
}

// Apply returns true if app metadata meets the criteria
func (t *Source) Apply(app models.AppMetadata) bool {
	if t == nil {
		return true
	}
	if t.Equal != "" && t.Equal != app.Source {
		return false
	}

	return true
}

// License filter
type License struct {
	Equal string
}

// Apply returns true if app metadata meets the criteria
func (t *License) Apply(app models.AppMetadata) bool {
	if t == nil {
		return true
	}
	if t.Equal != "" && t.Equal != app.License {
		return false
	}

	return true
}

// Description filter
type Description struct {
	HasText string
}

// Apply returns true if app metadata meets the criteria
func (d *Description) Apply(app models.AppMetadata) bool {
	if d == nil {
		return true
	}
	return strings.Contains(app.Description, d.HasText)
}
