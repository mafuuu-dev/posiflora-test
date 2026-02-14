package pg_order_query

import (
	"backend/core/pkg/pgscan"
	"backend/core/pkg/query"
	"backend/internal/domain/order/entity"
	"backend/internal/infrastructure/pg/order/mapper"
)

type CreateOrder struct {
	*query.Query
	shopID       int64
	number       string
	total        int64
	customerName string
}

func NewCreateOrder(factory *query.Factory) *CreateOrder {
	return &CreateOrder{
		Query: factory.Instance(),
	}
}

func (q *CreateOrder) Set(shopID int64, number string, total int64, customerName string) *CreateOrder {
	q.shopID = shopID
	q.number = number
	q.total = total
	q.customerName = customerName
	return q
}

func (q *CreateOrder) Execute() (*order_entity.Order, error) {
	row := q.QueryRow(q, q.shopID, q.number, q.total, q.customerName)
	return pgscan.ScanOne[order_entity.Order](row, pg_order_mapper.ToEntity)
}

func (q *CreateOrder) Sql() string {
	return `
		INSERT INTO orders (shop_id, number, total, customer_name)
		SELECT $1::BIGINT, $2::TEXT, $3::BIGINT, $4::TEXT
		RETURNING id, shop_id, number, total, customer_name, created_at
	`
}
