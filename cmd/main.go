package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/ali-aidaruly/finances-saktau/pkg/logger"

	"github.com/ali-aidaruly/finances-saktau/internal/composer"
	"github.com/ali-aidaruly/finances-saktau/internal/repository"
	"github.com/ali-aidaruly/finances-saktau/internal/repository/db"
	"github.com/ali-aidaruly/finances-saktau/internal/server"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/ali-aidaruly/finances-saktau/internal/config"
	"github.com/ali-aidaruly/finances-saktau/internal/telegram"
)

func main() {
	var cfg config.Config
	if err := config.ParseConfig(&cfg); err != nil {
		log.Fatal(err)
	}

	logg := logger.NewLogger(&cfg.Logger)
	logg.Info().Interface("config", cfg).Msg("The gathered config")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(logg.WithContext(ctx))
	defer cancel()

	db, err := connectDB(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	bot, err := telegram.NewBot(cfg.TelegramBot)
	if err != nil {
		panic(err)
	}

	fmt.Println("bot is running...")
	respChan, reqChan := bot.Run()

	repos := repository.New(db)

	composer := composer.NewComposer(repos)

	srv := server.NewServer(composer, reqChan, respChan)
	srv.Run(ctx)
}

func connectDB(cfg config.DbConfig) (*db.DB, error) {
	sqlDb, err := sql.Open(cfg.Driver, cfg.ConnectionString)

	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to database")
	}

	if err := sqlDb.Ping(); err != nil {
		return nil, errors.Wrap(err, "cannot ping database")
	}
	dbx := sqlx.NewDb(sqlDb, cfg.Driver)

	dbx.SetMaxOpenConns(cfg.MaxOpenConnections)
	dbx.SetMaxIdleConns(cfg.MaxIdleConnections)
	dbx.SetConnMaxLifetime(cfg.ConnectionMaxLifetime)

	return db.NewDB(dbx), nil
}
