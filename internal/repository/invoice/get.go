package invoice

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ali-aidaruly/finances-saktau/internal/models/sort"
	"github.com/ali-aidaruly/finances-saktau/pkg/errs"
	"github.com/ali-aidaruly/finances-saktau/pkg/pointer"

	sq "github.com/Masterminds/squirrel"
	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/internal/models/filters"
	"github.com/pkg/errors"
)

func (r *repo) Get(ctx context.Context, filter filters.InvoiceQuery) ([]models.Invoice, error) {

	q := sq.Select(
		"i.id id",
		"i.user_telegram_id user_telegram_id",
		"i.category_id as category_id",
		"i.amount as amount",
		"i.currency as currency",
		"i.description as description",
		"i.created_at as created_at",
		"i.updated_at as updated_at",
		"c.name as category_name",
	).
		From(fmt.Sprintf("invoice i")).
		Join("category c ON i.category_id = c.id").
		PlaceholderFormat(sq.Dollar).Where("i.deleted_at IS NULL")

	q = applyFilter(q, filter.InvoiceFilter)
	q = applySort(q, filter.Sort)

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

func (r *repo) AmountSumGroupByCategory(ctx context.Context, filter filters.InvoiceSumQuery) ([]models.InvoiceSumByCategory, error) {

	q := sq.Select("c.name AS category_name", "SUM(i.amount)::text AS total_amount").
		From("invoice i").
		Join("category c ON i.category_id = c.id").
		Where(sq.Expr("i.created_at BETWEEN ? AND ? AND i.user_telegram_id = ?", filter.From, filter.To, filter.UserId)).
		GroupBy("c.name").
		OrderBy("total_amount DESC").
		PlaceholderFormat(sq.Dollar)

	queryStr, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	invoices := make([]models.InvoiceSumByCategory, 0)
	if err = r.db.SelectContext(ctx, &invoices, queryStr, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.InvoiceSumByCategory{}, nil
		}
		return nil, errs.FromPostgresError(err)
	}

	return invoices, nil
}

func applyFilter(q sq.SelectBuilder, f filters.InvoiceFilter) sq.SelectBuilder {

	if len(f.UserIds) > 0 {
		q = q.Where(sq.Eq{"i.user_telegram_id": f.UserIds})
	}

	if len(f.CategoryIds) > 0 {
		q = q.Where(sq.Eq{"category_id": f.CategoryIds})
	}

	if f.FromDate != nil {
		if f.TillDate == nil {
			f.TillDate = pointer.Of(time.Now())
		}

		q = q.Where(sq.Expr("i.created_at BETWEEN ? AND ?", *f.FromDate, *f.TillDate))
	}

	return q
}

func applySort(q sq.SelectBuilder, sort sort.Sort) sq.SelectBuilder {

	if sort.IsValid() {
		q = q.OrderByClause(fmt.Sprintf("%s %s", sort.Field(), sort.SQLDir()))
	} else {
		q = q.OrderBy("id desc")
	}

	return q
}
