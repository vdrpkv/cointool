// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	currency "github.com/vdrpkv/cointool/internal/currency"
)

// FiatCurrencyRecognizer is an autogenerated mock type for the FiatCurrencyRecognizer type
type FiatCurrencyRecognizer struct {
	mock.Mock
}

// RecognizeFiatCurrency provides a mock function with given fields: ctx, from
func (_m *FiatCurrencyRecognizer) RecognizeFiatCurrency(ctx context.Context, from currency.Symbol) (bool, error) {
	ret := _m.Called(ctx, from)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, currency.Symbol) bool); ok {
		r0 = rf(ctx, from)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, currency.Symbol) error); ok {
		r1 = rf(ctx, from)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFiatCurrencyRecognizer interface {
	mock.TestingT
	Cleanup(func())
}

// NewFiatCurrencyRecognizer creates a new instance of FiatCurrencyRecognizer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFiatCurrencyRecognizer(t mockConstructorTestingTNewFiatCurrencyRecognizer) *FiatCurrencyRecognizer {
	mock := &FiatCurrencyRecognizer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
