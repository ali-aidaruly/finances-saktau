package models

import "time"

type Subscription struct {
	Id              int      `db:"id"`
	UserId          int      `db:"user_telegram_id"`
	Name            string   `db:"name"`
	Amount          string   `db:"amount"`
	Currency        Currency `db:"currency"`
	PaymentInterval string   `db:"payment_interval"`
	Description     *string  `db:"description"`

	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type CreateSubscription struct {
	UserTelegramId  int
	Name            string
	Amount          string
	Currency        Currency
	PaymentInterval string
	Description     *string
}
