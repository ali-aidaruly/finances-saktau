package user

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/pkg/errors"
)

func (s *service) Create(ctx context.Context, user *models.User) error {
	err := s.repo.Create(ctx, user)
	if err != nil {
		return errors.Wrap(err, "error during creating user") // TODO
	}

	return nil
}
