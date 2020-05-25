// Code generated by MockGen. DO NOT EDIT.
// Source: storage/client.go

// Package clientmock is a generated GoMock package.
package clientmock

import (
	orm "github.com/astaxie/beego/orm"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetByXApiKeyAndStatus mocks base method
func (m *MockClient) GetByXApiKeyAndStatus(arg0 orm.Ormer, arg1 string, arg2 int32) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByXApiKeyAndStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByXApiKeyAndStatus indicates an expected call of GetByXApiKeyAndStatus
func (mr *MockClientMockRecorder) GetByXApiKeyAndStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByXApiKeyAndStatus", reflect.TypeOf((*MockClient)(nil).GetByXApiKeyAndStatus), arg0, arg1, arg2)
}