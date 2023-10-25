package invoice

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/internal/models/filters"
	"github.com/ali-aidaruly/finances-saktau/internal/pkg/errs"
	"github.com/ali-aidaruly/finances-saktau/internal/pkg/pointer"
	"github.com/pkg/errors"
)

func (r *repo) Get(ctx context.Context, filter filters.InvoiceFilter) ([]models.Invoice, error) {

	q := sq.Select("*").
		From(invoiceTableName).
		PlaceholderFormat(sq.Dollar)

	q = applyFilter(q, filter)
	queryStr, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	invoices := make([]models.Invoice, 0)
	if err = r.db.Unsafe(ctx).SelectContext(ctx, &invoices, queryStr, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Invoice{}, nil
		}
		return nil, errs.FromPostgresError(err)
	}

	return invoices, nil
}

func applyFilter(q sq.SelectBuilder, f filters.InvoiceFilter) sq.SelectBuilder {

	if len(f.UserIds) > 0 {
		q = q.Where(sq.Eq{"user_telegram_id": f.UserIds})
	}

	if f.FromDate != nil {
		if f.TillDate == nil {
			f.TillDate = pointer.Of(time.Now())
		}

		q = q.Where(sq.Expr("created_at::DATE BETWEEN ? AND ?", *f.FromDate, *f.TillDate))
	}

	return q
}
