package filters

import "time"

type InvoiceFilter struct {
	UserIds  []int
	FromDate *time.Time
	TillDate *time.Time
}
