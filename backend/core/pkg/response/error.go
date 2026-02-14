package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Success bool   `json:"success"`
}

func Error(c *fiber.Ctx, err string, code int) error {
	return c.Status(code).JSON(ErrorResponse{
		Error:   err,
		Code:    code,
		Success: false,
	})
}

func NotValidRequest(c *fiber.Ctx, err error) error {
	return Error(c, err.Error(), http.StatusUnprocessableEntity)
}
