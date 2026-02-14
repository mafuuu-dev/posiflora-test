package pg_tgl_query

import (
	"backend/core/pkg/pgscan"
	"backend/core/pkg/query"
	"backend/internal/domain/telegram_log/dto"
	"backend/internal/infrastructure/pg/telegram_log/mapper"
)

type GetStats struct {
	*query.Query
	shopID int64
}

func NewGetStats(factory *query.Factory) *GetStats {
	return &GetStats{
		Query: factory.Instance(),
	}
}

func (q *GetStats) Set(shopID int64) *GetStats {
	q.shopID = shopID
	return q
}

func (q *GetStats) Execute() (*tgl_dto.StatsDTO, error) {
	row := q.QueryRow(q, q.shopID)
	return pgscan.ScanOne[tgl_dto.StatsDTO](row, pg_tgl_mapper.ToStatsDto)
}

func (q *GetStats) Sql() string {
	return `
		SELECT 
			MAX(sent_at) AS last_sent_at,
			COUNT(*) FILTER (WHERE status = 'SENT') AS sent_count,
			COUNT(*) FILTER (WHERE status = 'FAILED') AS failed_count
		FROM telegram_send_log
		WHERE shop_id = $1 AND sent_at::DATE >= NOW() - INTERVAL '7 days'
	`
}
