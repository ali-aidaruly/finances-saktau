package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type TX interface {
	DoTx(ctx context.Context, txFunc func(context.Context) error) error
}

type ctxTxKey struct{}

func GetTx(ctx context.Context) (*sqlx.Tx, bool) {
	tx, ok := ctx.Value(ctxTxKey{}).(*sqlx.Tx)
	if !ok {
		return nil, false
	}
	if tx == nil {
		return nil, false
	}

	return tx, true
}
