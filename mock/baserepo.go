// Code generated by MockGen. DO NOT EDIT.
// Source: service/base/baserepository.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// BaseRepository is a mock of InterfaceBaseRepository interface
type BaseRepository struct {
	ctrl     *gomock.Controller
	recorder *BaseRepositoryMockRecorder
}

// BaseRepositoryMockRecorder is the mock recorder for BaseRepository
type BaseRepositoryMockRecorder struct {
	mock *BaseRepository
}

// NewBaseRepository creates a new mock instance
func NewBaseRepository(ctrl *gomock.Controller) *BaseRepository {
	mock := &BaseRepository{ctrl: ctrl}
	mock.recorder = &BaseRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *BaseRepository) EXPECT() *BaseRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *BaseRepository) Create(param interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", param)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *BaseRepositoryMockRecorder) Create(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*BaseRepository)(nil).Create), param)
}

// FindById mocks base method
func (m *BaseRepository) FindById(id uint, out interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindById indicates an expected call of FindById
func (mr *BaseRepositoryMockRecorder) FindById(id, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*BaseRepository)(nil).FindById), id, out)
}

// Find mocks base method
func (m *BaseRepository) Find(filter, out interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", filter, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// Find indicates an expected call of Find
func (mr *BaseRepositoryMockRecorder) Find(filter, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*BaseRepository)(nil).Find), filter, out)
}

// FindAll mocks base method
func (m *BaseRepository) FindAll(filter, outs interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", filter, outs)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindAll indicates an expected call of FindAll
func (mr *BaseRepositoryMockRecorder) FindAll(filter, outs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*BaseRepository)(nil).FindAll), filter, outs)
}

// Update mocks base method
func (m *BaseRepository) Update(filter, param interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", filter, param)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *BaseRepositoryMockRecorder) Update(filter, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*BaseRepository)(nil).Update), filter, param)
}

// Delete mocks base method
func (m *BaseRepository) Delete(filter interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *BaseRepositoryMockRecorder) Delete(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*BaseRepository)(nil).Delete), filter)
}

// Remove mocks base method
func (m *BaseRepository) Remove(filter interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *BaseRepositoryMockRecorder) Remove(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*BaseRepository)(nil).Remove), filter)
}
