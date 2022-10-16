package fiat

import (
	"context"
	"fmt"

	"github.com/vdrpkv/cointool/internal/domain/entity/currency"
)

type UseCaseRecognizeFiatCurrency interface {
	DoUseCaseRecognizeFiatCurrency(
		ctx context.Context,
		symbol currency.Symbol,
	) (
		bool,
		error,
	)
}

type useCaseRecognizeFiatCurrency struct {
	fiatCurrencyClient FiatCurrencyClient
}

func NewUseCaseRecognizeFiatCurrency(
	fiatCurrencyRecognizer FiatCurrencyClient,
) UseCaseRecognizeFiatCurrency {
	return useCaseRecognizeFiatCurrency{
		fiatCurrencyClient: fiatCurrencyRecognizer,
	}
}

func (u useCaseRecognizeFiatCurrency) DoUseCaseRecognizeFiatCurrency(
	ctx context.Context,
	symbol currency.Symbol,
) (
	bool,
	error,
) {
	fiat, err := u.fiatCurrencyClient.RecognizeFiatCurrency(ctx, symbol)
	if err != nil {
		return false, fmt.Errorf("recognize fiat client: %w", err)
	}

	return fiat, nil
}
