package pg_tgi_query

import (
	"backend/core/pkg/pgscan"
	"backend/core/pkg/query"
	"backend/internal/domain/telegram_integration/entity"
	"backend/internal/infrastructure/pg/telegram_integration/mapper"
)

type UpdateIntegration struct {
	*query.Query
	shopID    int64
	botToken  string
	chatID    string
	isEnabled bool
}

func NewUpdateIntegration(factory *query.Factory) *UpdateIntegration {
	return &UpdateIntegration{
		Query: factory.Instance(),
	}
}

func (q *UpdateIntegration) Set(shopID int64, botToken string, chatID string, isEnabled bool) *UpdateIntegration {
	q.shopID = shopID
	q.botToken = botToken
	q.chatID = chatID
	q.isEnabled = isEnabled
	return q
}

func (q *UpdateIntegration) Execute() (*tgi_entity.TelegramIntegration, error) {
	row := q.QueryRow(q, q.shopID, q.botToken, q.chatID, q.isEnabled)
	return pgscan.ScanOne[tgi_entity.TelegramIntegration](row, pg_tgi_mapper.ToEntity)
}

func (q *UpdateIntegration) Sql() string {
	return `
		UPDATE telegram_integrations 
		SET bot_token = $2::TEXT, chat_id = $3::TEXT, is_enabled = $4::BOOLEAN 
		WHERE shop_id = $1::BIGINT
		RETURNING id, shop_id, bot_token, chat_id, is_enabled, created_at, updated_at
	`
}
