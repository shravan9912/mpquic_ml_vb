// Code generated by MockGen. DO NOT EDIT.
// Source: ../handshake/connection_parameters_manager.go

package mocks

import (
	time "time"

	"github.com/shravan9912/mpquic_ml_vb/internal/handshake"
	protocol "github.com/shravan9912/mpquic_ml_vb/internal/protocol"
	gomock "github.com/golang/mock/gomock"
)

// MockConnectionParametersManager is a mock of ConnectionParametersManager interface
type MockConnectionParametersManager struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionParametersManagerMockRecorder
}

// MockConnectionParametersManagerMockRecorder is the mock recorder for MockConnectionParametersManager
type MockConnectionParametersManagerMockRecorder struct {
	mock *MockConnectionParametersManager
}

// NewMockConnectionParametersManager creates a new mock instance
func NewMockConnectionParametersManager(ctrl *gomock.Controller) *MockConnectionParametersManager {
	mock := &MockConnectionParametersManager{ctrl: ctrl}
	mock.recorder = &MockConnectionParametersManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockConnectionParametersManager) EXPECT() *MockConnectionParametersManagerMockRecorder {
	return _m.recorder
}

// SetFromMap mocks base method
func (_m *MockConnectionParametersManager) SetFromMap(_param0 map[handshake.Tag][]byte) error {
	ret := _m.ctrl.Call(_m, "SetFromMap", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFromMap indicates an expected call of SetFromMap
func (_mr *MockConnectionParametersManagerMockRecorder) SetFromMap(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetFromMap", arg0)
}

// GetHelloMap mocks base method
func (_m *MockConnectionParametersManager) GetHelloMap() (map[handshake.Tag][]byte, error) {
	ret := _m.ctrl.Call(_m, "GetHelloMap")
	ret0, _ := ret[0].(map[handshake.Tag][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHelloMap indicates an expected call of GetHelloMap
func (_mr *MockConnectionParametersManagerMockRecorder) GetHelloMap() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetHelloMap")
}

// GetSendStreamFlowControlWindow mocks base method
func (_m *MockConnectionParametersManager) GetSendStreamFlowControlWindow() protocol.ByteCount {
	ret := _m.ctrl.Call(_m, "GetSendStreamFlowControlWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetSendStreamFlowControlWindow indicates an expected call of GetSendStreamFlowControlWindow
func (_mr *MockConnectionParametersManagerMockRecorder) GetSendStreamFlowControlWindow() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSendStreamFlowControlWindow")
}

// GetSendConnectionFlowControlWindow mocks base method
func (_m *MockConnectionParametersManager) GetSendConnectionFlowControlWindow() protocol.ByteCount {
	ret := _m.ctrl.Call(_m, "GetSendConnectionFlowControlWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetSendConnectionFlowControlWindow indicates an expected call of GetSendConnectionFlowControlWindow
func (_mr *MockConnectionParametersManagerMockRecorder) GetSendConnectionFlowControlWindow() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSendConnectionFlowControlWindow")
}

// GetReceiveStreamFlowControlWindow mocks base method
func (_m *MockConnectionParametersManager) GetReceiveStreamFlowControlWindow() protocol.ByteCount {
	ret := _m.ctrl.Call(_m, "GetReceiveStreamFlowControlWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetReceiveStreamFlowControlWindow indicates an expected call of GetReceiveStreamFlowControlWindow
func (_mr *MockConnectionParametersManagerMockRecorder) GetReceiveStreamFlowControlWindow() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetReceiveStreamFlowControlWindow")
}

// GetMaxReceiveStreamFlowControlWindow mocks base method
func (_m *MockConnectionParametersManager) GetMaxReceiveStreamFlowControlWindow() protocol.ByteCount {
	ret := _m.ctrl.Call(_m, "GetMaxReceiveStreamFlowControlWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetMaxReceiveStreamFlowControlWindow indicates an expected call of GetMaxReceiveStreamFlowControlWindow
func (_mr *MockConnectionParametersManagerMockRecorder) GetMaxReceiveStreamFlowControlWindow() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMaxReceiveStreamFlowControlWindow")
}

// GetReceiveConnectionFlowControlWindow mocks base method
func (_m *MockConnectionParametersManager) GetReceiveConnectionFlowControlWindow() protocol.ByteCount {
	ret := _m.ctrl.Call(_m, "GetReceiveConnectionFlowControlWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetReceiveConnectionFlowControlWindow indicates an expected call of GetReceiveConnectionFlowControlWindow
func (_mr *MockConnectionParametersManagerMockRecorder) GetReceiveConnectionFlowControlWindow() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetReceiveConnectionFlowControlWindow")
}

// GetMaxReceiveConnectionFlowControlWindow mocks base method
func (_m *MockConnectionParametersManager) GetMaxReceiveConnectionFlowControlWindow() protocol.ByteCount {
	ret := _m.ctrl.Call(_m, "GetMaxReceiveConnectionFlowControlWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetMaxReceiveConnectionFlowControlWindow indicates an expected call of GetMaxReceiveConnectionFlowControlWindow
func (_mr *MockConnectionParametersManagerMockRecorder) GetMaxReceiveConnectionFlowControlWindow() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMaxReceiveConnectionFlowControlWindow")
}

// GetMaxOutgoingStreams mocks base method
func (_m *MockConnectionParametersManager) GetMaxOutgoingStreams() uint32 {
	ret := _m.ctrl.Call(_m, "GetMaxOutgoingStreams")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetMaxOutgoingStreams indicates an expected call of GetMaxOutgoingStreams
func (_mr *MockConnectionParametersManagerMockRecorder) GetMaxOutgoingStreams() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMaxOutgoingStreams")
}

// GetMaxIncomingStreams mocks base method
func (_m *MockConnectionParametersManager) GetMaxIncomingStreams() uint32 {
	ret := _m.ctrl.Call(_m, "GetMaxIncomingStreams")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetMaxIncomingStreams indicates an expected call of GetMaxIncomingStreams
func (_mr *MockConnectionParametersManagerMockRecorder) GetMaxIncomingStreams() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMaxIncomingStreams")
}

// GetIdleConnectionStateLifetime mocks base method
func (_m *MockConnectionParametersManager) GetIdleConnectionStateLifetime() time.Duration {
	ret := _m.ctrl.Call(_m, "GetIdleConnectionStateLifetime")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetIdleConnectionStateLifetime indicates an expected call of GetIdleConnectionStateLifetime
func (_mr *MockConnectionParametersManagerMockRecorder) GetIdleConnectionStateLifetime() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetIdleConnectionStateLifetime")
}

// TruncateConnectionID mocks base method
func (_m *MockConnectionParametersManager) TruncateConnectionID() bool {
	ret := _m.ctrl.Call(_m, "TruncateConnectionID")
	ret0, _ := ret[0].(bool)
	return ret0
}

// TruncateConnectionID indicates an expected call of TruncateConnectionID
func (_mr *MockConnectionParametersManagerMockRecorder) TruncateConnectionID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TruncateConnectionID")
}
