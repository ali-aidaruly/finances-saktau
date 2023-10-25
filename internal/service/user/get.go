package user

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
)

func (s *service) GetByTelegramId(ctx context.Context, telegramId int) (models.User, error) {
	user, err := s.repo.GetByTelegramId(ctx, telegramId)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
