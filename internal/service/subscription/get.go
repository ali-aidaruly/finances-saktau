package subscription

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/pkg/errors"
)

func (s *service) GetAll(ctx context.Context, userTelegramID int) ([]models.Subscription, error) {
	categories, err := s.repo.GetAll(ctx, userTelegramID)
	if err != nil {
		return nil, errors.Wrap(err, "error during getting all subscriptions")
	}

	return categories, nil
}
