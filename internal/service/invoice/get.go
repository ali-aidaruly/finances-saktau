package invoice

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/internal/models/filters"
)

func (s *service) Get(ctx context.Context, filter filters.InvoiceQuery) ([]models.Invoice, error) {
	invoices, err := s.Get(ctx, filter)
	if err != nil {
		return nil, err
	}

	return invoices, nil
}
