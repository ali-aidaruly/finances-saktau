package composer

import (
	"context"
	"math"
	"strconv"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
	"github.com/ali-aidaruly/finances-saktau/internal/models/filters"
	"github.com/ali-aidaruly/finances-saktau/internal/models/sort"
	"github.com/ali-aidaruly/finances-saktau/pkg/slice"
)

// TODO: index on category_name
// TODO: index on created_at

type GetReportPayload struct {
	TotalSum    int
	InvoiceSums []models.InvoiceSumByCategory
}

func (c *composer) GetReport(ctx context.Context, filter filters.InvoiceSumQuery) (GetReportPayload, error) {
	invoices, err := c.invoiceMan.AmountSumGroupByCategory(ctx, filter)
	if err != nil {
		return GetReportPayload{}, err
	}

	sumFloat, err := slice.Reduce(invoices, func(num float64, inv models.InvoiceSumByCategory) (float64, error) {
		temp, err := strconv.ParseFloat(inv.TotalAmount, 64)
		if err != nil {
			return 0, err
		}

		return num + temp, nil
	})
	if err != nil {
		return GetReportPayload{}, err
	}

	sum := int(math.Ceil(sumFloat))

	return GetReportPayload{
		TotalSum:    sum,
		InvoiceSums: invoices,
	}, nil
}

func (c *composer) GetInvoices(ctx context.Context, invoice GetInvoicesFilter) (GetInvoicesPayload, error) {
	//// TODO: should check user-existence ?

	filter := filters.InvoiceQuery{
		InvoiceFilter: filters.InvoiceFilter{
			UserIds:     []int{invoice.UserId},
			CategoryIds: nil,
			FromDate:    &invoice.FromDate,
			TillDate:    &invoice.TillDate,
		},
		Sort: sort.By("created_at"),
	}

	if invoice.CategoryName != nil {
		category, err := c.categoryMan.GetByName(ctx, *invoice.CategoryName)
		if err != nil {
			return GetInvoicesPayload{}, err
		}

		filter.CategoryIds = []int{category.Id}
	}

	invoices, err := c.invoiceMan.Get(ctx, filter)
	if err != nil {
		return GetInvoicesPayload{}, err
	}

	sumFloat, err := slice.Reduce(invoices, func(num float64, inv models.Invoice) (float64, error) {
		temp, err := strconv.ParseFloat(inv.Amount, 64)
		if err != nil {
			return 0, err
		}

		return num + temp, nil
	})
	if err != nil {
		return GetInvoicesPayload{}, err
	}

	sum := int(math.Ceil(sumFloat))

	res := GetInvoicesPayload{
		Sum:      sum,
		Invoices: invoices,
	}

	return res, nil
}
