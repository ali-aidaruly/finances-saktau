package telegram

import "time"

type Config struct {
	BotToken         string        `env:"BOT_TOKEN,required"`
	GetUpdateTimeout time.Duration `env:"GetUpdateTimeout,required"`
}
