package repository

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/internal/models/filters"
	"github.com/ali-aidaruly/finances-saktau/internal/repository/category"
	"github.com/ali-aidaruly/finances-saktau/internal/repository/db"
	"github.com/ali-aidaruly/finances-saktau/internal/repository/invoice"
	"github.com/ali-aidaruly/finances-saktau/internal/repository/user"
)

type InvoiceRepo interface {
	Create(ctx context.Context, invoice models.CreateInvoice) (int, error)

	Get(ctx context.Context, filter filters.InvoiceQuery) ([]models.Invoice, error)

	AmountSumGroupByCategory(ctx context.Context, filter filters.InvoiceSumQuery) ([]models.InvoiceSumByCategory, error)
}

type CategoryRepo interface {
	Create(ctx context.Context, category *models.Category) error

	GetByName(ctx context.Context, name string) (models.Category, error)
	GetAll(ctx context.Context, userTelegramID int) ([]*models.Category, error)
	Exists(ctx context.Context, filter filters.CategoryFilter) (bool, error)
}

type UserRepo interface {
	Create(ctx context.Context, user *models.User) error

	GetByTelegramId(ctx context.Context, telegramId int) (models.User, error)
}

type Repository struct {
	UserRepo
	CategoryRepo
	InvoiceRepo
}

func New(db db.ExecUnsafe) Repository {
	userRepo := user.NewRepo(db)
	categoryRepo := category.NewRepo(db)
	invoiceRepo := invoice.NewRepo(db)

	return Repository{
		UserRepo:     userRepo,
		CategoryRepo: categoryRepo,
		InvoiceRepo:  invoiceRepo,
	}
}
