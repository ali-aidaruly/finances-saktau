package user

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/pkg/errs"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/pkg/errors"
)

func (r *repo) Create(ctx context.Context, user *models.User) error {
	q := `INSERT INTO "user" (telegram_id, telegram_username, first_name, last_name)
		  VALUES (:telegram_id, :telegram_username, :first_name, :last_name) 
		  RETURNING created_at, updated_at`

	q, args, err := r.db.BindNamed(q, user)
	if err != nil {
		return errors.Wrap(err, "cannot bind variables for query")
	}

	if err := r.db.QueryRowxContext(ctx, q, args...).StructScan(user); err != nil {
		return errs.FromPostgresError(err)
	}

	return nil
}
