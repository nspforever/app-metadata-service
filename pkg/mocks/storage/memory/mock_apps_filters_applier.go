// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/guannan.du/workspace/go/src/github.com/nspforever/app-metadata-service/pkg/storage/memory/repository.go

// Package memory is a generated GoMock package.
package memory

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	app "github.com/nspforever/app-metadata-service/pkg/filtering/app"
	models "github.com/nspforever/app-metadata-service/pkg/models"
)

// MockappFiltersApplier is a mock of appFiltersApplier interface.
type MockappFiltersApplier struct {
	ctrl     *gomock.Controller
	recorder *MockappFiltersApplierMockRecorder
}

// MockappFiltersApplierMockRecorder is the mock recorder for MockappFiltersApplier.
type MockappFiltersApplierMockRecorder struct {
	mock *MockappFiltersApplier
}

// NewMockappFiltersApplier creates a new mock instance.
func NewMockappFiltersApplier(ctrl *gomock.Controller) *MockappFiltersApplier {
	mock := &MockappFiltersApplier{ctrl: ctrl}
	mock.recorder = &MockappFiltersApplierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockappFiltersApplier) EXPECT() *MockappFiltersApplierMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockappFiltersApplier) Apply(filters *app.Filters, app models.AppMetadata) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", filters, app)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockappFiltersApplierMockRecorder) Apply(filters, app interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockappFiltersApplier)(nil).Apply), filters, app)
}