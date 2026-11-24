package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("Переменная окружения TELEGRAM_BOT_TOKEN не задана")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := bot.New(token)
	if err != nil {
		log.Fatalf("Не удалось создать Telegram-бота: %v", err)
	}

	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"/start",
		bot.MatchTypeExact,
		startHandler,
	)

	log.Println("Бот запущен и ожидает сообщения...")
	b.Start(ctx)
}

func startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Привет! Я бот для записи на услуги мастера. Скоро здесь появится возможность выбрать удобное время для посещения 💅",
	})
	if err != nil {
		log.Printf("Не удалось отправить приветствие: %v", err)
	}
}
