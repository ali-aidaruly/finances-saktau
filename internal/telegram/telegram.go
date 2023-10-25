package telegram

import (
	"log"
	"time"

	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	lastUpdateId     int
	getUpdateTimeout time.Duration
	bot              *tgbotapi.BotAPI
}

type responseChannel chan Response
type requestChannel chan Request

func (resp responseChannel) Respond(response Response) {
	resp <- response
}

func (req requestChannel) GetRequest() Request {
	return <-req
}

type Request struct {
	tgbotapi.Update
}

type Response struct {
	ChatID int
	Text   string
}

func NewBot(config Config) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		return nil, err
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &TelegramBot{
		getUpdateTimeout: config.GetUpdateTimeout,
		bot:              bot,
	}, nil
}

func (t *TelegramBot) Run() (responseChannel, requestChannel) {
	rspCh := responseChannel(make(chan Response))
	reqCh := requestChannel(make(chan Request))

	go t.DoRequests(reqCh)

	go t.ListenResponses(rspCh)

	return rspCh, reqCh
}

func (t *TelegramBot) DoRequests(reqCh requestChannel) {
	updates := t.getUpdates()
	for update := range updates {
		t.lastUpdateId = update.UpdateID + 1

		reqCh <- Request{update}
	}
}

func (t *TelegramBot) ListenResponses(rspCh responseChannel) {
	for r := range rspCh {
		go func(r Response) {
			msg := tgbotapi.NewMessage(int64(r.ChatID), r.Text)

			if _, err := t.bot.Send(msg); err != nil {
				logrus.Error(err)
			}
		}(r)
	}
}

func (t *TelegramBot) getUpdates() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(t.lastUpdateId)
	u.Timeout = int(t.getUpdateTimeout.Seconds())

	return t.bot.GetUpdatesChan(u)
}
