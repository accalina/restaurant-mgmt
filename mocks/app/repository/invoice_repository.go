// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/accalina/restaurant-mgmt/app/entity"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	model "github.com/accalina/restaurant-mgmt/app/model"
)

// InvoiceRepository is an autogenerated mock type for the InvoiceRepository type
type InvoiceRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, filter
func (_m *InvoiceRepository) Count(ctx context.Context, filter *model.InvoiceFilter) int {
	ret := _m.Called(ctx, filter)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *model.InvoiceFilter) int); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// FetchAll provides a mock function with given fields: ctx, filter
func (_m *InvoiceRepository) FetchAll(ctx context.Context, filter *model.InvoiceFilter) ([]entity.Invoice, error) {
	ret := _m.Called(ctx, filter)

	var r0 []entity.Invoice
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.InvoiceFilter) ([]entity.Invoice, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.InvoiceFilter) []entity.Invoice); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Invoice)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.InvoiceFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, filter
func (_m *InvoiceRepository) Find(ctx context.Context, filter *model.InvoiceFilter) (entity.Invoice, error) {
	ret := _m.Called(ctx, filter)

	var r0 entity.Invoice
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.InvoiceFilter) (entity.Invoice, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.InvoiceFilter) entity.Invoice); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(entity.Invoice)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.InvoiceFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: tx, data
func (_m *InvoiceRepository) Save(tx *gorm.DB, data *entity.Invoice) (*entity.Invoice, error) {
	ret := _m.Called(tx, data)

	var r0 *entity.Invoice
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *entity.Invoice) (*entity.Invoice, error)); ok {
		return rf(tx, data)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *entity.Invoice) *entity.Invoice); ok {
		r0 = rf(tx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Invoice)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *entity.Invoice) error); ok {
		r1 = rf(tx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewInvoiceRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewInvoiceRepository creates a new instance of InvoiceRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInvoiceRepository(t mockConstructorTestingTNewInvoiceRepository) *InvoiceRepository {
	mock := &InvoiceRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
