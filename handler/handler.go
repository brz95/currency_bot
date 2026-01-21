package handler

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/brz95/currency_bot/client/fxratesapi"
	"github.com/brz95/currency_bot/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HandleCurrencyConvert(ctx context.Context, b *bot.Bot, update *models.Update) {
	msg := update.Message.Text
	currencyFrom, currencyTo, err := utils.ParseCurrency(msg)

	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   err.Error(),
		})
		return
	}

	re := regexp.MustCompile(`\d+(\.\d+)?`)
	valString := re.FindString(msg)
	val, err := strconv.Atoi(valString)

	if err != nil {
		log.Println("err value to int", err)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Я не смог обработать число",
		})
		return
	}

	currency, err := fxratesapi.GetCurrencyRates()
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   err.Error(),
		})
		return
	}

	switch currencyFrom {
	case "RUB":
		res := float32(val) / currency.Rates.Rub
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("%d %s будет %.2f %s", val, currencyFrom, res, currencyTo),
		})
		return
	case "USD":
		res := float32(val) * currency.Rates.Rub
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("%d %s будет %.2f %s", val, currencyFrom, res, currencyTo),
		})
		return

	case "AED":
		res := (float32(val) * currency.Rates.Rub) / currency.Rates.Aed
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("%d %s будет %.2f %s", val, currencyFrom, res, currencyTo),
		})
		return
	case "TRY":
		res := (float32(val) * currency.Rates.Rub) / currency.Rates.Try
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("%d %s будет %.2f %s", val, currencyFrom, res, currencyTo),
		})
		return
	case "EUR":
		res := (float32(val) * currency.Rates.Rub) / currency.Rates.Eur
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("%d %s будет %.2f %s", val, currencyFrom, res, currencyTo),
		})
		return
	}
}

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Привет! Я бот который отправит тебе курсы валют",
	})
}
