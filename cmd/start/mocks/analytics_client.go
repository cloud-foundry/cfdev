// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cfdev/cmd/start (interfaces: AnalyticsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAnalyticsClient is a mock of AnalyticsClient interface
type MockAnalyticsClient struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyticsClientMockRecorder
}

// MockAnalyticsClientMockRecorder is the mock recorder for MockAnalyticsClient
type MockAnalyticsClientMockRecorder struct {
	mock *MockAnalyticsClient
}

// NewMockAnalyticsClient creates a new mock instance
func NewMockAnalyticsClient(ctrl *gomock.Controller) *MockAnalyticsClient {
	mock := &MockAnalyticsClient{ctrl: ctrl}
	mock.recorder = &MockAnalyticsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAnalyticsClient) EXPECT() *MockAnalyticsClientMockRecorder {
	return m.recorder
}

// Event mocks base method
func (m *MockAnalyticsClient) Event(arg0 string, arg1 ...map[string]interface{}) error {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Event", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Event indicates an expected call of Event
func (mr *MockAnalyticsClientMockRecorder) Event(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Event", reflect.TypeOf((*MockAnalyticsClient)(nil).Event), varargs...)
}

// PromptOptInIfNeeded mocks base method
func (m *MockAnalyticsClient) PromptOptInIfNeeded(arg0 string) error {
	ret := m.ctrl.Call(m, "PromptOptInIfNeeded", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PromptOptInIfNeeded indicates an expected call of PromptOptInIfNeeded
func (mr *MockAnalyticsClientMockRecorder) PromptOptInIfNeeded(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PromptOptInIfNeeded", reflect.TypeOf((*MockAnalyticsClient)(nil).PromptOptInIfNeeded), arg0)
}
