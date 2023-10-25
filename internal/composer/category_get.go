package composer

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/pkg/errors"
)

func (c *composer) GetAllCategories(ctx context.Context, userTelegramID int) ([]*models.Category, error) {
	categories, err := c.categoryMan.GetAll(ctx, userTelegramID)
	if err != nil {
		return nil, errors.Wrap(err, "error during getting all categories")
	}

	return categories, nil
}
