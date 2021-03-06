// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/tddbc/go_gotest/domain"
)

// MockFizzBuzzRepository is a mock of FizzBuzzRepository interface.
type MockFizzBuzzRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFizzBuzzRepositoryMockRecorder
}

// MockFizzBuzzRepositoryMockRecorder is the mock recorder for MockFizzBuzzRepository.
type MockFizzBuzzRepositoryMockRecorder struct {
	mock *MockFizzBuzzRepository
}

// NewMockFizzBuzzRepository creates a new mock instance.
func NewMockFizzBuzzRepository(ctrl *gomock.Controller) *MockFizzBuzzRepository {
	mock := &MockFizzBuzzRepository{ctrl: ctrl}
	mock.recorder = &MockFizzBuzzRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFizzBuzzRepository) EXPECT() *MockFizzBuzzRepositoryMockRecorder {
	return m.recorder
}

// Append mocks base method.
func (m *MockFizzBuzzRepository) Append(f *domain.FizzBuzz) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Append", f)
}

// Append indicates an expected call of Append.
func (mr *MockFizzBuzzRepositoryMockRecorder) Append(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockFizzBuzzRepository)(nil).Append), f)
}

// List mocks base method.
func (m *MockFizzBuzzRepository) List() []*domain.FizzBuzz {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*domain.FizzBuzz)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockFizzBuzzRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFizzBuzzRepository)(nil).List))
}
