// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/accalina/restaurant-mgmt/app/entity"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	model "github.com/accalina/restaurant-mgmt/app/model"
)

// TableRepository is an autogenerated mock type for the TableRepository type
type TableRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, filter
func (_m *TableRepository) Count(ctx context.Context, filter *model.TableFilter) int {
	ret := _m.Called(ctx, filter)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *model.TableFilter) int); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// FetchAll provides a mock function with given fields: ctx, filter
func (_m *TableRepository) FetchAll(ctx context.Context, filter *model.TableFilter) ([]entity.Table, error) {
	ret := _m.Called(ctx, filter)

	var r0 []entity.Table
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.TableFilter) ([]entity.Table, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.TableFilter) []entity.Table); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Table)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.TableFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, filter
func (_m *TableRepository) Find(ctx context.Context, filter *model.TableFilter) (*entity.Table, error) {
	ret := _m.Called(ctx, filter)

	var r0 *entity.Table
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.TableFilter) (*entity.Table, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.TableFilter) *entity.Table); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Table)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.TableFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: tx, data
func (_m *TableRepository) Save(tx *gorm.DB, data *entity.Table) (*entity.Table, error) {
	ret := _m.Called(tx, data)

	var r0 *entity.Table
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *entity.Table) (*entity.Table, error)); ok {
		return rf(tx, data)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *entity.Table) *entity.Table); ok {
		r0 = rf(tx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Table)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *entity.Table) error); ok {
		r1 = rf(tx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTableRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTableRepository creates a new instance of TableRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTableRepository(t mockConstructorTestingTNewTableRepository) *TableRepository {
	mock := &TableRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
