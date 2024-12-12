// Package nutanixschema defines the Nutanix configuration schema.
package nutanixschema

import (
	validator "github.com/go-playground/validator/v10"

	utils "homelab/pkg/utils"
)

// ConfigRaw is used for the initial YAML parsing and validation.
type ConfigRaw struct {
	Enabled interface{} `yaml:"enabled" json:"enabled"`
}

// Config is the final struct for validated configuration.
type Config struct {
	Enabled bool `yaml:"enabled" json:"enabled" validate:"required"`
}

// ValidateNutanix implements custom validation logic for Nutanix
func (v *Config) ValidateNutanix() error {
	validate := validator.New()

	if err := validate.Struct(v); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return utils.NewValidationError(validationErrors)
		}
		return err
	}
	return nil
}
