package category

import (
	"context"
	"database/sql"

	"github.com/ali-aidaruly/finances-saktau/pkg/errs"

	sq "github.com/Masterminds/squirrel"
	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/internal/models/filters"
	"github.com/pkg/errors"
)

func (r *repo) GetByName(ctx context.Context, name string) (models.Category, error) {

	q := sq.Select("*").
		From(categoryTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"name": name})

	queryStr, args, err := q.ToSql()
	if err != nil {
		return models.Category{}, err
	}

	var category models.Category
	if err = r.db.Unsafe(ctx).GetContext(ctx, &category, queryStr, args...); err != nil {
		return models.Category{}, errs.FromPostgresError(err)
	}

	return category, nil
}

func (r *repo) GetAll(ctx context.Context, userTelegramID int) ([]*models.Category, error) {
	q := `SELECT id, category, category_origin_typed, created_at, updated_at
			FROM category WHERE user_telegram_id = $1 AND deleted_at IS NULL`

	var selected []*models.Category
	if err := r.db.SelectContext(ctx, &selected, q, userTelegramID); err != nil && err != sql.ErrNoRows {
		return nil, errs.FromPostgresError(err)
	}

	return selected, nil
}

func (r *repo) Exists(ctx context.Context, filter filters.CategoryFilter) (bool, error) {
	q := sq.Select("1").
		From(categoryTableName).
		PlaceholderFormat(sq.Dollar)

	q = applyFilter(q, filter)
	q = q.Prefix("SELECT EXISTS (").Suffix(")")

	query, args, err := q.ToSql()
	if err != nil {
		return false, err
	}

	var exists bool
	if err = r.db.GetContext(ctx, &exists, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, errs.FromPostgresError(err)
	}

	return exists, nil
}

func applyFilter(q sq.SelectBuilder, filter filters.CategoryFilter) sq.SelectBuilder {
	if filter.Name != nil {
		q = q.Where(sq.Eq{"name": *filter.Name})
	}
	if filter.UserTelegramId != nil {
		q = q.Where(sq.Eq{"user_telegram_id": *filter.UserTelegramId})
	}

	return q
}
