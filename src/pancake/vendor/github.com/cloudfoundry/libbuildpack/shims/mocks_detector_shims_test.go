// Code generated by MockGen. DO NOT EDIT.
// Source: detector.go

// Package shims_test is a generated GoMock package.
package shims_test

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInstaller is a mock of Installer interface
type MockInstaller struct {
	ctrl     *gomock.Controller
	recorder *MockInstallerMockRecorder
}

// MockInstallerMockRecorder is the mock recorder for MockInstaller
type MockInstallerMockRecorder struct {
	mock *MockInstaller
}

// NewMockInstaller creates a new mock instance
func NewMockInstaller(ctrl *gomock.Controller) *MockInstaller {
	mock := &MockInstaller{ctrl: ctrl}
	mock.recorder = &MockInstallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInstaller) EXPECT() *MockInstallerMockRecorder {
	return m.recorder
}

// InstallCNBS mocks base method
func (m *MockInstaller) InstallCNBS(orderFile, installDir string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallCNBS", orderFile, installDir)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallCNBS indicates an expected call of InstallCNBS
func (mr *MockInstallerMockRecorder) InstallCNBS(orderFile, installDir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallCNBS", reflect.TypeOf((*MockInstaller)(nil).InstallCNBS), orderFile, installDir)
}

// InstallLifecycle mocks base method
func (m *MockInstaller) InstallLifecycle(dst string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallLifecycle", dst)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallLifecycle indicates an expected call of InstallLifecycle
func (mr *MockInstallerMockRecorder) InstallLifecycle(dst interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallLifecycle", reflect.TypeOf((*MockInstaller)(nil).InstallLifecycle), dst)
}
