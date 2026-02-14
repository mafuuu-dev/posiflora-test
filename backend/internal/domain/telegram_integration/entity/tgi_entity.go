package tgi_entity

import "time"

type TelegramIntegration struct {
	ID        int64     `json:"id" db:"id"`
	ShopID    int64     `json:"shop_id" db:"shop_id"`
	BotToken  string    `json:"bot_token" db:"bot_token"`
	ChatID    string    `json:"chat_id" db:"chat_id"`
	IsEnabled bool      `json:"is_enabled" db:"is_enabled"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
