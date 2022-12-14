// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/vdrpkv/cointool/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// UseCaseRecognizeFiatCurrency is an autogenerated mock type for the UseCaseRecognizeFiatCurrency type
type UseCaseRecognizeFiatCurrency struct {
	mock.Mock
}

// DoUseCaseRecognizeFiatCurrency provides a mock function with given fields: ctx, symbol
func (_m *UseCaseRecognizeFiatCurrency) DoUseCaseRecognizeFiatCurrency(ctx context.Context, symbol entity.CurrencySymbol) (bool, error) {
	ret := _m.Called(ctx, symbol)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, entity.CurrencySymbol) bool); ok {
		r0 = rf(ctx, symbol)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.CurrencySymbol) error); ok {
		r1 = rf(ctx, symbol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUseCaseRecognizeFiatCurrency interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCaseRecognizeFiatCurrency creates a new instance of UseCaseRecognizeFiatCurrency. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCaseRecognizeFiatCurrency(t mockConstructorTestingTNewUseCaseRecognizeFiatCurrency) *UseCaseRecognizeFiatCurrency {
	mock := &UseCaseRecognizeFiatCurrency{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
