package filters

import (
	"time"

	"github.com/ali-aidaruly/finances-saktau/internal/models/sort"
)

type InvoiceQuery struct {
	InvoiceFilter
	sort.Sort
}

type InvoiceFilter struct {
	UserIds     []int
	CategoryIds []int
	FromDate    *time.Time
	TillDate    *time.Time
}

type InvoiceSumQuery struct {
	UserId int
	From   time.Time
	To     time.Time
}
