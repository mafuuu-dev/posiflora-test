package pg_order

import (
	"backend/core/pkg/repository"
	"backend/internal/domain/order/entity"
	"backend/internal/infrastructure/pg/order/query"
)

type OrderRepository struct {
	*repository.Repository
}

func NewRepository(factory *repository.Factory) *OrderRepository {
	return &OrderRepository{
		Repository: factory.Instance(),
	}
}

func (r *OrderRepository) CreateOrder(
	shopID int64,
	number string,
	total int64,
	customerName string,
) (*order_entity.Order, error) {
	return pg_order_query.NewCreateOrder(r.Query()).
		Set(shopID, number, total, customerName).
		Execute()
}
