package request

import (
	"backend/core/pkg/errorsx"
	"backend/core/pkg/scope"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ValidationError struct {
	Errors  map[string]string `json:"errors"`
	Code    int               `json:"code"`
	Success bool              `json:"success"`
}

type Validator struct {
	scope *scope.Scope
}

func NewValidator(scope *scope.Scope) *Validator {
	return &Validator{
		scope: scope,
	}
}

func (v *Validator) Validate(c *fiber.Ctx, model any) error {
	if err := v.parse(c, model); err != nil {
		return err
	}

	if err := v.validation(model); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return err
	}

	return nil
}

func (v *Validator) parse(c *fiber.Ctx, model any) error {
	var err error

	if c.Method() == http.MethodGet {
		err = c.QueryParser(model)
	} else {
		err = c.BodyParser(model)
	}

	if err == nil {
		return nil
	}

	c.Status(http.StatusBadRequest)
	return errorsx.Wrap(err, "Error parsing request body")
}

func (v *Validator) validation(model any) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(model)
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		var errs []string
		for _, e := range validationErrors {
			errs = append(errs, fmt.Sprintf(
				"Field '%s' failed on '%s'",
				e.Field(),
				e.Tag(),
			))
		}
		return errors.New(strings.Join(errs, ", "))
	}

	return err
}
