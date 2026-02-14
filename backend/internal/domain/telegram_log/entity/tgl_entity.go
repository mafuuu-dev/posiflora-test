package tgl_entity

import "time"

type TelegramLog struct {
	ID        int64      `json:"id" db:"id"`
	ShopID    int64      `json:"shop_id" db:"shop_id"`
	OrderID   int64      `json:"order_id" db:"order_id"`
	Message   string     `json:"message" db:"message"`
	Status    string     `json:"status" db:"status"`
	Error     string     `json:"error" db:"error"`
	SentAt    *time.Time `json:"sent_at" db:"sent_at"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
}
