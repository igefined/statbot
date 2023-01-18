package telegram

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
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

	SaveQue struct {
		PreviousMessageId int
		Step              int8
		Message           string
		model.DepositSave
	}

	Config struct {
		ApiKey string `config:"TELEGRAM_API_KEY,required"`
	}
)

var saveQue = SaveQue{
	PreviousMessageId: 0,
	Step:              0,
}

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
			c.textMessageHandler(ctx, update)
			return
		}
	}
}

func (c *Client) commandHandler(ctx context.Context, update tgbotapi.Update) {
	switch update.Message.Text {
	case "/info":
		report, err := c.service.Report(ctx)
		if err != nil {
			log.Printf("error get info: %v", err)

			return
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, ToMessage(report))
		msg.ParseMode = parseMode

		c.api.Send(msg)
	case "/save":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Enter a symbol:")
		saveQue.PreviousMessageId = update.Message.MessageID
		saveQue.Step += 1
		c.api.Send(msg)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "This bot helps @lali_lollipup keep track of the investment wallet")
		c.api.Send(msg)
	}
}

func (c *Client) textMessageHandler(ctx context.Context, update tgbotapi.Update) {
	messageString := "bad request"
	if update.Message.MessageID-2 == saveQue.PreviousMessageId {
		c.saveQueHandler(ctx, &saveQue, update)
		messageString = saveQue.Message
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageString)
	c.api.Send(msg)
}

func (c *Client) saveQueHandler(ctx context.Context, que *SaveQue, update tgbotapi.Update) {
	switch que.Step {
	case 1:
		que.Step += 1
		que.Symbol = update.Message.Text
		que.PreviousMessageId = update.Message.MessageID
		que.Message = fmt.Sprintf("Enter count of %s", que.Symbol)
		return
	case 2:
		que.Step += 1
		count, _ := strconv.ParseFloat(update.Message.Text, 64)
		que.Count = count
		que.PreviousMessageId = update.Message.MessageID
		que.Message = fmt.Sprintf("Enter price of %s", que.Symbol)
		return
	case 3:
		que.Step += 1
		price, _ := strconv.ParseFloat(update.Message.Text, 64)
		que.PreviousMessageId = update.Message.MessageID
		que.PurchasePrice = price
		que.Message = fmt.Sprintf("%s saved!", que.Symbol)

		err := c.service.Save(ctx, &que.DepositSave)
		if err != nil {
			if errors.Is(err, model.ErrCoinNotFound) {
				log.Printf("error save currency token %v", err)
				que.Message = "currency is not supported"
			}

			que.Message = "internal bot error"
		}

		que.clean()

		return
	}
}

func ToMessage(c model.Capital) string {
	message := ""

	for k, v := range c {
		message += fmt.Sprintf("<strong>%s</strong>\ncount: %f\npurchase price: %f\nactual price: %f\ndifference in percent: %f%%\n\n", k, v.Count, v.PurchasePrice, v.ActualPrice, v.Difference)
	}

	return message
}

func (q *SaveQue) clean() {
	q.Step = 0
	q.PreviousMessageId = 0
}
