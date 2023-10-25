package composer

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/pkg/errors"
)

func (c *composer) CreateUser(ctx context.Context, user *models.User) error {
	err := c.userMan.Create(ctx, user)
	if err != nil {
		return errors.Wrap(err, "error during creating user") // TODO
	}

	return nil
}
