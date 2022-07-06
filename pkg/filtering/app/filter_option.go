package app

import (
	"strings"

	"github.com/pkg/errors"
)

// FilterOption is a function type that accepts a pointer to Filters
type FilterOption func(*Filters)

var errUnsupportedFilter = errors.New("unsupported filter")

// FilterOptionFactory is a factory method for creating FilterOption
func FilterOptionFactory(name, value string) (opt FilterOption, err error) {
	switch strings.ToLower(name) {
	case "title":
		opt = WithTitleEqual(value)
	case "version":
		opt = WithVersionEqual(value)
	case "maintainer_has_name":
		opt = WithMaintainerHasName(value)
	case "maintainer_has_email":
		opt = WithMaintainerHasEmail(value)
	case "company":
		opt = WithCompanyEqual(value)
	case "website":
		opt = WithWebsiteEqual(value)
	case "source":
		opt = WithSourceEqual(value)
	case "license":
		opt = WithLicenseEqual(value)
	case "description_has_text":
		opt = WithDescriptionHasText(value)
	default:
		opt = WithNoop()
		err = errors.Wrapf(errUnsupportedFilter, name)
	}
	return
}

// WithNoop returns a function type that do nothing to the filters
func WithNoop() FilterOption {
	return func(f *Filters) {}
}

// WithTitleEqual adds Title 'equal' to the filters
func WithTitleEqual(v string) FilterOption {
	return func(f *Filters) {
		f.Title = &Title{Equal: v}
	}
}

// WithVersionEqual adds Version 'equal' to the filters
func WithVersionEqual(v string) FilterOption {
	return func(f *Filters) {
		f.Version = &Version{Equal: v}
	}
}

// WithMaintainerHasName adds Maintainer 'has' name to the filters
func WithMaintainerHasName(v string) FilterOption {
	return func(f *Filters) {
		if f.Maintainer == nil {
			f.Maintainer = &Maintainer{}
		}
		f.Maintainer.HasName = v
	}
}

// WithMaintainerHasEmail adds Maintainer 'has' email to the filters
func WithMaintainerHasEmail(v string) FilterOption {
	return func(f *Filters) {
		if f.Maintainer == nil {
			f.Maintainer = &Maintainer{}
		}
		f.Maintainer.HasEmail = v
	}
}

// WithCompanyEqual adds Company 'equal' to the filters
func WithCompanyEqual(v string) FilterOption {
	return func(f *Filters) {
		f.Company = &Company{Equal: v}
	}
}

// WithWebsiteEqual adds Website 'equal' to the filters
func WithWebsiteEqual(v string) FilterOption {
	return func(f *Filters) {
		f.Website = &Website{Equal: v}
	}
}

// WithSourceEqual adds Source 'equal' to the filters
func WithSourceEqual(v string) FilterOption {
	return func(f *Filters) {
		f.Source = &Source{Equal: v}
	}
}

// WithLicenseEqual adds License 'equal' to the filters
func WithLicenseEqual(v string) FilterOption {
	return func(f *Filters) {
		f.License = &License{Equal: v}
	}
}

// WithDescriptionHasText adds Description 'has' text to the filters
func WithDescriptionHasText(v string) FilterOption {
	return func(f *Filters) {
		f.Description = &Description{HasText: v}
	}
}
