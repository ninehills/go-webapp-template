// Code generated by MockGen. DO NOT EDIT.
// Source: internal/dao/querier.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dao "github.com/ninehills/go-webapp-template/internal/dao"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockQuerier) CreateUser(ctx context.Context, arg dao.CreateUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockQuerierMockRecorder) CreateUser(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockQuerier)(nil).CreateUser), ctx, arg)
}

// DeleteUser mocks base method.
func (m *MockQuerier) DeleteUser(ctx context.Context, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockQuerierMockRecorder) DeleteUser(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockQuerier)(nil).DeleteUser), ctx, username)
}

// GetUser mocks base method.
func (m *MockQuerier) GetUser(ctx context.Context, username string) (dao.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, username)
	ret0, _ := ret[0].(dao.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockQuerierMockRecorder) GetUser(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockQuerier)(nil).GetUser), ctx, username)
}

// ListUser mocks base method.
func (m *MockQuerier) ListUser(ctx context.Context, arg dao.ListUserParams) ([]dao.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUser", ctx, arg)
	ret0, _ := ret[0].([]dao.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUser indicates an expected call of ListUser.
func (mr *MockQuerierMockRecorder) ListUser(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUser", reflect.TypeOf((*MockQuerier)(nil).ListUser), ctx, arg)
}

// QueryUser mocks base method.
func (m *MockQuerier) QueryUser(ctx context.Context, arg dao.QueryUserParams) ([]dao.User, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryUser", ctx, arg)
	ret0, _ := ret[0].([]dao.User)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryUser indicates an expected call of QueryUser.
func (mr *MockQuerierMockRecorder) QueryUser(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryUser", reflect.TypeOf((*MockQuerier)(nil).QueryUser), ctx, arg)
}

// UpdateUser mocks base method.
func (m *MockQuerier) UpdateUser(ctx context.Context, arg dao.UpdateUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockQuerierMockRecorder) UpdateUser(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockQuerier)(nil).UpdateUser), ctx, arg)
}
