// Code generated by MockGen. DO NOT EDIT.
// Source: service/user/service.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	user "github.com/kanca-studio/palang/service/user"
)

// UserService is a mock of Service interface
type UserService struct {
	ctrl     *gomock.Controller
	recorder *UserServiceMockRecorder
}

// UserServiceMockRecorder is the mock recorder for UserService
type UserServiceMockRecorder struct {
	mock *UserService
}

// NewUserService creates a new mock instance
func NewUserService(ctrl *gomock.Controller) *UserService {
	mock := &UserService{ctrl: ctrl}
	mock.recorder = &UserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *UserService) EXPECT() *UserServiceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *UserService) Create(param interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", param)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *UserServiceMockRecorder) Create(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*UserService)(nil).Create), param)
}

// FindById mocks base method
func (m *UserService) FindById(id uint, out interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindById indicates an expected call of FindById
func (mr *UserServiceMockRecorder) FindById(id, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*UserService)(nil).FindById), id, out)
}

// Find mocks base method
func (m *UserService) Find(filter, out interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", filter, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// Find indicates an expected call of Find
func (mr *UserServiceMockRecorder) Find(filter, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*UserService)(nil).Find), filter, out)
}

// FindAll mocks base method
func (m *UserService) FindAll(filter, outs interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", filter, outs)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindAll indicates an expected call of FindAll
func (mr *UserServiceMockRecorder) FindAll(filter, outs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*UserService)(nil).FindAll), filter, outs)
}

// Update mocks base method
func (m *UserService) Update(filter, param interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", filter, param)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *UserServiceMockRecorder) Update(filter, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*UserService)(nil).Update), filter, param)
}

// Delete mocks base method
func (m *UserService) Delete(filter interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *UserServiceMockRecorder) Delete(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*UserService)(nil).Delete), filter)
}

// Remove mocks base method
func (m *UserService) Remove(filter interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *UserServiceMockRecorder) Remove(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*UserService)(nil).Remove), filter)
}

// CreateUser mocks base method
func (m *UserService) CreateUser(identifierTypeStr user.IdentifierType, identifier, password string) (user.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", identifierTypeStr, identifier, password)
	ret0, _ := ret[0].(user.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *UserServiceMockRecorder) CreateUser(identifierTypeStr, identifier, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*UserService)(nil).CreateUser), identifierTypeStr, identifier, password)
}

// GetUserByIdentifier mocks base method
func (m *UserService) GetUserByIdentifier(identifierTypeStr user.IdentifierType, identifier string) (user.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByIdentifier", identifierTypeStr, identifier)
	ret0, _ := ret[0].(user.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByIdentifier indicates an expected call of GetUserByIdentifier
func (mr *UserServiceMockRecorder) GetUserByIdentifier(identifierTypeStr, identifier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByIdentifier", reflect.TypeOf((*UserService)(nil).GetUserByIdentifier), identifierTypeStr, identifier)
}

// IdentifierTypeToConst mocks base method
func (m *UserService) IdentifierTypeToConst(identifierType string) user.IdentifierType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdentifierTypeToConst", identifierType)
	ret0, _ := ret[0].(user.IdentifierType)
	return ret0
}

// IdentifierTypeToConst indicates an expected call of IdentifierTypeToConst
func (mr *UserServiceMockRecorder) IdentifierTypeToConst(identifierType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdentifierTypeToConst", reflect.TypeOf((*UserService)(nil).IdentifierTypeToConst), identifierType)
}
