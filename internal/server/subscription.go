package server

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/ali-aidaruly/finances-saktau/internal/composer"
	"github.com/ali-aidaruly/finances-saktau/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
)

func (s *Server) createSubscription(ctx context.Context, msg tgbotapi.Message) string {
	msgText := strings.TrimSpace(msg.Text)
	if msgText == "" {
		return "subscription command is empty, example: /addsub yandex-plus 1499 KZT monthly [description]"
	}

	parts := strings.Split(msgText, " ")
	if len(parts) != 4 && len(parts) != 5 {
		return "subscription command is empty, example: /addsub yandex-plus 1499 KZT monthly [description]"
	}

	logger := zerolog.Ctx(ctx).With().Str("command", "createSubscription").Logger()
	logger.Debug().Msg("creating subscription")

	if _, err := strconv.ParseFloat(parts[1], 64); err != nil {
		return "amount is not float"
	}

	currency := parts[2]

	if len(currency) != 3 {
		return "currency is longer than 3"
	}

	paymentInterval := parts[3]

	if paymentInterval != "monthly" && paymentInterval != "annual" {
		return "payment must be either monthly or annual"
	}

	var description *string

	if len(parts) == 5 {
		description = &parts[4]
	}

	create := models.CreateSubscription{
		UserTelegramId:  int(msg.From.ID),
		Name:            strings.ToLower(parts[0]),
		Amount:          parts[1],
		Currency:        models.Currency(parts[2]),
		PaymentInterval: paymentInterval,
		Description:     description,
	}

	err := s.composer.CreateSubscription(ctx, create)
	if err != nil {
		logger.Error().Err(err).Send()
		return "sorry something went wrong:( contact Author please"
	}

	return fmt.Sprintf("subscription %s successfully added!", strconv.Quote(parts[0])) // TODO: add emojis to return messages
}

func (s *Server) getAllSubscriptions(ctx context.Context, msg tgbotapi.Message) string {
	logger := zerolog.Ctx(ctx).With().Str("command", "getAllSubs").Logger()
	logger.Debug().Msg("getting all subscriptions")

	subsPayload, err := s.composer.GetAllSubscriptions(ctx, int(msg.From.ID))
	if err != nil {
		logger.Error().Err(err).Send()
		return "sorry something went wrong:( contact Author please"
	}
	logger.Debug().Interface("subscriptions", subsPayload).Send()

	res, err := getSubsResponseMessage(subsPayload)

	return res // TODO: add emojis to return messages
}

func getSubsResponseMessage(payload composer.GetSubscriptionsPayload) (string, error) {
	const size int = 10000

	builder := strings.Builder{}

	builder.Grow(size)

	totalMonthlyMsg := fmt.Sprintf("Total monthly: %d\n\n", payload.MonthlyTotal)

	_, err := builder.WriteString(totalMonthlyMsg)
	if err != nil {
		return "", err
	}

	for i, w := range payload.MonthlySubs {
		row := fmt.Sprintf("%d) %s: %s %s\n", i+1, w.Name, w.Amount, w.Currency)

		_, err = builder.WriteString(row)
		if err != nil {
			return "", err
		}
	}

	builder.WriteString("\n")

	totalAnnualMsg := fmt.Sprintf("Total annual: %d\n\n", payload.AnnualTotal)

	_, err = builder.WriteString(totalAnnualMsg)
	if err != nil {
		return "", err
	}

	for i, w := range payload.AnnualSubs {
		row := fmt.Sprintf("%d) %s: %s %s\n", i+1, w.Name, w.Amount, w.Currency)

		_, err = builder.WriteString(row)
		if err != nil {
			return "", err
		}
	}

	return builder.String(), nil
}
