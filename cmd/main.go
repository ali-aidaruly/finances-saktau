package main

import (
	"fmt"

	"github.com/ali-aidaruly/finances-saktau/internal/config"
	"github.com/ali-aidaruly/finances-saktau/internal/telegram"
)

func main() {
	var cfg config.Config
	if err := config.ParseConfig(&cfg); err != nil {
		panic(err)
	}

	fmt.Println(int(cfg.TelegramBot.GetUpdateTimeout))

	return

	fmt.Println()

	bot, err := telegram.NewBot(cfg.TelegramBot)
	if err != nil {
		panic(err)
	}

	fmt.Println("bot is running...")
	bot.Run()
}
