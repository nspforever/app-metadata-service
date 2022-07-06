package memory

import (
	"strconv"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nspforever/app-metadata-service/pkg/filtering/app"
	mock_memory "github.com/nspforever/app-metadata-service/pkg/mocks/storage/memory"
	"github.com/nspforever/app-metadata-service/pkg/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUpsertApp(t *testing.T) {
	Convey("Test upsert app", t, func() {
		Convey("Given an app metadata", func() {
			// initialize a storage with one existing record
			s := &Storage{
				apps: map[string]models.AppMetadata{
					"app 1@1.0.1": {
						Title:   "app 1",
						Version: "1.0.1",
						Company: "Upbound",
					},
				},
			}

			Convey("When it is a new app", func() {
				newApp := &models.AppMetadata{
					Title:   "app 2",
					Version: "1.0.1",
				}
				s.UpsertApp(newApp)
				Convey("The new app should be persisted into storage", func() {
					So(len(s.apps), ShouldEqual, 2)
					So(s.apps[newApp.Title+"@"+newApp.Version], ShouldResemble, *newApp)
				})
			})
			Convey("When it is an existing app", func() {
				existingApp := &models.AppMetadata{
					Title:   "app 1",
					Version: "1.0.1",
					Company: "Upbound Inc.",
				}
				s.UpsertApp(existingApp)
				Convey("The new app should be persisted into storage", func() {
					So(len(s.apps), ShouldEqual, 1)
					So(s.apps[existingApp.Title+"@"+existingApp.Version], ShouldResemble, *existingApp)
				})
			})
		})
	})
}

func TestGetApps(t *testing.T) {
	Convey("Test upsert app", t, func() {
		Convey("Given an app metadata", func() {
			// setup mock
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			filtersApplier := mock_memory.NewMockappFiltersApplier(ctrl)

			// initialize a storage with one existing record
			app1v1 := models.AppMetadata{
				Title:   "app1",
				Version: "1.0.1",
				Company: "Upbound",
			}
			app1v2 := models.AppMetadata{
				Title:   "app1",
				Version: "1.0.2",
				Company: "Upbound",
			}
			app2v1 := models.AppMetadata{
				Title:   "app2",
				Version: "1.0.1",
				Company: "Upbound",
			}

			s := &Storage{
				apps: map[string]models.AppMetadata{
					"app1@1.0.1": app1v1,
					"app1@1.0.2": app1v2,
					"app2@1.0.1": app2v1,
				},
				appFiltersApplier: filtersApplier,
			}

			Convey("When no app found by the filters", func() {
				filters := &app.Filters{}
				filtersApplier.EXPECT().Apply(filters, gomock.Any()).Return(false).Times(3)
				_, err := s.GetApps(filters)
				Convey("The app not found error should be returned", func() {
					So(err, ShouldBeError, ErrAppNotFound)
				})
			})
			Convey("When some apps meet the criteria", func() {
				filters := &app.Filters{}
				filtersApplier.EXPECT().Apply(filters, app1v1).Return(true).Times(1)
				filtersApplier.EXPECT().Apply(filters, app1v2).Return(false).Times(1)
				filtersApplier.EXPECT().Apply(filters, app2v1).Return(true).Times(1)
				appsFound, err := s.GetApps(filters)
				Convey("The app not found error should be returned", func() {
					So(err, ShouldBeNil)
					So(appsFound, ShouldContain, app1v1)
					So(appsFound, ShouldContain, app2v1)
				})
			})
		})
	})
}

func TestConcurrent(t *testing.T) {
	Convey("Test upsert apps concurrently", t, func() {
		appsWanted := 1000
		s := New()
		var wg sync.WaitGroup
		wg.Add(appsWanted)

		for i := 0; i < appsWanted; i++ {
			go func(i int) {
				newApp := &models.AppMetadata{
					Title:   "app " + strconv.FormatInt(int64(i), 10),
					Version: "1.0.1",
					Company: "Upbound Inc.",
				}
				s.UpsertApp(newApp)
				s.GetApps(&app.Filters{})
				wg.Done()
			}(i)
		}
		wg.Wait()
		So(len(s.apps), ShouldEqual, appsWanted)
	})
}
