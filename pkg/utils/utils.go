// Package utils stores utility functions for the project.
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// LogInfo is a helper function that logs a message and
// handles any potential logging errors by totally ignoring them :)
func LogInfo(ctx *pulumi.Context, format string, args ...interface{}) {

	// The user might either pass a format string with variables,
	// or just a string directly.

	var message string

	if len(args) > 0 {
		message = fmt.Sprintf(format, args...)
	} else {
		message = format
	}

	if err := ctx.Log.Info(message, nil); err != nil {
		fmt.Printf("logging error: %v\n", err)
	}

}

// LogWarn is a helper function that logs a warning message and
// handles any potential logging errors by totally ignoring them :)
func LogWarn(ctx *pulumi.Context, format string, args ...interface{}) {

	// The user might either pass a format string with variables,
	// or just a string directly.

	var message string

	if len(args) > 0 {
		message = fmt.Sprintf(format, args...)
	} else {
		message = format
	}

	if err := ctx.Log.Warn(message, nil); err != nil {
		fmt.Printf("logging error: %v\n", err)
	}

}

// LogError is a helper function that logs an error message and
// handles any potential logging errors by totally ignoring them :)
func LogError(ctx *pulumi.Context, format string, args ...interface{}) {

	// The user might either pass a format string with variables,
	// or just a string directly.

	var message string

	if len(args) > 0 {
		message = fmt.Sprintf(format, args...)
	} else {
		message = format
	}

	if err := ctx.Log.Error(message, nil); err != nil {
		fmt.Printf("logging error: %v\n", err)
	}

}

// TryObject loads and validates configuration data into a typed schema.
// It provides detailed error messages for missing fields and type mismatches.
// The key is optional, and if provided, the value to validate will be retrieved.
func TryObject(
	key string,
	input interface{},
	output interface{},
) error {
	// Get the value to validate based on key
	valueToValidate := input
	if key != "" {
		if m, ok := input.(map[string]interface{}); ok {
			var exists bool
			valueToValidate, exists = m[key]
			if !exists {
				return fmt.Errorf("configuration key '%s' not found", key)
			}
		}
	}

	// First pass: Convert to JSON to validate basic structure
	jsonBytes, err := json.Marshal(valueToValidate)
	if err != nil {
		return fmt.Errorf("failed to marshal input: %w", err)
	}

	// Use a decoder for strict validation
	decoder := json.NewDecoder(bytes.NewReader(jsonBytes))
	decoder.DisallowUnknownFields()

	// Second pass: Decode with strict type checking
	if err := decoder.Decode(output); err != nil {
		// Handle different types of JSON decode errors
		switch e := err.(type) {

		case *json.UnmarshalTypeError:
			return fmt.Errorf("type mismatch at field '%s': expected %s but got %s",
				e.Field, e.Type, e.Value)

		case *json.SyntaxError:
			return fmt.Errorf("invalid JSON syntax at position %d: %s",
				e.Offset, e.Error())

		default:
			// Check if it's an unknown field error
			if strings.Contains(err.Error(), "unknown field") {
				return fmt.Errorf("configuration contains unknown field: %s", err.Error())
			}
			return fmt.Errorf("validation error: %w", err)
		}
	}

	// Third pass: Validate using struct tags if the output implements Validator
	if validator, ok := output.(interface{ Validate() error }); ok {
		if err := validator.Validate(); err != nil {
			return fmt.Errorf("schema validation failed: %w", err)
		}
	}

	return nil
}
