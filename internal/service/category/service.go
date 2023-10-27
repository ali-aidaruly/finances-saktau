package category

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/internal/models/filters"
	"github.com/ali-aidaruly/finances-saktau/internal/repository"
)

type Manager interface {
	Create(ctx context.Context, category *models.Category) error

	GetByName(ctx context.Context, name string) (models.Category, error)
	GetAll(ctx context.Context, userTelegramID int) ([]*models.Category, error)
	Exists(ctx context.Context, filter filters.CategoryFilter) (bool, error)
}

type service struct {
	repo repository.CategoryRepo
}

var _ Manager = (*service)(nil)

func NewCategoryService(repo repository.CategoryRepo) *service {
	return &service{
		repo: repo,
	}
}
