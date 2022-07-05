package memory

import (
	"strconv"
	"sync"
	"testing"

	"github.com/nspforever/app-metadata-service/pkg/filtering/app"
	"github.com/nspforever/app-metadata-service/pkg/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUpsertApp(t *testing.T) {
	Convey("Test upsert app", t, func() {
		Convey("Given an app metadata", func() {
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

func TestConcurrentUpsertApps(t *testing.T) {
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
