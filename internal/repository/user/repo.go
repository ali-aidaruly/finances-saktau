package user

import (
	"github.com/ali-aidaruly/finances-saktau/internal/repository/db"
)

const (
	userTableName = "user"
)

type repo struct {
	db db.DBExecUnsafe
}

func NewRepo(db db.DBExecUnsafe) *repo {
	return &repo{db: db}
}
