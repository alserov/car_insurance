// Code generated by MockGen. DO NOT EDIT.
// Source: .\internal\db\repository.go

// Package mock is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/alserov/car_insurance/insurance/internal/service/models"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateInsuranceData mocks base method.
func (m *MockRepository) CreateInsuranceData(ctx context.Context, insData models.InsuranceData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInsuranceData", ctx, insData)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInsuranceData indicates an expected call of CreateInsuranceData.
func (mr *MockRepositoryMockRecorder) CreateInsuranceData(ctx, insData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInsuranceData", reflect.TypeOf((*MockRepository)(nil).CreateInsuranceData), ctx, insData)
}

// GetInsuranceData mocks base method.
func (m *MockRepository) GetInsuranceData(ctx context.Context, ownerAddr string) (models.InsuranceData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInsuranceData", ctx, ownerAddr)
	ret0, _ := ret[0].(models.InsuranceData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInsuranceData indicates an expected call of GetInsuranceData.
func (mr *MockRepositoryMockRecorder) GetInsuranceData(ctx, ownerAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInsuranceData", reflect.TypeOf((*MockRepository)(nil).GetInsuranceData), ctx, ownerAddr)
}

// UpdateInsuranceStatus mocks base method.
func (m *MockRepository) UpdateInsuranceStatus(ctx context.Context, id string, status uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInsuranceStatus", ctx, id, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInsuranceStatus indicates an expected call of UpdateInsuranceStatus.
func (mr *MockRepositoryMockRecorder) UpdateInsuranceStatus(ctx, id, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInsuranceStatus", reflect.TypeOf((*MockRepository)(nil).UpdateInsuranceStatus), ctx, id, status)
}

// MockOutbox is a mock of Outbox interface.
type MockOutbox struct {
	ctrl     *gomock.Controller
	recorder *MockOutboxMockRecorder
}

// MockOutboxMockRecorder is the mock recorder for MockOutbox.
type MockOutboxMockRecorder struct {
	mock *MockOutbox
}

// NewMockOutbox creates a new mock instance.
func NewMockOutbox(ctrl *gomock.Controller) *MockOutbox {
	mock := &MockOutbox{ctrl: ctrl}
	mock.recorder = &MockOutboxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOutbox) EXPECT() *MockOutboxMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOutbox) Create(ctx context.Context, item models.OutboxItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockOutboxMockRecorder) Create(ctx, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOutbox)(nil).Create), ctx, item)
}

// Delete mocks base method.
func (m *MockOutbox) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOutboxMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOutbox)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockOutbox) Get(ctx context.Context, status, groupID int) ([]models.OutboxItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, status, groupID)
	ret0, _ := ret[0].([]models.OutboxItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOutboxMockRecorder) Get(ctx, status, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOutbox)(nil).Get), ctx, status, groupID)
}
