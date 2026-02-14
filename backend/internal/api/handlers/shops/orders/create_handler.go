package handler_orders

import (
	"backend/core/pkg/errorsx"
	"backend/core/pkg/request"
	"backend/core/pkg/response"
	"backend/core/pkg/scope"
	"backend/internal/domain/order/entity"
	"backend/internal/services/request"
	"backend/internal/usecases/order"

	"github.com/gofiber/fiber/v2"
)

type createRequest struct {
	Number       string `json:"number" validate:"required"`
	Total        int64  `json:"total" validate:"required,gt=0"`
	CustomerName string `json:"customer_name" validate:"required"`
}

type CreateHandler struct {
	*request.Handler
}

func Create(scope *scope.Scope) []fiber.Handler {
	handler := &CreateHandler{
		Handler: request.NewHandler(scope),
	}

	return handler.Instance(handler)
}

func (h *CreateHandler) Handle() fiber.Handler {
	return func(c *fiber.Ctx) error {
		shopID, err := service_request.GetShopIDFromParams(c)
		if err != nil {
			return response.NotValidRequest(c, err)
		}

		model := &createRequest{}
		if err := h.Validator().Validate(c, model); err != nil {
			return response.NotValidRequest(c, err)
		}

		createOrder := app_order.NewCreateOrder(h.SC())
		order, status, err := createOrder.Create(*shopID, model.Number, model.Total, model.CustomerName)
		if err != nil {
			return response.NotValidRequest(c, errorsx.Wrapf(err, "Error creating order:"))
		}

		return response.Success(c, h.response(*order, status))
	}
}

func (h *CreateHandler) response(order order_entity.Order, status string) any {
	return fiber.Map{
		"order": fiber.Map{
			"id":            order.ID,
			"shop_id":       order.ShopID,
			"number":        order.Number,
			"total":         order.Total,
			"customer_name": order.CustomerName,
			"created_at":    order.CreatedAt,
		},
		"status": status,
	}
}
