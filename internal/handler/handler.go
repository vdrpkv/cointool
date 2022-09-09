package handler

import (
	"context"
	"fmt"

	"github.com/vdrpkv/cointool/internal/currency"
)

func HandleConvertCurrency(
	ctx context.Context,
	fiatCurrencyRecognizer FiatCurrencyRecognizer,
	exchangeRateGetter ExchangeRateGetter,
	amount currency.Amount,
	from, to currency.Symbol,
) (
	currency.Amount,
	error,
) {
	rate, err := HandleGetExchangeRate(
		ctx,
		fiatCurrencyRecognizer, exchangeRateGetter,
		from, to,
	)

	if err != nil {
		return 0, err
	}

	return rate.Convert(amount), nil
}

func HandleGetExchangeRate(
	ctx context.Context,
	fiatCurrencyRecognizer FiatCurrencyRecognizer,
	exchangeRateGetter ExchangeRateGetter,
	from, to currency.Symbol,
) (
	currency.ExchangeRate,
	error,
) {
	// Check is FROM currency fiat one.
	isFiat, err := fiatCurrencyRecognizer.RecognizeFiatCurrency(
		ctx, from,
	)

	if err != nil {
		return 0, fmt.Errorf("recognize fiat currency: %w", err)
	}

	// Flip symbols because ExchangeRateGetter accepts
	// only cryptocurrency symbols as FROM currency.
	if isFiat {
		from, to = to, from
	}

	rate, err := exchangeRateGetter.GetExchangeRate(ctx, from, to)
	if err != nil {
		return 0, fmt.Errorf("get exchange rate: %w", err)
	}

	// Flip exchange rate if first currency is fiat one.
	if isFiat {
		rate = rate.Flip()
	}

	return rate, nil
}

func HandleRecognizeFiatCurrency(
	ctx context.Context,
	fiatCurrencyRecognizer FiatCurrencyRecognizer,
	symbol currency.Symbol,
) (
	bool,
	error,
) {
	isFiat, err := fiatCurrencyRecognizer.RecognizeFiatCurrency(ctx, symbol)
	if err != nil {
		return false, fmt.Errorf("recognize fiat currency: %w", err)
	}

	return isFiat, nil
}

type (
	// ExchangeRateGetter gets exchange rate for given currency pair.
	// ExchangeRateGetter accepts cryptocurrency symbol as FROM and any symbol as TO.
	// ExchangeRateGetter may not find exchange rate if fiat currency symbol is passed as FROM.
	ExchangeRateGetter interface {
		GetExchangeRate(
			ctx context.Context,
			from, to currency.Symbol,
		) (
			currency.ExchangeRate,
			error,
		)
	}

	// FiatCurrencyRecognizer checks is given currency symbol is fiat one.
	FiatCurrencyRecognizer interface {
		RecognizeFiatCurrency(
			ctx context.Context,
			symbol currency.Symbol,
		) (
			bool,
			error,
		)
	}
)
