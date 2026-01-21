package utils

import (
	"errors"
	"slices"
	"strings"
	"unicode"
)

func ParseCurrency(msg string) (string, string, error) {
	var currencyFrom string
	currencyTo := "RUB"
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
		currencyTo = "USD"
	case slices.Contains(AedAliases, valueStr):
		currencyFrom = "AED"
	case slices.Contains(TryAliases, valueStr):
		currencyFrom = "TRY"
	case slices.Contains(EurAliases, valueStr):
		currencyFrom = "EUR"
	default:
		return currencyFrom, currencyTo, errors.New("Вы отправили неизвестную для меня валюту.\nПоддерживаемые валюты: RUB, USD, TRY, AED, EUR.")
	}

	return currencyFrom, currencyTo, nil
}
