package invoice

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/pkg/errs"
	"github.com/ali-aidaruly/finances-saktau/pkg/fmap"

	sq "github.com/Masterminds/squirrel"
	"github.com/ali-aidaruly/finances-saktau/internal/models"
)

func (r *repo) Create(ctx context.Context, invoice models.CreateInvoice) (int, error) {

	insert := map[string]any{
		"user_telegram_id": invoice.UserTelegramId,
		"category_id":      invoice.CategoryId,
		"amount":           sq.Expr("?::NUMERIC", invoice.Amount),
		"currency":         invoice.Currency,
	}

	if invoice.Description != nil {
		insert["description"] = *invoice.Description
	}

	cols, vals := fmap.KeyValues(insert)

	query, args, err := sq.Insert("invoice").
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
