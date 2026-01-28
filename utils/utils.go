package utils

import (
	"errors"
	"slices"
	"strings"
	"unicode"
)

func ParseCurrency(msg string) (string, error) {
	var currencyFrom string
	var value []string

	for _, v := range msg {
		if unicode.IsLetter(v) || unicode.IsSymbol(v) {
			value = append(value, string(v))
		}
	}

	valueStr := strings.ToLower(strings.Join(value, ""))

	switch {
	case slices.Contains(UsdAliases, valueStr):
		currencyFrom = "USD"
	case slices.Contains(RubAliases, valueStr):
		currencyFrom = "RUB"
	case slices.Contains(AedAliases, valueStr):
		currencyFrom = "AED"
	case slices.Contains(TryAliases, valueStr):
		currencyFrom = "TRY"
	case slices.Contains(EurAliases, valueStr):
		currencyFrom = "EUR"
	default:
		return currencyFrom, errors.New("Вы отправили неизвестную для меня валюту.\nПоддерживаемые валюты: RUB, USD, TRY, AED, EUR.")
	}

	return currencyFrom, nil
}

func SelectCurrencies(base string) string {
	var queryCurrencies []string

	switch base {
	case "USD":
		queryCurrencies = append(queryCurrencies, "RUB", "AED", "EUR", "TRY")
	case "RUB":
		queryCurrencies = append(queryCurrencies, "USD", "AED", "EUR", "TRY")
	case "TRY":
		queryCurrencies = append(queryCurrencies, "USD", "AED", "EUR", "RUB")
	case "EUR":
		queryCurrencies = append(queryCurrencies, "USD", "AED", "RUB", "TRY")
	case "AED":
		queryCurrencies = append(queryCurrencies, "USD", "TRY", "EUR", "RUB")
	}

	return strings.Join(queryCurrencies, ",")
}
