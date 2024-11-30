// Code generated by MockGen. DO NOT EDIT.
// Source: db.go
//
// Generated by this command:
//
//	mockgen -source db.go -destination db_mock.go -package db
//

// Package db is a generated GoMock package.
package db

import (
	reflect "reflect"
	time "time"

	model "github.com/icchy-san/uctest/db/model"
	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
	isgomock struct{}
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// GetInvoices mocks base method.
func (m *MockDB) GetInvoices(tx *gorm.DB, period_start_at, period_end_at *time.Time) ([]model.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInvoices", tx, period_start_at, period_end_at)
	ret0, _ := ret[0].([]model.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInvoices indicates an expected call of GetInvoices.
func (mr *MockDBMockRecorder) GetInvoices(tx, period_start_at, period_end_at any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvoices", reflect.TypeOf((*MockDB)(nil).GetInvoices), tx, period_start_at, period_end_at)
}
