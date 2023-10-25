package db

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type ExecUnsafe interface {
	Exec
	Unsafe(ctx context.Context) Exec
}

type Exec interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
}

type DB struct {
	db *sqlx.DB
}

func NewDB(sqlxDB *sqlx.DB) *DB {
	return &DB{db: sqlxDB}
}

func (d *DB) conn(ctx context.Context) Exec {
	tx, ok := GetTx(ctx)
	if ok {
		return tx
	}
	return d.db
}

func (d *DB) GetSqlxDB() *sqlx.DB {
	return d.db
}

func (d *DB) Unsafe(ctx context.Context) Exec {
	tx, ok := GetTx(ctx)
	if ok {
		return tx.Unsafe()
	}
	return d.db.Unsafe()
}

func (d *DB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return d.conn(ctx).ExecContext(ctx, query, args...)
}

func (d *DB) PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error) {
	return d.conn(ctx).PrepareNamedContext(ctx, query)
}

func (d *DB) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return d.conn(ctx).NamedExecContext(ctx, query, arg)
}

func (d *DB) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.conn(ctx).SelectContext(ctx, dest, query, args...)
}

func (d *DB) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.conn(ctx).GetContext(ctx, dest, query, args...)
}

func (d *DB) PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error) {
	return d.conn(ctx).PreparexContext(ctx, query)
}

func (d *DB) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return d.conn(ctx).QueryxContext(ctx, query, args...)
}

func (d *DB) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return d.conn(ctx).QueryRowxContext(ctx, query, args)
}

func (d *DB) MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result {
	return d.conn(ctx).MustExecContext(ctx, query, args...)
}

func (d *DB) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	return d.db.BindNamed(query, arg)
}

func (d *DB) DoTx(ctx context.Context, txFunc func(context.Context) error) error {
	tx, err := d.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if panicErr := recover(); panicErr != nil {
			_ = tx.Rollback()
			panic(panicErr)
		}
	}()

	ctx = context.WithValue(ctx, ctxTxKey{}, tx)

	if err = txFunc(ctx); err != nil {
		_ = tx.Rollback() // FIXME: handle both errors
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
