package subscription

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/internal/repository"
)

type Manager interface {
	Create(ctx context.Context, create models.CreateSubscription) (int, error)

	GetAll(ctx context.Context, userTelegramID int) ([]models.Subscription, error)
}

type service struct {
	repo repository.SubscriptionRepo
}

var _ Manager = (*service)(nil)

func NewService(repo repository.SubscriptionRepo) *service {
	return &service{
		repo: repo,
	}
}
