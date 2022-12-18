// Code generated by MockGen. DO NOT EDIT.
// Source: services/cypto_service.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCyptoService is a mock of CyptoService interface.
type MockCyptoService struct {
	ctrl     *gomock.Controller
	recorder *MockCyptoServiceMockRecorder
}

// MockCyptoServiceMockRecorder is the mock recorder for MockCyptoService.
type MockCyptoServiceMockRecorder struct {
	mock *MockCyptoService
}

// NewMockCyptoService creates a new mock instance.
func NewMockCyptoService(ctrl *gomock.Controller) *MockCyptoService {
	mock := &MockCyptoService{ctrl: ctrl}
	mock.recorder = &MockCyptoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCyptoService) EXPECT() *MockCyptoServiceMockRecorder {
	return m.recorder
}

// ComparePasswords mocks base method.
func (m *MockCyptoService) ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComparePasswords", hashedPwd, plainPwd)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ComparePasswords indicates an expected call of ComparePasswords.
func (mr *MockCyptoServiceMockRecorder) ComparePasswords(hashedPwd, plainPwd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComparePasswords", reflect.TypeOf((*MockCyptoService)(nil).ComparePasswords), hashedPwd, plainPwd)
}

// HashAndSalt mocks base method.
func (m *MockCyptoService) HashAndSalt(pwd []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashAndSalt", pwd)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashAndSalt indicates an expected call of HashAndSalt.
func (mr *MockCyptoServiceMockRecorder) HashAndSalt(pwd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashAndSalt", reflect.TypeOf((*MockCyptoService)(nil).HashAndSalt), pwd)
}
