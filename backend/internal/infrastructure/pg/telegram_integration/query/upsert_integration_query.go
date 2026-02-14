package pg_tgi_query

import (
	"backend/core/pkg/pgscan"
	"backend/core/pkg/query"
	"backend/internal/domain/telegram_integration/entity"
	"backend/internal/infrastructure/pg/telegram_integration/mapper"
)

type UpsertIntegration struct {
	*query.Query
	shopID    int64
	botToken  string
	chatID    string
	isEnabled bool
}

func NewUpsertIntegration(factory *query.Factory) *UpsertIntegration {
	return &UpsertIntegration{
		Query: factory.Instance(),
	}
}

func (q *UpsertIntegration) Set(shopID int64, botToken string, chatID string, isEnabled bool) *UpsertIntegration {
	q.shopID = shopID
	q.botToken = botToken
	q.chatID = chatID
	q.isEnabled = isEnabled
	return q
}

func (q *UpsertIntegration) Execute() (*tgi_entity.TelegramIntegration, error) {
	row := q.QueryRow(q, q.shopID, q.botToken, q.chatID, q.isEnabled)
	return pgscan.ScanOne[tgi_entity.TelegramIntegration](row, pg_tgi_mapper.ToEntity)
}

func (q *UpsertIntegration) Sql() string {
	return `
		INSERT INTO telegram_integrations (shop_id, bot_token, chat_id, is_enabled)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (shop_id) 
		DO UPDATE SET bot_token = $2::TEXT, chat_id = $3::TEXT, is_enabled = $4::BOOLEAN
		RETURNING id, shop_id, bot_token, chat_id, is_enabled, created_at, updated_at
	`
}
