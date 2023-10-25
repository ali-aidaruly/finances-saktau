package server

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
)

func (s *Server) createCategory(ctx context.Context, msg tgbotapi.Message) string {
	msgText := strings.TrimSpace(msg.Text)
	if msgText == "" {
		return "category name is empty! Example: /addcategory продукты" // TODO: add wizard
	}
	if len(strings.Split(msgText, " ")) > 1 {
		return "enter only one category, without any spaces! Example: /addcategory weekend-hangout" // TODO: format error text
	}

	categoryName := strings.ToLower(msgText)

	logger := zerolog.Ctx(ctx).With().Str("command", "createCategory").Logger()
	logger.Debug().Msg("creating category")

	category := &models.Category{
		UserTelegramId:      int(msg.From.ID),
		Category:            categoryName,
		CategoryOriginTyped: msgText,
	}

	err := s.composer.CreateCategory(ctx, category)
	if err != nil {
		logger.Error().Err(err).Send()
		return "sorry something went wrong:( contact Author please"
	}

	return fmt.Sprintf("category %s successfully added!", strconv.Quote(msgText)) // TODO: add emojis to return messages
}

func (s *Server) getAllCategories(ctx context.Context, msg tgbotapi.Message) string {
	logger := zerolog.Ctx(ctx).With().Str("command", "getAllCategories").Logger()
	logger.Debug().Msg("getting all categories")

	categories, err := s.composer.GetAllCategories(ctx, int(msg.From.ID))
	logger.Debug().Interface("categories", categories).Send()
	if err != nil {
		logger.Error().Err(err).Send()
		return "sorry something went wrong:( contact Author please"
	}

	var builder strings.Builder
	builder.Grow(256) // TODO: magic number
	rowNum := 1
	for rowNum <= len(categories) {
		_, err = builder.WriteString(fmt.Sprintf("%d) %s\n", rowNum, categories[rowNum-1].CategoryOriginTyped))
		if err != nil {
			logger.Error().Err(err).Send()
			return "sorry something went wrong:( contact Author please"
		}
		rowNum++
	}

	return builder.String() // TODO: add emojis to return messages
}
