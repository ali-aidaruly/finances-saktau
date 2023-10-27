package invoice

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/internal/models/filters"
)

func (s *service) AmountSumGroupByCategory(ctx context.Context, filter filters.InvoiceSumQuery) ([]models.InvoiceSumByCategory, error) {
	return s.repo.AmountSumGroupByCategory(ctx, filter)
}
