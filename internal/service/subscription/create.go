package subscription

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/pkg/errors"
)

func (s *service) Create(ctx context.Context, create models.CreateSubscription) (int, error) {
	id, err := s.repo.Create(ctx, create)
	if err != nil {
		return 0, errors.Wrap(err, "error during creating subscription")
	}

	return id, nil
}
