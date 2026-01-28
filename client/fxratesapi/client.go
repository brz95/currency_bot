package fxratesapi

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/brz95/currency_bot/utils"
)

type FxApiRatesResp struct {
	Rates FxApiRates `json:"rates"`
}

type FxApiRates struct {
	Rub float32 `json:"RUB"`
	Usd float32 `json:"USD"`
	Aed float32 `json:"AED"`
	Try float32 `json:"TRY"`
	Eur float32 `json:"EUR"`
}

func GetCurrencyRates(base string) (FxApiRatesResp, error) {
	u := &url.URL{
		Scheme: "https",
		Host:   "api.fxratesapi.com",
		Path:   "/latest",
	}

	q := u.Query()
	q.Set("base", base)
	currencies := utils.SelectCurrencies(base)
	q.Set("currencies", currencies)
	q.Set("api_key", os.Getenv("FX_ACCESS_TOKEN"))
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		log.Println("err fxratesapi: ", err)
		return FxApiRatesResp{}, errors.New("я не смог получить курсы валют")
	}

	var fxApiRates FxApiRatesResp
	err = json.NewDecoder(resp.Body).Decode(&fxApiRates)
	if err != nil {
		log.Println("err fxRatesResp to JSON: ", err)
		return FxApiRatesResp{}, errors.New("я не смог обработать курсы валют")
	}

	return fxApiRates, nil
}
