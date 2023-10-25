package category

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/pkg/errors"
)

func (s *service) Create(ctx context.Context, category *models.Category) error {
	err := s.repo.Create(ctx, category)
	if err != nil {
		return errors.Wrap(err, "error during creating category")
	}

	return nil
}
