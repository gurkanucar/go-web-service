package validator

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Validator struct {
	validate *validator.Validate
}

func New() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

func (v *Validator) Validate(i interface{}) map[string]string {
	if err := v.validate.Struct(i); err != nil {
		errors := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			field := strings.ToLower(e.Field())
			errors[field] = formatError(e)
		}
		return errors
	}
	return nil
}

func ValidatedBody[T any](c *fiber.Ctx, v *Validator) (T, error) {
	var payload T

	if err := c.BodyParser(&payload); err != nil {
		return payload, fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
	}

	if errs := v.Validate(payload); errs != nil {
		fiberErr := fiber.NewError(fiber.StatusBadRequest, "Validation failed")
		errJSON, _ := json.Marshal(errs)
		fiberErr.Message = string(errJSON)
		return payload, fiberErr
	}

	return payload, nil
}

func ValidatedQuery[T any](c *fiber.Ctx, v *Validator) (T, error) {
	var payload T

	if err := c.QueryParser(&payload); err != nil {
		return payload, fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid query parameters: %v", err))
	}

	if errs := v.Validate(payload); errs != nil {
		fiberErr := fiber.NewError(fiber.StatusBadRequest, "Validation failed")
		errJSON, _ := json.Marshal(errs)
		fiberErr.Message = string(errJSON)
		return payload, fiberErr
	}

	return payload, nil
}

func formatError(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Minimum length/value is %s", e.Param())
	case "max":
		return fmt.Sprintf("Maximum length/value is %s", e.Param())
	case "lte":
		return fmt.Sprintf("Must be less than or equal to %s", e.Param())
	case "gte":
		return fmt.Sprintf("Must be greater than or equal to %s", e.Param())
	case "len":
		return fmt.Sprintf("Length must be exactly %s", e.Param())
	case "alphanum":
		return "Must be alphanumeric"
	case "numeric":
		return "Must be numeric"
	case "url":
		return "Must be a valid URL"
	default:
		return "Invalid value"
	}
}
