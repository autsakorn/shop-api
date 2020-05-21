// Code generated by MockGen. DO NOT EDIT.
// Source: storage/product.go

// Package mock is a generated GoMock package.
package mock

import (
	orm "github.com/astaxie/beego/orm"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	models "shop-api/models"
)

// MockProduct is a mock of Product interface
type MockProduct struct {
	ctrl     *gomock.Controller
	recorder *MockProductMockRecorder
}

// MockProductMockRecorder is the mock recorder for MockProduct
type MockProductMockRecorder struct {
	mock *MockProduct
}

// NewMockProduct creates a new mock instance
func NewMockProduct(ctrl *gomock.Controller) *MockProduct {
	mock := &MockProduct{ctrl: ctrl}
	mock.recorder = &MockProductMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProduct) EXPECT() *MockProductMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockProduct) Add(arg0 orm.Ormer, arg1 *models.Product) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockProductMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockProduct)(nil).Add), arg0, arg1)
}

// Delete mocks base method
func (m *MockProduct) Delete(arg0 orm.Ormer, arg1 *models.Product) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockProductMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProduct)(nil).Delete), arg0, arg1)
}

// GetAll mocks base method
func (m *MockProduct) GetAll(arg0 orm.Ormer, arg1 map[string]string, arg2 []string, arg3, arg4 int64) ([]models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockProductMockRecorder) GetAll(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockProduct)(nil).GetAll), arg0, arg1, arg2, arg3, arg4)
}

// GetByID mocks base method
func (m *MockProduct) GetByID(arg0 orm.Ormer, arg1 int64) (models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockProductMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockProduct)(nil).GetByID), arg0, arg1)
}

// UpdateByID mocks base method
func (m *MockProduct) UpdateByID(arg0 orm.Ormer, arg1 *models.Product) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByID indicates an expected call of UpdateByID
func (mr *MockProductMockRecorder) UpdateByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockProduct)(nil).UpdateByID), arg0, arg1)
}
