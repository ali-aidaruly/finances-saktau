package composer

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models/filters"
)

// get invoices by date

func (c *composer) GetInvoices(ctx context.Context, filter filters.InvoiceFilter) error {
	//// TODO: should check user-existence ?
	//
	//c.invoiceMan.(ctx, filter)
	//if err != nil {
	//	return errors.New("TODO")
	//}

	c.invoiceMan.Get(ctx, filter)

	return nil
}
