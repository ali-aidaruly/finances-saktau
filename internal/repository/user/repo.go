package user

import (
	"github.com/ali-aidaruly/finances-saktau/internal/repository/db"
)

const (
	userTableName = "user"
)

type repo struct {
	db db.ExecUnsafe
}

func NewRepo(db db.ExecUnsafe) *repo {
	return &repo{db: db}
}
