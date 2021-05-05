// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger-labs/minbft/core/internal/clientstate (interfaces: Provider,State)

// Package mock_clientstate is a generated GoMock package.
package mock_clientstate

import (
	gomock "github.com/golang/mock/gomock"
	clientstate "github.com/hyperledger-labs/minbft/core/internal/clientstate"
	messages "github.com/hyperledger-labs/minbft/messages"
	reflect "reflect"
)

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// ClientState mocks base method
func (m *MockProvider) ClientState(arg0 uint32) clientstate.State {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientState", arg0)
	ret0, _ := ret[0].(clientstate.State)
	return ret0
}

// ClientState indicates an expected call of ClientState
func (mr *MockProviderMockRecorder) ClientState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientState", reflect.TypeOf((*MockProvider)(nil).ClientState), arg0)
}

// Clients mocks base method
func (m *MockProvider) Clients() []uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clients")
	ret0, _ := ret[0].([]uint32)
	return ret0
}

// Clients indicates an expected call of Clients
func (mr *MockProviderMockRecorder) Clients() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clients", reflect.TypeOf((*MockProvider)(nil).Clients))
}

// MockState is a mock of State interface
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// AddReply mocks base method
func (m *MockState) AddReply(arg0 messages.Reply) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddReply", arg0)
}

// AddReply indicates an expected call of AddReply
func (mr *MockStateMockRecorder) AddReply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddReply", reflect.TypeOf((*MockState)(nil).AddReply), arg0)
}

// CaptureRequestSeq mocks base method
func (m *MockState) CaptureRequestSeq(arg0 uint64) (bool, func()) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CaptureRequestSeq", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(func())
	return ret0, ret1
}

// CaptureRequestSeq indicates an expected call of CaptureRequestSeq
func (mr *MockStateMockRecorder) CaptureRequestSeq(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CaptureRequestSeq", reflect.TypeOf((*MockState)(nil).CaptureRequestSeq), arg0)
}

// PrepareRequestSeq mocks base method
func (m *MockState) PrepareRequestSeq(arg0 uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareRequestSeq", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareRequestSeq indicates an expected call of PrepareRequestSeq
func (mr *MockStateMockRecorder) PrepareRequestSeq(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareRequestSeq", reflect.TypeOf((*MockState)(nil).PrepareRequestSeq), arg0)
}

// ReplyChannel mocks base method
func (m *MockState) ReplyChannel(arg0 uint64) <-chan messages.Reply {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplyChannel", arg0)
	ret0, _ := ret[0].(<-chan messages.Reply)
	return ret0
}

// ReplyChannel indicates an expected call of ReplyChannel
func (mr *MockStateMockRecorder) ReplyChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplyChannel", reflect.TypeOf((*MockState)(nil).ReplyChannel), arg0)
}

// RetireRequestSeq mocks base method
func (m *MockState) RetireRequestSeq(arg0 uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetireRequestSeq", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetireRequestSeq indicates an expected call of RetireRequestSeq
func (mr *MockStateMockRecorder) RetireRequestSeq(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetireRequestSeq", reflect.TypeOf((*MockState)(nil).RetireRequestSeq), arg0)
}

// StartPrepareTimer mocks base method
func (m *MockState) StartPrepareTimer(arg0 uint64, arg1 func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartPrepareTimer", arg0, arg1)
}

// StartPrepareTimer indicates an expected call of StartPrepareTimer
func (mr *MockStateMockRecorder) StartPrepareTimer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartPrepareTimer", reflect.TypeOf((*MockState)(nil).StartPrepareTimer), arg0, arg1)
}

// StartRequestTimer mocks base method
func (m *MockState) StartRequestTimer(arg0 uint64, arg1 func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartRequestTimer", arg0, arg1)
}

// StartRequestTimer indicates an expected call of StartRequestTimer
func (mr *MockStateMockRecorder) StartRequestTimer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartRequestTimer", reflect.TypeOf((*MockState)(nil).StartRequestTimer), arg0, arg1)
}

// StopAllTimers mocks base method
func (m *MockState) StopAllTimers() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopAllTimers")
}

// StopAllTimers indicates an expected call of StopAllTimers
func (mr *MockStateMockRecorder) StopAllTimers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopAllTimers", reflect.TypeOf((*MockState)(nil).StopAllTimers))
}

// StopPrepareTimer mocks base method
func (m *MockState) StopPrepareTimer(arg0 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopPrepareTimer", arg0)
}

// StopPrepareTimer indicates an expected call of StopPrepareTimer
func (mr *MockStateMockRecorder) StopPrepareTimer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopPrepareTimer", reflect.TypeOf((*MockState)(nil).StopPrepareTimer), arg0)
}

// StopRequestTimer mocks base method
func (m *MockState) StopRequestTimer(arg0 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopRequestTimer", arg0)
}

// StopRequestTimer indicates an expected call of StopRequestTimer
func (mr *MockStateMockRecorder) StopRequestTimer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopRequestTimer", reflect.TypeOf((*MockState)(nil).StopRequestTimer), arg0)
}

// UnprepareRequestSeq mocks base method
func (m *MockState) UnprepareRequestSeq() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnprepareRequestSeq")
}

// UnprepareRequestSeq indicates an expected call of UnprepareRequestSeq
func (mr *MockStateMockRecorder) UnprepareRequestSeq() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnprepareRequestSeq", reflect.TypeOf((*MockState)(nil).UnprepareRequestSeq))
}
