package app

import (
	"strings"

	"github.com/pkg/errors"
)

type FilterOption func(*Filters)

var errUnsupportedFilter = errors.New("unsupported filter")

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

func WithNoop() FilterOption {
	return func(f *Filters) {}
}

func WithTitleEqual(v string) FilterOption {
	return func(f *Filters) {
		f.Title = &Title{Equal: v}
	}
}

func WithVersionEqual(v string) FilterOption {
	return func(f *Filters) {
		f.Version = &Version{Equal: v}
	}
}

func WithMaintainerHasName(v string) FilterOption {
	return func(f *Filters) {
		if f.Maintainer == nil {
			f.Maintainer = &Maintainer{}
		}
		f.Maintainer.HasName = v
	}
}

func WithMaintainerHasEmail(v string) FilterOption {
	return func(f *Filters) {
		if f.Maintainer == nil {
			f.Maintainer = &Maintainer{}
		}
		f.Maintainer.HasEmail = v
	}
}

func WithCompanyEqual(v string) FilterOption {
	return func(f *Filters) {
		f.Company = &Company{Equal: v}
	}
}

func WithWebsiteEqual(v string) FilterOption {
	return func(f *Filters) {
		f.Website = &Website{Equal: v}
	}
}

func WithSourceEqual(v string) FilterOption {
	return func(f *Filters) {
		f.Source = &Source{Equal: v}
	}
}

func WithLicenseEqual(v string) FilterOption {
	return func(f *Filters) {
		f.License = &License{Equal: v}
	}
}

func WithDescriptionHasText(v string) FilterOption {
	return func(f *Filters) {
		f.Description = &Description{HasText: v}
	}
}
