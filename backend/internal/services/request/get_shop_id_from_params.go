package service_request

import (
	"backend/core/pkg/errorsx"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetShopIDFromParams(c *fiber.Ctx) (*int64, error) {
	shopID, err := strconv.ParseInt(c.Params("shop_id"), 10, 64)

	if err != nil {
		return nil, errorsx.New("Invalid shop id")
	}
	if shopID <= 0 {
		return nil, errorsx.New("Shop id must be greater than 0")
	}

	return &shopID, nil
}
