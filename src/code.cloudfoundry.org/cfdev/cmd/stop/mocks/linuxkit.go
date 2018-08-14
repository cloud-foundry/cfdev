// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cfdev/cmd/stop (interfaces: Hypervisor)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHypervisor is a mock of Hypervisor interface
type MockHypervisor struct {
	ctrl     *gomock.Controller
	recorder *MockHypervisorMockRecorder
}

// MockHypervisorMockRecorder is the mock recorder for MockHypervisor
type MockHypervisorMockRecorder struct {
	mock *MockHypervisor
}

// NewMockHypervisor creates a new mock instance
func NewMockHypervisor(ctrl *gomock.Controller) *MockHypervisor {
	mock := &MockHypervisor{ctrl: ctrl}
	mock.recorder = &MockHypervisorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHypervisor) EXPECT() *MockHypervisorMockRecorder {
	return m.recorder
}

// Destroy mocks base method
func (m *MockHypervisor) Destroy(arg0 string) error {
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
func (mr *MockHypervisorMockRecorder) Destroy(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockHypervisor)(nil).Destroy), arg0)
}

// Stop mocks base method
func (m *MockHypervisor) Stop(arg0 string) error {
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockHypervisorMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockHypervisor)(nil).Stop), arg0)
}