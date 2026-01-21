package fxratesapi

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
)

type FxApiRatesResp struct {
	Rates FxApiRates `json:"rates"`
}

type FxApiRates struct {
	Rub float32 `json:"RUB"`
	Aed float32 `json:"AED"`
	Try float32 `json:"TRY"`
	Eur float32 `json:"EUR"`
}

func GetCurrencyRates() (FxApiRatesResp, error) {
	u := &url.URL{
		Scheme: "https",
		Host:   "api.fxratesapi.com",
		Path:   "/latest",
	}

	q := u.Query()
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
