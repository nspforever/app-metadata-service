// +build smoke

package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/nspforever/app-metadata-service/pkg/models"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"
)

const baseURL = "http://localhost:9999/api/v1/apps"

var apps = getAllApps()

func TestAppMetadataEndpoints(t *testing.T) {
	Convey("Test app metadata endpoints", t, func() {
		Convey("Test upsert app metadata endpoint", func() {
			type test struct {
				testName    string
				contentType string
				filename    string
			}
			tests := []test{
				{
					testName:    "Upsert app1",
					contentType: "application/x-yaml",
					filename:    "app1.yaml",
				},
				{
					testName:    "Upsert app2",
					contentType: "application/x-yaml",
					filename:    "app2.yaml",
				},
				{
					testName:    "Upsert app3",
					contentType: "application/x-yaml",
					filename:    "app3.yaml",
				},
			}

			for _, tc := range tests {
				Convey(tc.testName, func() {
					payload, err := loadAppMetadataFromYaml(tc.filename)
					So(err, ShouldBeNil)
					req, err := http.NewRequest("PUT", baseURL, bytes.NewBuffer(payload))
					req.Header.Set("Content-type", "application/x-yaml")
					So(err, ShouldBeNil)
					client := &http.Client{}
					res, err := client.Do(req)
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, http.StatusOK)
				})
			}
		})
		Convey("Test search app metadata endpoint", func() {
			type test struct {
				testName    string
				queryString string
				appsFound   []models.AppMetadata
			}
			client := &http.Client{}
			tests := []test{
				{
					testName:    "Search by title",
					queryString: "?title=Valid%20App1",
					appsFound:   []models.AppMetadata{apps[0]},
				},
				{
					testName:    "Search by maintainer name",
					queryString: "?maintainer_has_name=AppTwo%20Maintainer",
					appsFound:   []models.AppMetadata{apps[1]},
				},
				{
					testName:    "Search by company",
					queryString: "?company=Upbound%20Inc",
					appsFound:   []models.AppMetadata{apps[1], apps[2]},
				},

				{
					testName:    "description has text",
					queryString: "?description_has_text=app",
					appsFound:   []models.AppMetadata{apps[0], apps[1], apps[2]},
				},
			}

			for _, tc := range tests {
				Convey(tc.testName, func() {
					req, err := http.NewRequest("GET", baseURL+tc.queryString, nil)
					So(err, ShouldBeNil)
					res, err := client.Do(req)
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, http.StatusOK)
					bodyBytes, err := ioutil.ReadAll(res.Body)
					defer res.Body.Close()
					So(err, ShouldBeNil)
					var resObj models.AppSearchResponse
					json.Unmarshal(bodyBytes, &resObj)
					So(resObj.Count, ShouldEqual, len(tc.appsFound))
					for _, app := range tc.appsFound {
						So(resObj.Data, ShouldContain, app)
					}
				})
			}
		})
	})

}

func getAllApps() (apps []models.AppMetadata) {
	for i := 0; i < 3; i++ {
		appYaml, _ := loadAppMetadataFromYaml(fmt.Sprintf("app%d.yaml", i+1))
		var app models.AppMetadata
		yaml.Unmarshal([]byte(appYaml), &app)
		apps = append(apps, app)
	}

	return
}

func loadAppMetadataFromYaml(filename string) ([]byte, error) {
	_, curPath, _, _ := runtime.Caller(0)
	fullPath := filepath.Join(filepath.Dir(curPath), "../../", "test-data", filename)
	return ioutil.ReadFile(fullPath)
}
