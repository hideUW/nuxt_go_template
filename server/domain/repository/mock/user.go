// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/user.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/hideUW/nuxt-go-chat-app/server/domain/model"
	repository "github.com/hideUW/nuxt-go-chat-app/server/domain/repository"
)

// MockUserRepository is a mock of UserRepository interface
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// GetUserByID mocks base method
func (m_2 *MockUserRepository) GetUserByID(m repository.SQLManager, id uint32) (*model.User, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "GetUserByID", m, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID
func (mr *MockUserRepositoryMockRecorder) GetUserByID(m, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserRepository)(nil).GetUserByID), m, id)
}

// GetUserByName mocks base method
func (m_2 *MockUserRepository) GetUserByName(m repository.SQLManager, name string) (*model.User, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "GetUserByName", m, name)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName
func (mr *MockUserRepositoryMockRecorder) GetUserByName(m, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockUserRepository)(nil).GetUserByName), m, name)
}

// InsertUser mocks base method
func (m_2 *MockUserRepository) InsertUser(m repository.SQLManager, user *model.User) (uint32, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "InsertUser", m, user)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser
func (mr *MockUserRepositoryMockRecorder) InsertUser(m, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockUserRepository)(nil).InsertUser), m, user)
}

// UpdateUser mocks base method
func (m_2 *MockUserRepository) UpdateUser(m repository.SQLManager, id uint32, user *model.User) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "UpdateUser", m, id, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockUserRepositoryMockRecorder) UpdateUser(m, id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepository)(nil).UpdateUser), m, id, user)
}

// DeleteUser mocks base method
func (m_2 *MockUserRepository) DeleteUser(m repository.SQLManager, id uint32) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "DeleteUser", m, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockUserRepositoryMockRecorder) DeleteUser(m, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepository)(nil).DeleteUser), m, id)
}
