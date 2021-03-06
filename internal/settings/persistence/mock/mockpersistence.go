// Code generated by MockGen. DO NOT EDIT.
// Source: persistence.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	persistence "github.com/n7down/kuiper/internal/settings/persistence"
	reflect "reflect"
)

// MockPersistence is a mock of Persistence interface.
type MockPersistence struct {
	ctrl     *gomock.Controller
	recorder *MockPersistenceMockRecorder
}

// MockPersistenceMockRecorder is the mock recorder for MockPersistence.
type MockPersistenceMockRecorder struct {
	mock *MockPersistence
}

// NewMockPersistence creates a new mock instance.
func NewMockPersistence(ctrl *gomock.Controller) *MockPersistence {
	mock := &MockPersistence{ctrl: ctrl}
	mock.recorder = &MockPersistenceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersistence) EXPECT() *MockPersistenceMockRecorder {
	return m.recorder
}

// CreateBatCaveSetting mocks base method.
func (m *MockPersistence) CreateBatCaveSetting(settings persistence.BatCaveSetting) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBatCaveSetting", settings)
	ret0, _ := ret[0].(int64)
	return ret0
}

// CreateBatCaveSetting indicates an expected call of CreateBatCaveSetting.
func (mr *MockPersistenceMockRecorder) CreateBatCaveSetting(settings interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBatCaveSetting", reflect.TypeOf((*MockPersistence)(nil).CreateBatCaveSetting), settings)
}

// GetBatCaveSetting mocks base method.
func (m *MockPersistence) GetBatCaveSetting(deviceID string) (bool, persistence.BatCaveSetting) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBatCaveSetting", deviceID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(persistence.BatCaveSetting)
	return ret0, ret1
}

// GetBatCaveSetting indicates an expected call of GetBatCaveSetting.
func (mr *MockPersistenceMockRecorder) GetBatCaveSetting(deviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBatCaveSetting", reflect.TypeOf((*MockPersistence)(nil).GetBatCaveSetting), deviceID)
}

// UpdateBatCaveSetting mocks base method.
func (m *MockPersistence) UpdateBatCaveSetting(settings persistence.BatCaveSetting) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBatCaveSetting", settings)
	ret0, _ := ret[0].(int64)
	return ret0
}

// UpdateBatCaveSetting indicates an expected call of UpdateBatCaveSetting.
func (mr *MockPersistenceMockRecorder) UpdateBatCaveSetting(settings interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBatCaveSetting", reflect.TypeOf((*MockPersistence)(nil).UpdateBatCaveSetting), settings)
}
