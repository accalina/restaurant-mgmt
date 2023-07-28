// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/accalina/restaurant-mgmt/app/entity"
	mock "github.com/stretchr/testify/mock"

	model "github.com/accalina/restaurant-mgmt/app/model"
)

// InvoiceService is an autogenerated mock type for the InvoiceService type
type InvoiceService struct {
	mock.Mock
}

// CreateInvoice provides a mock function with given fields: ctx, data
func (_m *InvoiceService) CreateInvoice(ctx context.Context, data model.InvoiceCreateOrUpdateModel) (*model.InvoiceResponse, error) {
	ret := _m.Called(ctx, data)

	var r0 *model.InvoiceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.InvoiceCreateOrUpdateModel) (*model.InvoiceResponse, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.InvoiceCreateOrUpdateModel) *model.InvoiceResponse); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.InvoiceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.InvoiceCreateOrUpdateModel) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteInvoice provides a mock function with given fields: ctx, id
func (_m *InvoiceService) DeleteInvoice(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllInvoice provides a mock function with given fields: ctx, filter
func (_m *InvoiceService) GetAllInvoice(ctx context.Context, filter *model.InvoiceFilter) ([]model.InvoiceResponse, model.Meta, error) {
	ret := _m.Called(ctx, filter)

	var r0 []model.InvoiceResponse
	var r1 model.Meta
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.InvoiceFilter) ([]model.InvoiceResponse, model.Meta, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.InvoiceFilter) []model.InvoiceResponse); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.InvoiceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.InvoiceFilter) model.Meta); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(model.Meta)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *model.InvoiceFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetDetailInvoice provides a mock function with given fields: ctx, id
func (_m *InvoiceService) GetDetailInvoice(ctx context.Context, id string) (model.InvoiceResponse, error) {
	ret := _m.Called(ctx, id)

	var r0 model.InvoiceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.InvoiceResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.InvoiceResponse); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.InvoiceResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateInvoice provides a mock function with given fields: ctx, data
func (_m *InvoiceService) UpdateInvoice(ctx context.Context, data model.InvoiceCreateOrUpdateModel) (*entity.Invoice, error) {
	ret := _m.Called(ctx, data)

	var r0 *entity.Invoice
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.InvoiceCreateOrUpdateModel) (*entity.Invoice, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.InvoiceCreateOrUpdateModel) *entity.Invoice); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Invoice)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.InvoiceCreateOrUpdateModel) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewInvoiceService interface {
	mock.TestingT
	Cleanup(func())
}

// NewInvoiceService creates a new instance of InvoiceService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInvoiceService(t mockConstructorTestingTNewInvoiceService) *InvoiceService {
	mock := &InvoiceService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
