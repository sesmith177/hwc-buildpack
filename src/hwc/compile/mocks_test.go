// Code generated by MockGen. DO NOT EDIT.
// Source: compile.go

package compile_test

import (
	libbuildpack "github.com/cloudfoundry/libbuildpack"
	gomock "github.com/golang/mock/gomock"
)

// MockManifest is a mock of Manifest interface
type MockManifest struct {
	ctrl     *gomock.Controller
	recorder *MockManifestMockRecorder
}

// MockManifestMockRecorder is the mock recorder for MockManifest
type MockManifestMockRecorder struct {
	mock *MockManifest
}

// NewMockManifest creates a new mock instance
func NewMockManifest(ctrl *gomock.Controller) *MockManifest {
	mock := &MockManifest{ctrl: ctrl}
	mock.recorder = &MockManifestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockManifest) EXPECT() *MockManifestMockRecorder {
	return _m.recorder
}

// DefaultVersion mocks base method
func (_m *MockManifest) DefaultVersion(_param0 string) (libbuildpack.Dependency, error) {
	ret := _m.ctrl.Call(_m, "DefaultVersion", _param0)
	ret0, _ := ret[0].(libbuildpack.Dependency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DefaultVersion indicates an expected call of DefaultVersion
func (_mr *MockManifestMockRecorder) DefaultVersion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DefaultVersion", arg0)
}

// InstallDependency mocks base method
func (_m *MockManifest) InstallDependency(_param0 libbuildpack.Dependency, _param1 string) error {
	ret := _m.ctrl.Call(_m, "InstallDependency", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallDependency indicates an expected call of InstallDependency
func (_mr *MockManifestMockRecorder) InstallDependency(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InstallDependency", arg0, arg1)
}