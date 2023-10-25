package server

import (
	"context"
	"strconv"
	"strings"

	"github.com/ali-aidaruly/finances-saktau/internal/composer"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *Server) createInvoice(ctx context.Context, msg tgbotapi.Message) string {
	msgText := strings.TrimSpace(msg.Text)
	if msgText == "" {
		return "invoice is empty! Example: /addinvoice parking 100 KZT (currency is optional)" // TODO: add wizard
	}

	parts := strings.Split(msgText, " ")
	if !(2 <= len(parts) && len(parts) <= 3) {
		return "TODO: "
	}

	if _, err := strconv.ParseFloat(parts[1], 64); err != nil {
		return "TODO: "
	}

	var currency *string
	category, amount := parts[0], parts[1]

	if len(parts) == 3 {
		currency = &parts[2]

		if len(*currency) != 3 {
			return "TODO: "
		}
	}

	err := s.composer.CreateInvoice(ctx, composer.CreateInvoice{
		UserTelegramId: int(msg.From.ID), // TODO: handle if msg.From is nil
		Category:       category,
		Amount:         amount,
		Currency:       currency,
	})
	if err != nil {
		return "TODO: "
	}

	return "success!"
}
