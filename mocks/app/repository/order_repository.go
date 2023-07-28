// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/accalina/restaurant-mgmt/app/entity"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	model "github.com/accalina/restaurant-mgmt/app/model"
)

// OrderRepository is an autogenerated mock type for the OrderRepository type
type OrderRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, filter
func (_m *OrderRepository) Count(ctx context.Context, filter *model.OrderFilter) int {
	ret := _m.Called(ctx, filter)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *model.OrderFilter) int); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// FetchAll provides a mock function with given fields: ctx, filter
func (_m *OrderRepository) FetchAll(ctx context.Context, filter *model.OrderFilter) ([]entity.Order, error) {
	ret := _m.Called(ctx, filter)

	var r0 []entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.OrderFilter) ([]entity.Order, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.OrderFilter) []entity.Order); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.OrderFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, filter
func (_m *OrderRepository) Find(ctx context.Context, filter *model.OrderFilter) (*entity.Order, error) {
	ret := _m.Called(ctx, filter)

	var r0 *entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.OrderFilter) (*entity.Order, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.OrderFilter) *entity.Order); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.OrderFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: tx, data
func (_m *OrderRepository) Save(tx *gorm.DB, data *entity.Order) (*entity.Order, error) {
	ret := _m.Called(tx, data)

	var r0 *entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *entity.Order) (*entity.Order, error)); ok {
		return rf(tx, data)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *entity.Order) *entity.Order); ok {
		r0 = rf(tx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *entity.Order) error); ok {
		r1 = rf(tx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewOrderRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderRepository creates a new instance of OrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderRepository(t mockConstructorTestingTNewOrderRepository) *OrderRepository {
	mock := &OrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
