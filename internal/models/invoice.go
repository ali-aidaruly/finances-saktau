package models

import "time"

type Invoice struct {
	Id          int      `db:"id""`
	UserId      int      `db:"user_telegram_id"`
	CategoryId  int      `db:"category_id"`
	Amount      string   `db:"amount"`
	Currency    Currency `db:"currency"`
	Description *string  `db:"description"`

	CreatedAt time.Time `db:"created_at"`
}

type CreateInvoice struct {
	UserTelegramId int
	CategoryId     int
	Amount         string
	Currency       Currency
	Description    *string
}
