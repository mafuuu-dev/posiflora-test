package order_entity

import "time"

type Order struct {
	ID           int64     `json:"id" db:"id"`
	ShopID       int64     `json:"shop_id" db:"shop_id"`
	Number       string    `json:"number" db:"number"`
	Total        int64     `json:"total" db:"total"`
	CustomerName string    `json:"customer_name" db:"customer_name"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
