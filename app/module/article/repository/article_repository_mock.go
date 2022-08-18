// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bangadam/go-fiber-starter/app/module/article/repository (interfaces: IArticleRepository)

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	ent "github.com/bangadam/go-fiber-starter/internal/ent"
	gomock "github.com/golang/mock/gomock"
)

// MockIArticleRepository is a mock of IArticleRepository interface.
type MockIArticleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIArticleRepositoryMockRecorder
}

// MockIArticleRepositoryMockRecorder is the mock recorder for MockIArticleRepository.
type MockIArticleRepositoryMockRecorder struct {
	mock *MockIArticleRepository
}

// NewMockIArticleRepository creates a new mock instance.
func NewMockIArticleRepository(ctrl *gomock.Controller) *MockIArticleRepository {
	mock := &MockIArticleRepository{ctrl: ctrl}
	mock.recorder = &MockIArticleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIArticleRepository) EXPECT() *MockIArticleRepositoryMockRecorder {
	return m.recorder
}

// GetArticles mocks base method.
func (m *MockIArticleRepository) GetArticles() ([]*ent.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticles")
	ret0, _ := ret[0].([]*ent.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticles indicates an expected call of GetArticles.
func (mr *MockIArticleRepositoryMockRecorder) GetArticles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticles", reflect.TypeOf((*MockIArticleRepository)(nil).GetArticles))
}