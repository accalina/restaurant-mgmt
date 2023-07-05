// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/accalina/restaurant-mgmt/app/entity"
	mock "github.com/stretchr/testify/mock"

	model "github.com/accalina/restaurant-mgmt/app/model"
)

// FoodService is an autogenerated mock type for the FoodService type
type FoodService struct {
	mock.Mock
}

// CreateFood provides a mock function with given fields: ctx, data
func (_m *FoodService) CreateFood(ctx context.Context, data model.FoodCreateOrUpdateModel) (*entity.Food, error) {
	ret := _m.Called(ctx, data)

	var r0 *entity.Food
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.FoodCreateOrUpdateModel) (*entity.Food, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.FoodCreateOrUpdateModel) *entity.Food); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Food)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.FoodCreateOrUpdateModel) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFood provides a mock function with given fields: ctx, id
func (_m *FoodService) DeleteFood(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllFood provides a mock function with given fields: ctx, filter
func (_m *FoodService) GetAllFood(ctx context.Context, filter *model.FoodFilter) ([]model.FoodResponse, model.Meta, error) {
	ret := _m.Called(ctx, filter)

	var r0 []model.FoodResponse
	var r1 model.Meta
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.FoodFilter) ([]model.FoodResponse, model.Meta, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.FoodFilter) []model.FoodResponse); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.FoodResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.FoodFilter) model.Meta); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(model.Meta)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *model.FoodFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetDetailFood provides a mock function with given fields: ctx, id
func (_m *FoodService) GetDetailFood(ctx context.Context, id string) (model.FoodResponse, error) {
	ret := _m.Called(ctx, id)

	var r0 model.FoodResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.FoodResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.FoodResponse); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.FoodResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFood provides a mock function with given fields: ctx, data
func (_m *FoodService) UpdateFood(ctx context.Context, data model.FoodCreateOrUpdateModel) (*entity.Food, error) {
	ret := _m.Called(ctx, data)

	var r0 *entity.Food
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.FoodCreateOrUpdateModel) (*entity.Food, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.FoodCreateOrUpdateModel) *entity.Food); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Food)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.FoodCreateOrUpdateModel) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFoodService interface {
	mock.TestingT
	Cleanup(func())
}

// NewFoodService creates a new instance of FoodService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFoodService(t mockConstructorTestingTNewFoodService) *FoodService {
	mock := &FoodService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
