package telegram

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	lastUpdateId     int
	getUpdateTimeout time.Duration
	bot              *tgbotapi.BotAPI
}

func NewBot(config Config) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		return nil, err
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &TelegramBot{
		bot: bot,
	}, nil
}

func (t *TelegramBot) Run() {
	for update := range t.getUpdates() {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			t.bot.Send(msg)
		}
	}
}

func (t *TelegramBot) getUpdates() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(t.lastUpdateId)
	u.Timeout = int(t.getUpdateTimeout)

	return t.bot.GetUpdatesChan(u)
}
