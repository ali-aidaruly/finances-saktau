package server

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/ali-aidaruly/finances-saktau/pkg/logger"

	"github.com/ali-aidaruly/finances-saktau/internal/composer"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *Server) createInvoice(ctx context.Context, msg tgbotapi.Message) string {
	msgText := strings.TrimSpace(msg.Text)
	if msgText == "" {
		return "invoice is empty! Example: /addinvoice parking 100 KZT (currency is optional)" // TODO: add wizard
	}

	logger.Logger.Debug().Msg("creating invoice")

	parts := strings.Split(msgText, " ")
	if !(2 <= len(parts) && len(parts) <= 3) {
		return "TODO: it is not between 2 and 3"
	}

	if _, err := strconv.ParseFloat(parts[1], 64); err != nil {
		return "TODO: it is not float"
	}

	var currency *string
	category, amount := parts[0], parts[1]

	if len(parts) == 3 {
		currency = &parts[2]

		if len(*currency) != 3 {
			return "TODO: currency is longer than 3"
		}
	}

	err := s.composer.CreateInvoice(ctx, composer.CreateInvoice{
		UserTelegramId: int(msg.From.ID), // TODO: handle if msg.From is nil
		Category:       category,
		Amount:         amount,
		Currency:       currency,
	})
	if err != nil {
		fmt.Println("pzdc")
		logger.Logger.Err(err).Send()
		return err.Error()
	}

	return "success!"
}

func (s *Server) getInvoices(ctx context.Context, msg tgbotapi.Message) string {
	msgText := strings.TrimSpace(msg.Text)
	if msgText == "" {
		msgText = "today" // TODO: add wizard
	}

	parts := strings.Split(msgText, " ")
	if len(parts) != 1 {
		return "TODO: not one"
	}

	date := strings.ToLower(parts[0])

	fn, ok := dateParser[date]
	if !ok {
		return "TODO: not that date"
	}

	from, to := fn()

	payload, err := s.composer.GetInvoices(ctx, composer.GetInvoicesFilter{
		UserId:   int(msg.From.ID), // TODO: handle from is nil
		FromDate: from,
		TillDate: to,
	})

	if err != nil {
		return "TODO: getinvoices err" + err.Error()
	}

	res, err := getInvoicesResponseMessage(payload)
	if err != nil {
		return "TODO: getinvoices resp meesage"
	}

	return res
}

func getInvoicesResponseMessage(payload composer.GetInvoicesPayload) (string, error) {
	const size int = 10000

	builder := strings.Builder{}

	builder.Grow(size)

	totalMsg := fmt.Sprintf("Total: %d KZT\n", payload.Sum)

	_, err := builder.WriteString(totalMsg)
	if err != nil {
		return "", err
	}

	const layout = "15:04"
	for i, w := range payload.Invoices {
		row := fmt.Sprintf("%d) %s: %s KZT %s\n", i+1, w.CategoryName, w.Amount, w.CreatedAt.Local().Format(layout))

		_, err = builder.WriteString(row)
		if err != nil {
			return "", err
		}
	}

	return builder.String(), nil
}
