package invoice

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/pkg/errors"
)

func (s *service) Create(ctx context.Context, invoice models.CreateInvoice) (int, error) {

	// TODO: should I check if user exists in db?

	id, err := s.repo.Create(ctx, invoice)
	if err != nil {
		return 0, errors.Wrap(err, "error during creating category")
	}

	return id, nil
}
