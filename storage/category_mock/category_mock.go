// Code generated by MockGen. DO NOT EDIT.
// Source: storage/category.go

// Package categorymock is a generated GoMock package.
package categorymock

import (
	orm "github.com/astaxie/beego/orm"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	models "shop-api/models"
)

// MockCategory is a mock of Category interface
type MockCategory struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryMockRecorder
}

// MockCategoryMockRecorder is the mock recorder for MockCategory
type MockCategoryMockRecorder struct {
	mock *MockCategory
}

// NewMockCategory creates a new mock instance
func NewMockCategory(ctrl *gomock.Controller) *MockCategory {
	mock := &MockCategory{ctrl: ctrl}
	mock.recorder = &MockCategoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCategory) EXPECT() *MockCategoryMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockCategory) Add(arg0 orm.Ormer, arg1 *models.Category) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockCategoryMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCategory)(nil).Add), arg0, arg1)
}

// Delete mocks base method
func (m *MockCategory) Delete(arg0 orm.Ormer, arg1 *models.Category) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockCategoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCategory)(nil).Delete), arg0, arg1)
}

// GetAll mocks base method
func (m *MockCategory) GetAll(arg0 orm.Ormer, arg1 map[string]string, arg2 []string, arg3, arg4 int64) ([]models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockCategoryMockRecorder) GetAll(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCategory)(nil).GetAll), arg0, arg1, arg2, arg3, arg4)
}

// GetByID mocks base method
func (m *MockCategory) GetByID(arg0 orm.Ormer, arg1 int64) (models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockCategoryMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockCategory)(nil).GetByID), arg0, arg1)
}

// GetByName mocks base method
func (m *MockCategory) GetByName(arg0 orm.Ormer, arg1 string) (models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", arg0, arg1)
	ret0, _ := ret[0].(models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName
func (mr *MockCategoryMockRecorder) GetByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockCategory)(nil).GetByName), arg0, arg1)
}

// UpdateByID mocks base method
func (m *MockCategory) UpdateByID(arg0 orm.Ormer, arg1 *models.Category) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByID indicates an expected call of UpdateByID
func (mr *MockCategoryMockRecorder) UpdateByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockCategory)(nil).UpdateByID), arg0, arg1)
}
