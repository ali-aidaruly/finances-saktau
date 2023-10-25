package composer

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/models"
)

func (c *composer) CreateInvoice(ctx context.Context, invoice CreateInvoice) error {
	// TODO: should check user-existence ?

	category, err := c.categoryMan.GetByName(ctx, invoice.Category)
	if err != nil {
		return err
	}

	var currency models.Currency
	if invoice.Currency == nil {
		user, err := c.userMan.GetByTelegramId(ctx, invoice.UserTelegramId)
		if err != nil {
			return err
		}

		currency = user.DefaultCurrency
	}

	_, err = c.invoiceMan.Create(ctx, models.CreateInvoice{
		UserTelegramId: invoice.UserTelegramId,
		CategoryId:     category.Id,
		Amount:         invoice.Amount,
		Currency:       currency,
	})
	if err != nil {
		return err
	}

	return nil
}
