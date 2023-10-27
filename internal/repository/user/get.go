package user

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/pkg/errs"

	sq "github.com/Masterminds/squirrel"
	"github.com/ali-aidaruly/finances-saktau/internal/models"
)

func (r *repo) GetByTelegramId(ctx context.Context, telegramId int) (models.User, error) {

	q := sq.Select("*").
		From(userTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"telegram_id": telegramId})

	queryStr, args, err := q.ToSql()
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	if err = r.db.Unsafe(ctx).GetContext(ctx, &user, queryStr, args...); err != nil {
		return models.User{}, errs.FromPostgresError(err)
	}

	return user, nil
}
