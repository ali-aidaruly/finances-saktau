package server

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/internal/pkg/pointer"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
)

func (s *Server) createUser(ctx context.Context, msg tgbotapi.Message) string {
	logger := zerolog.Ctx(ctx).With().Str("command", "createUser").Logger()
	logger.Debug().Msg("creating user")

	var lastName *string
	if msg.From.LastName != "" {
		lastName = pointer.Of(msg.From.LastName)
	}

	user := &models.User{
		TelegramId:       int(msg.Chat.ID),
		TelegramUsername: msg.From.UserName,
		FirstName:        msg.From.FirstName,
		LastName:         lastName,
	}

	err := s.composer.CreateUser(ctx, user)
	if err != nil {
		logger.Error().Err(err).Send()
		return "you are already registered!:)" // TODO: should handle error and return proper text
	}

	return "you are welcome!"
}
