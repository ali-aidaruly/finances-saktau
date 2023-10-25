package category

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/internal/pkg/errs"
	"github.com/pkg/errors"
)

func (r *repo) Create(ctx context.Context, category *models.Category) error {
	q := `INSERT INTO categories (user_telegram_id, category, category_origin_typed)
		  VALUES (:user_telegram_id, :category, :category_origin_typed) 
		  RETURNING id, created_at, updated_at`

	q, args, err := r.db.BindNamed(q, category)
	if err != nil {
		return errors.Wrap(err, "cannot bind variables for query") // TODO: error wrapping
	}

	if err := r.db.QueryRowxContext(ctx, q, args...).StructScan(category); err != nil {
		return errs.FromPostgresError(err)
	}

	return nil
}
