// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\erons\TESTESAPI\config\database\database.go
//
// Generated by this command:
//
//	mockgen -source=C:\Users\erons\TESTESAPI\config\database\database.go -destination=C:\Users\erons\TESTESAPI\database\gomock\mocks_database.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	sql "database/sql"
	reflect "reflect"

	database "github.com/eron97/testesAPI/config/database"
	models "github.com/eron97/testesAPI/config/models"
	gomock "go.uber.org/mock/gomock"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// CloseDB mocks base method.
func (m *MockDatabase) CloseDB() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseDB")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseDB indicates an expected call of CloseDB.
func (mr *MockDatabaseMockRecorder) CloseDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseDB", reflect.TypeOf((*MockDatabase)(nil).CloseDB))
}

// Create mocks base method.
func (m *MockDatabase) Create(data any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDatabaseMockRecorder) Create(data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDatabase)(nil).Create), data)
}

// Delete mocks base method.
func (m *MockDatabase) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDatabaseMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDatabase)(nil).Delete), id)
}

// InitDB mocks base method.
func (m *MockDatabase) InitDB(config database.Config) (*sql.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitDB", config)
	ret0, _ := ret[0].(*sql.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitDB indicates an expected call of InitDB.
func (mr *MockDatabaseMockRecorder) InitDB(config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitDB", reflect.TypeOf((*MockDatabase)(nil).InitDB), config)
}

// Read mocks base method.
func (m *MockDatabase) Read(condition string) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", condition)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockDatabaseMockRecorder) Read(condition any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockDatabase)(nil).Read), condition)
}

// Update mocks base method.
func (m *MockDatabase) Update(data any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDatabaseMockRecorder) Update(data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDatabase)(nil).Update), data)
}