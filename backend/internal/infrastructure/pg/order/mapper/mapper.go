package pg_order_mapper

import (
	"backend/core/pkg/pgscan"
	"backend/internal/domain/order/entity"
)

func ToEntity(row pgscan.Scannable) (order_entity.Order, error) {
	var order order_entity.Order
	err := row.Scan(
		&order.ID,
		&order.ShopID,
		&order.Number,
		&order.Total,
		&order.CustomerName,
		&order.CreatedAt,
	)

	return order, err
}
