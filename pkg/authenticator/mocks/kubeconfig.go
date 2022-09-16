// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/authenticator/kubeconfig.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockKubeconfigClient is a mock of KubeconfigClient interface.
type MockKubeconfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockKubeconfigClientMockRecorder
}

// MockKubeconfigClientMockRecorder is the mock recorder for MockKubeconfigClient.
type MockKubeconfigClientMockRecorder struct {
	mock *MockKubeconfigClient
}

// NewMockKubeconfigClient creates a new mock instance.
func NewMockKubeconfigClient(ctrl *gomock.Controller) *MockKubeconfigClient {
	mock := &MockKubeconfigClient{ctrl: ctrl}
	mock.recorder = &MockKubeconfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubeconfigClient) EXPECT() *MockKubeconfigClientMockRecorder {
	return m.recorder
}

// GetKubeconfig mocks base method.
func (m *MockKubeconfigClient) GetKubeconfig(ctx context.Context, clusterName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKubeconfig", ctx, clusterName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKubeconfig indicates an expected call of GetKubeconfig.
func (mr *MockKubeconfigClientMockRecorder) GetKubeconfig(ctx, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKubeconfig", reflect.TypeOf((*MockKubeconfigClient)(nil).GetKubeconfig), ctx, clusterName)
}