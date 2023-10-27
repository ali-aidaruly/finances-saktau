package composer

import (
	"time"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
)

type CreateInvoice struct {
	UserTelegramId int
	Category       string
	Amount         string
	Currency       *string
}

type ListInvoices struct {
	UserId   int
	Category *string
	FromDate *time.Time
	ToDate   *time.Time
}

type GetInvoicesPayload struct {
	Sum      int
	Invoices []models.Invoice
}

type GetInvoicesFilter struct {
	UserId       int
	CategoryName *string
	FromDate     time.Time
	TillDate     time.Time
}

type GetReportFilter struct {
	UserId   int
	FromDate *time.Time
	TillDate *time.Time
}
