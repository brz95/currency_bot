package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/brz95/currency_bot/handler"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler.HandleCurrencyConvert),
	}

	b, err := bot.New(os.Getenv("BOT_TOKEN"), opts...)
	if err != nil {
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handler.StartHandler)
	log.Println("Bot started!")
	b.Start(ctx)
}
