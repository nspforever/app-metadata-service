package rest

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	//mocks_searching "github.com/nspforever/app-metadata-service/pkg/mocks/searching"
	mocks_upserting "github.com/nspforever/app-metadata-service/pkg/mocks/upserting"
	"github.com/nspforever/app-metadata-service/pkg/models"
)

var (
	app1Yaml = `
title: Valid App1
version: 0.0.1
maintainer:
- name: firstmaintainer app1
  email: firstmaintainer@hotmail.com
company: Random Inc
website: https://website.com
source: https://github.com/random/repo
license: Apache-2.0
description: |
  ### Interesting Title,
  Some application content, and description`
	app1JSON = `
{
	"title": "Valid App1",
	"version": "0.0.1",
	"maintainer": [
		{
			"name": "firstmaintainer app1",
			"email": "firstmaintainer@hotmail.com"
		}
	],
	"company": "Random Inc",
	"website": "https://website.com",
	"source": "https://github.com/random/repo",
	"license": "Apache-2.0",
	"description": "### Interesting Title,\nSome application content, and description"
}`
)

func TestUpsertApp(t *testing.T) {
	Convey("Test upsert app", t, func() {
		// Test Setup
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		upserter := mocks_upserting.NewMockService(ctrl)
		handler := NewHandler("localhost:9999", upserter, nil)
		testingServer := httptest.NewServer(handler.router)
		url := testingServer.URL + "/apps"
		defer testingServer.Close()
		client := &http.Client{}

		var app1 models.AppMetadata
		yaml.Unmarshal([]byte(app1Yaml), &app1)
		err := errors.New("upserter error")

		Convey("Test handler logic", func() {
			type test struct {
				testName    string
				contentType string
				payload     string
				app         models.AppMetadata
				err         error
				errCond     string
				statusCode  int
			}
			tests := []test{
				{
					testName:    "Given YAML payload",
					contentType: "application/x-yaml",
					payload:     app1Yaml,
					app:         app1,
					err:         nil,
					errCond:     "no",
					statusCode:  http.StatusOK,
				},
				{
					testName:    "Given JSON payload",
					contentType: "application/json",
					payload:     app1JSON,
					app:         app1,
					err:         nil,
					errCond:     "no",
					statusCode:  http.StatusOK,
				},
				{
					testName:    "Given another YAML payload",
					contentType: "application/x-yaml",
					payload:     app1Yaml,
					app:         app1,
					err:         err,
					errCond:     "an",
					statusCode:  http.StatusOK,
				},
				{
					testName:    "Given another JSON payload",
					contentType: "application/json",
					payload:     app1JSON,
					app:         app1,
					err:         err,
					errCond:     "an",
					statusCode:  http.StatusOK,
				},
			}

			for _, tc := range tests {
				Convey(tc.testName, func() {
					req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(tc.payload)))
					So(err, ShouldBeNil)
					req.Header.Set("Content-type", tc.contentType)
					Convey(fmt.Sprintf("When upsert app metadata %s error", tc.errCond), func() {
						upserter.EXPECT().UpsertApp(&tc.app).Return(tc.err).Times(1)
						Convey(fmt.Sprintf("The response should be %d", tc.statusCode), func() {
							res, err := client.Do(req)
							So(err, ShouldBeNil)
							So(res.StatusCode, ShouldEqual, tc.statusCode)
						})
					})
				})
			}
		})

		Convey("Test validation against YAML payload", func() {
			_, curPath, _, _ := runtime.Caller(0)
			testDataPath := filepath.Join(filepath.Dir(curPath), "../../../", "test-data/invalid_apps.yaml")
			yamlFile, err := ioutil.ReadFile(testDataPath)

			So(err, ShouldBeNil)
			dec := yaml.NewDecoder(bytes.NewBuffer(yamlFile))
			var invalidAppNodes []yaml.Node
			for {
				var node yaml.Node
				err = dec.Decode(&node)
				if errors.Is(err, io.EOF) {
					break
				}
				So(err, ShouldBeNil)
				invalidAppNodes = append(invalidAppNodes, node)
			}
			errMsgs := []string{
				"validation for 'Title' failed on the 'required' tag",
				"validation for 'Version' failed on the 'required' tag",
				"validation for 'Maintainer' failed on the 'required' tag",
				"validation for 'Company' failed on the 'required' tag",
				"validation for 'Website' failed on the 'required' tag",
				"validation for 'Source' failed on the 'required' tag",
				"validation for 'License' failed on the 'required' tag",
				"validation for 'Description' failed on the 'required' tag",
				"validation for 'Email' failed on the 'email' tag",
				"validation for 'Website' failed on the 'url' tag",
				"validation for 'Name' failed on the 'required' tag",
				"validation for 'Email' failed on the 'required' tag",
			}

			for i, appNode := range invalidAppNodes {
				appYaml, err := yaml.Marshal(&appNode)
				So(err, ShouldBeNil)
				req, err := http.NewRequest("PUT", url, bytes.NewBuffer(appYaml))
				So(err, ShouldBeNil)
				req.Header.Set("Content-type", "application/x-yaml")
				res, err := client.Do(req)
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
				bodyBytes, err := ioutil.ReadAll(res.Body)
				defer res.Body.Close()
				So(err, ShouldBeNil)
				So(string(bodyBytes), ShouldContainSubstring, errMsgs[i])
			}
		})
	})
}
