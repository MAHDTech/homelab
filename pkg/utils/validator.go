// Package utils stores utility functions for the project.
package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidationError provides a user-friendly error format for validation failures
type ValidationError struct {
	Errors []string
}

// NewValidationError customises the error message
// for the validator.ValidationErrors.
func NewValidationError(
	errors validator.ValidationErrors,
) *ValidationError {
	var msgs []string

	for _, err := range errors {

		switch err.Tag() {

		case "required":
			msgs = append(msgs, fmt.Sprintf("field '%s' is required", err.Field()))

		case "required_if":
			msgs = append(
				msgs,
				fmt.Sprintf("field '%s' is required when '%s'", err.Field(), err.Param()),
			)

		case "oneof":
			msgs = append(
				msgs,
				fmt.Sprintf("field '%s' must be one of [%s]", err.Field(), err.Param()),
			)

		case "pattern":
			msgs = append(
				msgs,
				fmt.Sprintf(
					"field '%s' contains invalid characters (only alphanumeric, underscore and hyphen allowed)",
					err.Field(),
				),
			)

		case "startswith":
			msgs = append(
				msgs,
				fmt.Sprintf("field '%s' must start with '%s'", err.Field(), err.Param()),
			)

		case "max":
			msgs = append(
				msgs,
				fmt.Sprintf("field '%s' exceeds maximum length of %s", err.Field(), err.Param()),
			)

		case "min":
			msgs = append(
				msgs,
				fmt.Sprintf(
					"field '%s' is shorter than minimum length of %s",
					err.Field(),
					err.Param(),
				),
			)

		default:
			msgs = append(
				msgs,
				fmt.Sprintf("field '%s' failed validation: %s", err.Field(), err.Tag()),
			)

		}
	}

	return &ValidationError{Errors: msgs}

}

// Error returns a string representation of the ValidationError
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed:\n- %s", strings.Join(e.Errors, "\n- "))
}
