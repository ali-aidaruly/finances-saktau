package subscription

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/pkg/errs"
	"github.com/ali-aidaruly/finances-saktau/pkg/fmap"

	sq "github.com/Masterminds/squirrel"
	"github.com/ali-aidaruly/finances-saktau/internal/models"
)

func (r *repo) Create(ctx context.Context, create models.CreateSubscription) (int, error) {

	insert := map[string]any{
		"user_telegram_id": create.UserTelegramId,
		"name":             create.Name,
		"amount":           sq.Expr("?::NUMERIC", create.Amount),
		"currency":         create.Currency,
		"payment_interval": create.PaymentInterval,
	}

	if create.Description != nil {
		insert["description"] = *create.Description
	}

	cols, vals := fmap.KeyValues(insert)

	query, args, err := sq.Insert(subscriptionTableName).
		Columns(cols...).
		Values(vals...).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return 0, err
	}

	var id int
	if err = r.db.GetContext(ctx, &id, query, args...); err != nil {
		return 0, errs.FromPostgresError(err)
	}

	return id, nil
}
