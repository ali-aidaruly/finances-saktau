package category

import (
	"github.com/ali-aidaruly/finances-saktau/internal/repository/db"
)

const (
	categoryTableName = "category"
)

type repo struct {
	db db.DBExecUnsafe
}

func NewRepo(db db.DBExecUnsafe) *repo {
	return &repo{db: db}
}
