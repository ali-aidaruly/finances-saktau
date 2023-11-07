package subscription

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/pkg/errs"
	"github.com/pkg/errors"
)

func (r *repo) GetAll(ctx context.Context, userTelegramID int) ([]models.Subscription, error) {

	q := sq.Select("*").
		From(subscriptionTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"user_telegram_id": userTelegramID})

	queryStr, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	fmt.Print(queryStr)

	var selected = make([]models.Subscription, 0)
	if err = r.db.Unsafe(ctx).SelectContext(ctx, &selected, queryStr, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errs.FromPostgresError(err)
	}

	return selected, nil
}
