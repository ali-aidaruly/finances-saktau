package invoice

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/internal/models/filters"
	"github.com/ali-aidaruly/finances-saktau/internal/repository"
)

func NewInvoiceService(repo repository.InvoiceRepo) *service {
	return &service{
		repo: repo,
	}
}

type Manager interface {
	Create(ctx context.Context, invoice models.CreateInvoice) (int, error)

	Get(ctx context.Context, filter filters.InvoiceFilter) ([]models.Invoice, error)
}

type service struct {
	repo repository.InvoiceRepo
}
