// Package globalschema defines the global configuration schema.
package globalschema

import (
	validator "github.com/go-playground/validator/v10"

	utils "homelab/pkg/utils"
)

// ConfigRaw is used for the initial YAML parsing and validation.
type ConfigRaw struct {
	Enabled interface{} `yaml:"enabled" json:"enabled"`
	Debug   interface{} `yaml:"debug"   json:"debug"`
}

// Config is the final struct for validated configuration.
type Config struct {
	Enabled bool `yaml:"enabled" json:"enabled" validate:"required"`
	Debug   bool `yaml:"debug"   json:"debug"   validate:"required" default:"false"`
}

// ValidateGlobal implements custom validation logic for Global
func (v *Config) ValidateGlobal() error {
	validate := validator.New()

	if err := validate.Struct(v); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return utils.NewValidationError(validationErrors)
		}
		return err
	}
	return nil
}
