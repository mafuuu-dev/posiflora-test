package order_repository

import "backend/internal/domain/order/entity"

type OrderRepository interface {
	CreateOrder(shopID int64, number string, total int64, customerName string) (*order_entity.Order, error)
}
