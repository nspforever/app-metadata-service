package upserting

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	mocks_upserting "github.com/nspforever/app-metadata-service/pkg/mocks/upserting"
	"github.com/nspforever/app-metadata-service/pkg/models"
)

func TestUpsertApp(t *testing.T) {
	Convey("Test upsert app", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repo := mocks_upserting.NewMockRepository(ctrl)
		svc := NewService(repo)
		newApp := &models.AppMetadata{
			Title:   "app 2",
			Version: "1.0.1",
		}
		Convey("When upsert apps metadata succeeded", func() {
			repo.EXPECT().UpsertApp(newApp).Return(nil).Times(1)
			err := svc.UpsertApp(newApp)
			Convey("No error should be returned", func() {
				So(err, ShouldBeNil)
			})
		})
		Convey("When upsert app metadata failed", func() {
			expectErr := errors.New("storage error")
			repo.EXPECT().UpsertApp(newApp).Return(expectErr).Times(1)
			err := svc.UpsertApp(newApp)
			Convey("An error should be returned", func() {
				So(err, ShouldBeError, expectErr)
			})
		})
	})
}
