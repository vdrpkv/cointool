package rate

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/mocks"
)

func Test_RateCommandHandler(t *testing.T) {
	var (
		zeroExchangeRateValue = 0.00001

		stubCtx  = context.Background()
		stubFrom = currency.Symbol("STUB_FROM")
		stubTo   = currency.Symbol("STUB_TO")
		stubErr  = errors.New("stub error")
	)

	type (
		testCaseGiveStubs struct {
			rate currency.ExchangeRate
			err  error
		}

		testCaseGive struct {
			stubs testCaseGiveStubs
		}

		testCaseWant struct {
			rate currency.ExchangeRate
			err  error
		}

		testCase struct {
			name string
			give testCaseGive
			want testCaseWant
		}
	)

	for _, testCase := range []testCase{
		{
			"any error",
			testCaseGive{
				stubs: testCaseGiveStubs{
					err: stubErr,
				},
			},
			testCaseWant{
				err: stubErr,
			},
		},
		{
			"negative exchange rate error",
			testCaseGive{
				stubs: testCaseGiveStubs{
					rate: -1,
				},
			},
			testCaseWant{
				err: currency.ErrExchangeRateNegative,
			},
		},
		{
			"zero exchange rate error",
			testCaseGive{
				stubs: testCaseGiveStubs{
					rate: currency.ExchangeRate(zeroExchangeRateValue / 10),
				},
			},
			testCaseWant{
				err: currency.ErrExchangeRateZero,
			},
		},
		{
			"success",
			testCaseGive{
				stubs: testCaseGiveStubs{
					rate: 0.99,
				},
			},
			testCaseWant{
				rate: 0.99,
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			exchangeRateGetterMock := mocks.NewExchangeRateGetter(t)
			exchangeRateGetterMock.
				On(
					"GetExchangeRate", stubCtx, stubFrom, stubTo,
				).
				Return(
					testCase.give.stubs.rate, testCase.give.stubs.err,
				)

			rateCommandHandler := New(
				exchangeRateGetterMock,
				zeroExchangeRateValue,
			)
			gotRate, gotErr := rateCommandHandler.HandleRateCommand(
				stubCtx, stubFrom, stubTo,
			)

			require.ErrorIs(t, gotErr, testCase.want.err)
			require.Equal(t, testCase.want.rate, gotRate)
		})
	}
}
