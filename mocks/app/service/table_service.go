// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/accalina/restaurant-mgmt/app/model"
	mock "github.com/stretchr/testify/mock"
)

// TableService is an autogenerated mock type for the TableService type
type TableService struct {
	mock.Mock
}

// CreateTable provides a mock function with given fields: ctx, data
func (_m *TableService) CreateTable(ctx context.Context, data model.TableCreateOrUpdateModel) (*model.TableResponse, error) {
	ret := _m.Called(ctx, data)

	var r0 *model.TableResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.TableCreateOrUpdateModel) (*model.TableResponse, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.TableCreateOrUpdateModel) *model.TableResponse); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TableResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.TableCreateOrUpdateModel) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTable provides a mock function with given fields: ctx, id
func (_m *TableService) DeleteTable(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTable provides a mock function with given fields: ctx, filter
func (_m *TableService) GetAllTable(ctx context.Context, filter *model.TableFilter) ([]model.TableResponse, model.Meta, error) {
	ret := _m.Called(ctx, filter)

	var r0 []model.TableResponse
	var r1 model.Meta
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.TableFilter) ([]model.TableResponse, model.Meta, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.TableFilter) []model.TableResponse); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.TableResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.TableFilter) model.Meta); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(model.Meta)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *model.TableFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetDetailTable provides a mock function with given fields: ctx, id
func (_m *TableService) GetDetailTable(ctx context.Context, id string) (*model.TableResponse, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.TableResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.TableResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.TableResponse); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TableResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTable provides a mock function with given fields: ctx, data
func (_m *TableService) UpdateTable(ctx context.Context, data model.TableCreateOrUpdateModel) (*model.TableResponse, error) {
	ret := _m.Called(ctx, data)

	var r0 *model.TableResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.TableCreateOrUpdateModel) (*model.TableResponse, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.TableCreateOrUpdateModel) *model.TableResponse); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TableResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.TableCreateOrUpdateModel) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTableService interface {
	mock.TestingT
	Cleanup(func())
}

// NewTableService creates a new instance of TableService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTableService(t mockConstructorTestingTNewTableService) *TableService {
	mock := &TableService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
