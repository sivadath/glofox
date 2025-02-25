// Code generated by MockGen. DO NOT EDIT.
// Source: storage/storage.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	models "github.com/sivadath/glofox/models"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AddBooking mocks base method.
func (m *MockStorage) AddBooking(ctx context.Context, booking models.Booking) (models.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBooking", ctx, booking)
	ret0, _ := ret[0].(models.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBooking indicates an expected call of AddBooking.
func (mr *MockStorageMockRecorder) AddBooking(ctx, booking interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBooking", reflect.TypeOf((*MockStorage)(nil).AddBooking), ctx, booking)
}

// AddClass mocks base method.
func (m *MockStorage) AddClass(ctx context.Context, class models.Class) (models.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddClass", ctx, class)
	ret0, _ := ret[0].(models.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddClass indicates an expected call of AddClass.
func (mr *MockStorageMockRecorder) AddClass(ctx, class interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddClass", reflect.TypeOf((*MockStorage)(nil).AddClass), ctx, class)
}

// GetBookings mocks base method.
func (m *MockStorage) GetBookings(ctx context.Context) ([]models.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookings", ctx)
	ret0, _ := ret[0].([]models.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookings indicates an expected call of GetBookings.
func (mr *MockStorageMockRecorder) GetBookings(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookings", reflect.TypeOf((*MockStorage)(nil).GetBookings), ctx)
}

// GetClasses mocks base method.
func (m *MockStorage) GetClasses(ctx context.Context) ([]models.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClasses", ctx)
	ret0, _ := ret[0].([]models.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClasses indicates an expected call of GetClasses.
func (mr *MockStorageMockRecorder) GetClasses(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClasses", reflect.TypeOf((*MockStorage)(nil).GetClasses), ctx)
}

// GetClassesByDate mocks base method.
func (m *MockStorage) GetClassesByDate(ctx context.Context, date time.Time) ([]models.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClassesByDate", ctx, date)
	ret0, _ := ret[0].([]models.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClassesByDate indicates an expected call of GetClassesByDate.
func (mr *MockStorageMockRecorder) GetClassesByDate(ctx, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClassesByDate", reflect.TypeOf((*MockStorage)(nil).GetClassesByDate), ctx, date)
}

// MockClassSchema is a mock of ClassSchema interface.
type MockClassSchema struct {
	ctrl     *gomock.Controller
	recorder *MockClassSchemaMockRecorder
}

// MockClassSchemaMockRecorder is the mock recorder for MockClassSchema.
type MockClassSchemaMockRecorder struct {
	mock *MockClassSchema
}

// NewMockClassSchema creates a new mock instance.
func NewMockClassSchema(ctrl *gomock.Controller) *MockClassSchema {
	mock := &MockClassSchema{ctrl: ctrl}
	mock.recorder = &MockClassSchemaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClassSchema) EXPECT() *MockClassSchemaMockRecorder {
	return m.recorder
}

// AddClass mocks base method.
func (m *MockClassSchema) AddClass(ctx context.Context, class models.Class) (models.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddClass", ctx, class)
	ret0, _ := ret[0].(models.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddClass indicates an expected call of AddClass.
func (mr *MockClassSchemaMockRecorder) AddClass(ctx, class interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddClass", reflect.TypeOf((*MockClassSchema)(nil).AddClass), ctx, class)
}

// GetClasses mocks base method.
func (m *MockClassSchema) GetClasses(ctx context.Context) ([]models.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClasses", ctx)
	ret0, _ := ret[0].([]models.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClasses indicates an expected call of GetClasses.
func (mr *MockClassSchemaMockRecorder) GetClasses(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClasses", reflect.TypeOf((*MockClassSchema)(nil).GetClasses), ctx)
}

// GetClassesByDate mocks base method.
func (m *MockClassSchema) GetClassesByDate(ctx context.Context, date time.Time) ([]models.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClassesByDate", ctx, date)
	ret0, _ := ret[0].([]models.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClassesByDate indicates an expected call of GetClassesByDate.
func (mr *MockClassSchemaMockRecorder) GetClassesByDate(ctx, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClassesByDate", reflect.TypeOf((*MockClassSchema)(nil).GetClassesByDate), ctx, date)
}

// MockBookingSchema is a mock of BookingSchema interface.
type MockBookingSchema struct {
	ctrl     *gomock.Controller
	recorder *MockBookingSchemaMockRecorder
}

// MockBookingSchemaMockRecorder is the mock recorder for MockBookingSchema.
type MockBookingSchemaMockRecorder struct {
	mock *MockBookingSchema
}

// NewMockBookingSchema creates a new mock instance.
func NewMockBookingSchema(ctrl *gomock.Controller) *MockBookingSchema {
	mock := &MockBookingSchema{ctrl: ctrl}
	mock.recorder = &MockBookingSchemaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookingSchema) EXPECT() *MockBookingSchemaMockRecorder {
	return m.recorder
}

// AddBooking mocks base method.
func (m *MockBookingSchema) AddBooking(ctx context.Context, booking models.Booking) (models.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBooking", ctx, booking)
	ret0, _ := ret[0].(models.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBooking indicates an expected call of AddBooking.
func (mr *MockBookingSchemaMockRecorder) AddBooking(ctx, booking interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBooking", reflect.TypeOf((*MockBookingSchema)(nil).AddBooking), ctx, booking)
}

// GetBookings mocks base method.
func (m *MockBookingSchema) GetBookings(ctx context.Context) ([]models.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookings", ctx)
	ret0, _ := ret[0].([]models.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookings indicates an expected call of GetBookings.
func (mr *MockBookingSchemaMockRecorder) GetBookings(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookings", reflect.TypeOf((*MockBookingSchema)(nil).GetBookings), ctx)
}
