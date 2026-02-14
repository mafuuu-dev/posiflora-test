package pg_tgi_query

import (
	"backend/core/pkg/pgscan"
	"backend/core/pkg/query"
	"backend/internal/domain/telegram_integration/entity"
	"backend/internal/infrastructure/pg/telegram_integration/mapper"
)

type CreateIntegration struct {
	*query.Query
	shopID    int64
	botToken  string
	chatID    string
	isEnabled bool
}

func NewCreateIntegration(factory *query.Factory) *CreateIntegration {
	return &CreateIntegration{
		Query: factory.Instance(),
	}
}

func (q *CreateIntegration) Set(shopID int64, botToken string, chatID string, isEnabled bool) *CreateIntegration {
	q.shopID = shopID
	q.botToken = botToken
	q.chatID = chatID
	q.isEnabled = isEnabled
	return q
}

func (q *CreateIntegration) Execute() (*tgi_entity.TelegramIntegration, error) {
	row := q.QueryRow(q, q.shopID, q.botToken, q.chatID, q.isEnabled)
	return pgscan.ScanOne[tgi_entity.TelegramIntegration](row, pg_tgi_mapper.ToEntity)
}

func (q *CreateIntegration) Sql() string {
	return `
		INSERT INTO telegram_integrations (shop_id, bot_token, chat_id, is_enabled)
		SELECT $1::BIGINT, $2::TEXT, $3::TEXT, $4::BOOLEAN
		RETURNING id, shop_id, bot_token, chat_id, is_enabled, created_at, updated_at
	`
}
