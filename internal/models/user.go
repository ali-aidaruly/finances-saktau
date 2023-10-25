package models

import "time"

type Currency string

var validCurrencies = map[Currency]struct{}{
	"KZT": {},
}

func (c Currency) IsValid() bool {
	_, ok := validCurrencies[c]

	return ok
}

func (c Currency) String() string {
	return string(c)
}

type User struct {
	TelegramId       int      `db:"telegram_id"`
	TelegramUsername string   `db:"telegram_username"`
	FirstName        string   `db:"first_name"`
	LastName         *string  `db:"last_name"`
	DefaultCurrency  Currency `db:"default_currency"`

	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
