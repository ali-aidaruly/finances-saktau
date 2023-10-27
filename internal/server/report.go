package server

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ali-aidaruly/finances-saktau/internal/composer"
	"github.com/ali-aidaruly/finances-saktau/internal/models/filters"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *Server) getReport(ctx context.Context, msg tgbotapi.Message) string {
	msgText := strings.TrimSpace(msg.Text)
	if msgText == "" {
		return "date is empty" // TODO: add wizard
	}

	parts := strings.Split(msgText, " ")
	if len(parts) != 1 {
		return "TODO: not one report"
	}

	date := strings.ToLower(parts[0])

	fn, ok := dateParser[date]
	if !ok {
		return "TODO: invalid dateparses"
	}

	from, to := fn()

	payload, err := s.composer.GetReport(ctx, filters.InvoiceSumQuery{
		UserId: int(msg.From.ID), // TODO: handle from optional
		From:   from,
		To:     to,
	})

	if err != nil {
		return "TODO: getreport error in compo" + err.Error()
	}

	res, err := getReportResponseMessage(payload, from, to)
	if err != nil {
		return "TODO: get-report-resp-msg"
	}

	return res
}

func getReportResponseMessage(payload composer.GetReportPayload, from, to time.Time) (string, error) {
	const size int = 10000

	builder := strings.Builder{}

	builder.Grow(size)

	const layout = "2.1.2006"
	totalMsg := fmt.Sprintf("Total: %d KZT (%s - %s)\n", payload.TotalSum, from.Format(layout), to.Format(layout))

	_, err := builder.WriteString(totalMsg)
	if err != nil {
		return "", err
	}

	for i, w := range payload.InvoiceSums {
		row := fmt.Sprintf("%d) %s: %s KZT\n", i+1, w.CategoryName, w.TotalAmount)

		_, err = builder.WriteString(row)
		if err != nil {
			return "", err
		}
	}

	return builder.String(), nil
}
