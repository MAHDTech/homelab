// Package utils stores utility functions for the project.
package utils

import (
	"encoding/json"
	"fmt"

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

// TryObject loads an optional configuration value by its key into the output variable,
// or returns an error if unable to do so.
func TryObject(
	input interface{},
	key string,
	output interface{},
) error {
	// If input is a map, try to get the value by key
	if m, ok := input.(map[string]interface{}); ok {
		if value, exists := m[key]; exists {
			return tryObject(value, output)
		}
		return fmt.Errorf("key '%s' not found in configuration", key)
	}

	// If input is not a map, try to unmarshal the entire input
	return tryObject(input, output)
}

// tryObject is a helper function that attempts to load an optional configuration value
// into the output variable, or returns an error if unable to do so.
func tryObject(
	input interface{},
	output interface{},
) error {
	// Convert input to JSON bytes based on its type
	var jsonBytes []byte
	switch v := input.(type) {
	case string:
		jsonBytes = []byte(v)
	case []byte:
		jsonBytes = v
	case map[string]interface{}:
		var err error
		jsonBytes, err = json.Marshal(v)
		if err != nil {
			return fmt.Errorf("failed to marshal map: %w", err)
		}
	default:
		// If it's any other type, marshal it to JSON first
		var err error
		jsonBytes, err = json.Marshal(v)
		if err != nil {
			return fmt.Errorf("failed to marshal input: %w", err)
		}
	}

	if err := json.Unmarshal(jsonBytes, output); err != nil {
		return fmt.Errorf("failed to unmarshal into output: %w", err)
	}

	return nil
}
