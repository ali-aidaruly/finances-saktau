package composer

import (
	"context"
	"fmt"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
)

func (c *composer) CreateSubscription(ctx context.Context, create models.CreateSubscription) error {
	_, err := c.subsMan.Create(ctx, create)
	if err != nil {
		return fmt.Errorf("error during creating subscription: %w", err)
	}

	return nil
}
