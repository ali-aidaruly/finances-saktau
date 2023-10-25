package invoice

import (
	"github.com/ali-aidaruly/finances-saktau/internal/repository/db"
)

const (
	invoiceTableName = "invoice"
)

type repo struct {
	db db.DBExecUnsafe
}

func NewRepo(db db.DBExecUnsafe) *repo {
	return &repo{db: db}
}
