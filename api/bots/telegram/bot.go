package telegram

import (
	"context"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/igilgyrg/statbot/internal/currency"
	"github.com/igilgyrg/statbot/model"
)

const parseMode = "html"

type (
	Client struct {
		api     *tgbotapi.BotAPI
		service currency.Service
	}

	Config struct {
		ApiKey string `config:"TELEGRAM_API_KEY,required"`
	}
)

func NewBot(service currency.Service) *Client {
	cfg := Config{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := confita.NewLoader(env.NewBackend()).Load(ctx, &cfg); err != nil {
		log.Fatalf("sendsay: failed to load config: %s", err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.ApiKey)
	if err != nil {
		log.Fatal(err)
	}

	return &Client{
		api:     bot,
		service: service,
	}
}

func (c *Client) Run(ctx context.Context) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.api.GetUpdatesChan(u)

	for update := range updates {
		c.messageHandler(ctx, update)
	}
}

func (c *Client) messageHandler(ctx context.Context, update tgbotapi.Update) {
	if update.Message != nil {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.From.UserName != "lali_lollipup" && update.Message.From.UserName != "igpoma" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "forbidden")
			c.api.Send(msg)

			return
		}

		switch string(update.Message.Text[0]) {
		case "/":
			c.commandHandler(ctx, update)
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "bad request")
			c.api.Send(msg)

			return
		}
	}
}

func (c *Client) commandHandler(ctx context.Context, update tgbotapi.Update) {
	switch update.Message.Text {
	case "/info":
		report, err := c.service.Report(ctx)
		if err != nil {
			log.Printf("error get info: %w", err)

			return
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, ToMessage(report))
		msg.ParseMode = parseMode

		c.api.Send(msg)
	case "/save":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "SAVE")
		c.api.Send(msg)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "This bot helps @lali_lollipup keep track of the investment wallet")
		c.api.Send(msg)
	}
}

func ToMessage(c model.Capital) string {
	message := ""

	for k, v := range c {
		message += fmt.Sprintf("<strong>%s</strong>\ncount: %f\npurchase price: %f\nactual price: %f\ndifference in percent: %f%%\n\n", k, v.Count, v.PurchasePrice, v.ActualPrice, v.Difference)
	}

	return message
}
