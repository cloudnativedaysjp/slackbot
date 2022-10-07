// Code generated by MockGen. DO NOT EDIT.
// Source: slack_factory.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	slack "github.com/cloudnativedaysjp/seaman/seaman/infra/slack"
	gomock "github.com/golang/mock/gomock"
	slack0 "github.com/slack-go/slack"
)

// MockSlackClientFactory is a mock of SlackClientFactory interface.
type MockSlackClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockSlackClientFactoryMockRecorder
}

// MockSlackClientFactoryMockRecorder is the mock recorder for MockSlackClientFactory.
type MockSlackClientFactoryMockRecorder struct {
	mock *MockSlackClientFactory
}

// NewMockSlackClientFactory creates a new mock instance.
func NewMockSlackClientFactory(ctrl *gomock.Controller) *MockSlackClientFactory {
	mock := &MockSlackClientFactory{ctrl: ctrl}
	mock.recorder = &MockSlackClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSlackClientFactory) EXPECT() *MockSlackClientFactoryMockRecorder {
	return m.recorder
}

// New mocks base method.
func (m *MockSlackClientFactory) New(client slack0.Client) (slack.SlackClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", client)
	ret0, _ := ret[0].(slack.SlackClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New.
func (mr *MockSlackClientFactoryMockRecorder) New(client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockSlackClientFactory)(nil).New), client)
}
