package models

import "time"

type Invoice struct {
	Id           int      `db:"id"`
	UserId       int      `db:"user_telegram_id"`
	CategoryId   int      `db:"category_id"`
	CategoryName string   `db:"category_name"`
	Amount       string   `db:"amount"`
	Currency     Currency `db:"currency"`
	Description  *string  `db:"description"`

	CreatedAt time.Time `db:"created_at"`
}

type InvoiceSumByCategory struct {
	CategoryName string `db:"category_name"`
	TotalAmount  string `db:"total_amount"`
}

type CreateInvoice struct {
	UserTelegramId int
	CategoryId     int
	Amount         string
	Currency       Currency
	Description    *string
}
