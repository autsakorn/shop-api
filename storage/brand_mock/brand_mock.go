// Code generated by MockGen. DO NOT EDIT.
// Source: storage/brand.go

// Package brandmock is a generated GoMock package.
package brandmock

import (
	orm "github.com/astaxie/beego/orm"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	models "shop-api/models"
)

// MockBrand is a mock of Brand interface
type MockBrand struct {
	ctrl     *gomock.Controller
	recorder *MockBrandMockRecorder
}

// MockBrandMockRecorder is the mock recorder for MockBrand
type MockBrandMockRecorder struct {
	mock *MockBrand
}

// NewMockBrand creates a new mock instance
func NewMockBrand(ctrl *gomock.Controller) *MockBrand {
	mock := &MockBrand{ctrl: ctrl}
	mock.recorder = &MockBrandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBrand) EXPECT() *MockBrandMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockBrand) Add(arg0 orm.Ormer, arg1 *models.Brand) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockBrandMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockBrand)(nil).Add), arg0, arg1)
}

// Delete mocks base method
func (m *MockBrand) Delete(arg0 orm.Ormer, arg1 *models.Brand) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockBrandMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBrand)(nil).Delete), arg0, arg1)
}

// GetAll mocks base method
func (m *MockBrand) GetAll(arg0 orm.Ormer, arg1 map[string]string, arg2 []string, arg3, arg4 int64) ([]models.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]models.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockBrandMockRecorder) GetAll(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockBrand)(nil).GetAll), arg0, arg1, arg2, arg3, arg4)
}

// GetByID mocks base method
func (m *MockBrand) GetByID(arg0 orm.Ormer, arg1 int64) (models.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(models.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockBrandMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockBrand)(nil).GetByID), arg0, arg1)
}

// UpdateByID mocks base method
func (m *MockBrand) UpdateByID(arg0 orm.Ormer, arg1 *models.Brand) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByID indicates an expected call of UpdateByID
func (mr *MockBrandMockRecorder) UpdateByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockBrand)(nil).UpdateByID), arg0, arg1)
}