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
	if update.Message == nil {
		return
	}
	msg := update.Message.Text
	currencyFrom, err := utils.ParseCurrency(msg)

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

	currency, err := fxratesapi.GetCurrencyRates(currencyFrom)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   err.Error(),
		})
		return
	}

	m := map[string]float32{
		"RUB": float32(val) * currency.Rates.Rub,
		"USD": float32(val) * currency.Rates.Usd,
		"EUR": float32(val) * currency.Rates.Eur,
		"TRY": float32(val) * currency.Rates.Try,
		"AED": float32(val) * currency.Rates.Aed,
	}

	var messageToUser string
	for k, v := range m {
		if v != 0 {
			messageToUser += fmt.Sprintf("%s: %.2f\n", k, v)
		}
	}

	switch currencyFrom {
	case "RUB":
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   messageToUser,
		})
		return
	case "USD":
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   messageToUser,
		})
		return

	case "AED":
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   messageToUser,
		})
		return
	case "TRY":
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   messageToUser,
		})
		return
	case "EUR":
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   messageToUser,
		})
		return
	}
}

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	currency, err := fxratesapi.GetCurrencyRates("RUB")
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   err.Error(),
		})
		return
	}

	m := map[string]float32{
		"USD": 1 / currency.Rates.Usd,
		"EUR": 1 / currency.Rates.Eur,
		"TRY": 1 / currency.Rates.Try,
		"AED": 1 / currency.Rates.Aed,
	}

	var messageToUser string
	for k, v := range m {
		messageToUser += fmt.Sprintf("1 %s: %.2f RUB\n", k, v)
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Привет! Я бот который отправит тебе курсы валют\n\n%s", messageToUser),
	})
}
