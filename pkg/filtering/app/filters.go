package app

import (
	"strings"

	"github.com/nspforever/app-metadata-service/pkg/models"
)

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

func NewFilters(opts ...FilterOption) *Filters {
	filters := &Filters{}
	for _, opt := range opts {
		opt(filters)
	}
	return filters
}

func (f *Filters) Apply(app models.AppMetadata) bool {
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

type Title struct {
	Equal string
}

func (t *Title) Apply(app models.AppMetadata) bool {
	if t == nil {
		return true
	}
	if t.Equal != "" && t.Equal != app.Title {
		return false
	}

	return true
}

type Version struct {
	Equal string
}

func (v *Version) Apply(app models.AppMetadata) bool {
	if v == nil {
		return true
	}

	if v.Equal != "" && v.Equal != app.Version {
		return false
	}

	return true
}

type Maintainer struct {
	HasName  string
	HasEmail string
}

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

type Company struct {
	Equal string
}

func (t *Company) Apply(app models.AppMetadata) bool {
	if t == nil {
		return true
	}
	if t.Equal != "" && t.Equal != app.Company {
		return false
	}

	return true
}

type Website struct {
	Equal string
}

func (t *Website) Apply(app models.AppMetadata) bool {
	if t == nil {
		return true
	}
	if t.Equal != "" && t.Equal != app.Website {
		return false
	}

	return true
}

type Source struct {
	Equal string
}

func (t *Source) Apply(app models.AppMetadata) bool {
	if t == nil {
		return true
	}
	if t.Equal != "" && t.Equal != app.Source {
		return false
	}

	return true
}

type License struct {
	Equal string
}

func (t *License) Apply(app models.AppMetadata) bool {
	if t == nil {
		return true
	}
	if t.Equal != "" && t.Equal != app.License {
		return false
	}

	return true
}

type Description struct {
	HasText string
}

func (d *Description) Apply(app models.AppMetadata) bool {
	if d == nil {
		return true
	}
	return strings.Contains(app.Description, d.HasText)
}
