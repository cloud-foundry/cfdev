// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cfdev/cmd/deploy-service (interfaces: MetaDataReader)

// Package mocks is a generated GoMock package.
package mocks

import (
	metadata "code.cloudfoundry.org/cfdev/metadata"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMetaDataReader is a mock of MetaDataReader interface
type MockMetaDataReader struct {
	ctrl     *gomock.Controller
	recorder *MockMetaDataReaderMockRecorder
}

// MockMetaDataReaderMockRecorder is the mock recorder for MockMetaDataReader
type MockMetaDataReaderMockRecorder struct {
	mock *MockMetaDataReader
}

// NewMockMetaDataReader creates a new mock instance
func NewMockMetaDataReader(ctrl *gomock.Controller) *MockMetaDataReader {
	mock := &MockMetaDataReader{ctrl: ctrl}
	mock.recorder = &MockMetaDataReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetaDataReader) EXPECT() *MockMetaDataReaderMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockMetaDataReader) Read(arg0 string) (metadata.Metadata, error) {
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(metadata.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockMetaDataReaderMockRecorder) Read(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockMetaDataReader)(nil).Read), arg0)
}
