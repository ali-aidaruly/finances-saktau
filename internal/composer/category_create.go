package composer

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/pkg/errors"
)

func (c *composer) CreateCategory(ctx context.Context, category *models.Category) error {
	err := c.categoryMan.Create(ctx, category)
	if err != nil {
		return errors.Wrap(err, "error during creating category")
	}

	return nil
}
