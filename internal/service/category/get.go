package category

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/internal/models/filters"
	"github.com/pkg/errors"
)

func (s *service) GetByName(ctx context.Context, name string) (models.Category, error) {
	category, err := s.repo.GetByName(ctx, name)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *service) GetAll(ctx context.Context, userTelegramID int) ([]*models.Category, error) {
	categories, err := s.repo.GetAll(ctx, userTelegramID)
	if err != nil {
		return nil, errors.Wrap(err, "error during getting all categories")
	}

	return categories, nil
}

func (s *service) Exists(ctx context.Context, filter filters.CategoryFilter) (bool, error) {
	exists, err := s.repo.Exists(ctx, filter)
	if err != nil {
		return false, errors.Wrap(err, "error during getting all categories") // TODO: wrap
	}

	return exists, nil
}
