package models

import "time"

type Category struct {
	Id                  int    `db:"id"`
	UserTelegramId      int    `db:"user_telegram_id"`
	Category            string `db:"category"`
	CategoryOriginTyped string `db:"category_origin_typed"`

	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
