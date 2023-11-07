package composer

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/service/subscription"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/internal/models/filters"
	"github.com/ali-aidaruly/finances-saktau/internal/repository"
	"github.com/ali-aidaruly/finances-saktau/internal/service/category"
	"github.com/ali-aidaruly/finances-saktau/internal/service/invoice"
	"github.com/ali-aidaruly/finances-saktau/internal/service/user"
)

type composer struct {
	invoiceMan  invoice.Manager
	categoryMan category.Manager
	userMan     user.Manager
	subsMan     subscription.Manager
}

type Composer interface {
	CreateInvoice(ctx context.Context, invoice CreateInvoice) error
	GetInvoices(ctx context.Context, invoice GetInvoicesFilter) (GetInvoicesPayload, error)

	CreateUser(ctx context.Context, user *models.User) error
	GetReport(ctx context.Context, filter filters.InvoiceSumQuery) (GetReportPayload, error)

	CreateCategory(ctx context.Context, category *models.Category) error
	GetAllCategories(ctx context.Context, userTelegramID int) ([]*models.Category, error)

	CreateSubscription(ctx context.Context, create models.CreateSubscription) error
	GetAllSubscriptions(ctx context.Context, userTelegramID int) (GetSubscriptionsPayload, error)
}

var _ Composer = (*composer)(nil)

func NewComposer(repos repository.Repository) *composer {

	userService := user.NewUserService(repos.UserRepo)
	categoryService := category.NewCategoryService(repos.CategoryRepo)
	invoiceService := invoice.NewInvoiceService(repos.InvoiceRepo)
	subsService := subscription.NewService(repos.SubscriptionRepo)

	return &composer{
		userMan:     userService,
		categoryMan: categoryService,
		invoiceMan:  invoiceService,
		subsMan:     subsService,
	}
}
