package searching

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/nspforever/app-metadata-service/pkg/filtering/app"
	mocks_searching "github.com/nspforever/app-metadata-service/pkg/mocks/searching"
	"github.com/nspforever/app-metadata-service/pkg/models"
	"github.com/nspforever/app-metadata-service/pkg/storage/memory"
)

func TestSearchApps(t *testing.T) {
	Convey("Test search apps", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repo := mocks_searching.NewMockRepository(ctrl)
		svc := NewService(repo)
		filters := &app.Filters{}
		apps := []models.AppMetadata{
			{
				Title:   "app1",
				Version: "1.1.1",
			},
			{
				Title:   "app2",
				Version: "2.2.2",
			},
		}
		Convey("When GetApps returns app metadata", func() {
			repo.EXPECT().GetApps(filters).Return(apps, nil).Times(1)
			appsFound, err := svc.SearchApps(filters)
			Convey("Metadata of found apps should be returned without an error", func() {
				So(err, ShouldBeNil)
				So(appsFound, ShouldResemble, apps)
			})
		})
		Convey("When GetApps returns no app metadata", func() {
			repo.EXPECT().GetApps(filters).Return(nil, memory.ErrAppNotFound).Times(1)

			_, err := svc.SearchApps(filters)
			Convey("ErrAppNotFound should be returned", func() {
				So(err, ShouldBeError, memory.ErrAppNotFound)
			})
		})
		Convey("When GetApps returns an error", func() {
			expectErr := errors.New("storage error")
			repo.EXPECT().GetApps(filters).Return(nil, expectErr).Times(1)

			_, err := svc.SearchApps(filters)
			Convey("An error should be returned", func() {
				So(err, ShouldBeError, expectErr)
			})
		})
	})
}
