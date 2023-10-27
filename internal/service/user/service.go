package user

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/internal/repository"
)

func NewUserService(repo repository.UserRepo) *service {
	return &service{
		repo: repo,
	}
}

type service struct {
	repo repository.UserRepo
}

var _ Manager = (*service)(nil)

type Manager interface {
	Create(ctx context.Context, user *models.User) error

	GetByTelegramId(ctx context.Context, telegramId int) (models.User, error)
}
