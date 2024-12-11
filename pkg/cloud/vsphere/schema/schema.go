// Package vsphereschema contains the structs for the VMware configuration.
package vsphereschema

import (
	"errors"
	"fmt"
	"regexp"

	validator "github.com/go-playground/validator/v10"
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	utils "homelab/pkg/utils"
)

// ConfigRaw is used for the initial YAML parsing and validation.
type ConfigRaw struct {
	Enabled          interface{} `yaml:"enabled"            json:"enabled"`
	Common           interface{} `yaml:"common"             json:"common"`
	HostsAndClusters interface{} `yaml:"hosts_and_clusters" json:"hosts_and_clusters"`
	VirtualMachines  interface{} `yaml:"virtual_machines"   json:"virtual_machines"`
	Storage          interface{} `yaml:"storage"            json:"storage"`
	Network          interface{} `yaml:"network"            json:"network"`
	Supervisor       interface{} `yaml:"supervisor"         json:"supervisor"`
}

// Config is the final struct for validated configuration.
type Config struct {
	// Enabled is whether the vSphere configuration is enabled.
	Enabled bool `yaml:"enabled" json:"enabled" default:"false" validate:"required"`

	// Common is the common configuration for the vSphere environment.
	// This is required if Enabled is true.
	Common Common `yaml:"common" json:"common" validate:"omitempty"`

	// HostsAndClusters is the configuration for the hosts and clusters.
	// This is required if Enabled is true.
	HostsAndClusters HostsAndClusters `yaml:"hosts_and_clusters" json:"hosts_and_clusters" validate:"omitempty"`

	// Storage is the configuration for the storage.
	// This is required if Enabled is true.
	Storage Storage `yaml:"storage" json:"storage" validate:"omitempty"`

	// Network is the configuration for the network.
	// This is required if Enabled is true.
	Network Network `yaml:"network" json:"network" validate:"omitempty"`

	// VirtualMachines is the configuration for the virtual machines.
	// This is required if Enabled is true.
	VirtualMachines VirtualMachines `yaml:"virtual_machines" json:"virtual_machines" validate:"omitempty"`

	// Supervisor is the configuration for the supervisor.
	// This is required if Enabled is true.
	Supervisor Supervisor `yaml:"supervisor" json:"supervisor" validate:"omitempty"`
}

// ValidateVSphere implements custom validation logic for VSphere
func (vSphereConfig *Config) ValidateVSphere(ctx *pulumi.Context) error {
	validate := validator.New()

	/*
		#########################
		Custom validation functions
		#########################
	*/

	// Validator: 'vSphereName'
	if err := validate.RegisterValidation("vSphereName", validateVSphereName); err != nil {
		return errors.New("failed to register vSphereName validator: %w" + err.Error())
	}

	/*
		#########################
		Validation
		#########################
	*/

	// Validate the 'vsphere' top-level section.
	if err := validate.Struct(vSphereConfig); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			validationErr := utils.NewValidationError(validationErrors)
			return fmt.Errorf("\n%s", validationErr.Error())
		}
		return err
	}

	// If vSphere is not enabled, return nil as there is nothing to validate.
	if !vSphereConfig.Enabled {
		return nil
	}

	// If vSphere is enabled, begin validating the nested sections.
	if vSphereConfig.Enabled {

		// Section: Common
		if err := validate.Struct(vSphereConfig.Common); err != nil {
			utils.LogInfo(ctx, "✨ Validating 'vsphere.common' configuration...")
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				// If Common.Enabled is false, ignore all validation errors except the Enabled field itself
				if !vSphereConfig.Common.Enabled {
					return nil
				}
				// If Common.Enabled is true, return all validation errors
				validationErr := utils.NewValidationError(validationErrors)
				return fmt.Errorf(
					"💥 The 'common' section under the 'vsphere' cloud provider in the configuration file is invalid:\n%s",
					validationErr.Error(),
				)
			}
			return err
		}

		// Section: HostsAndClusters
		if err := validate.Struct(vSphereConfig.HostsAndClusters); err != nil {
			utils.LogInfo(ctx, "✨ Validating 'vsphere.hosts_and_clusters' configuration...")
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				// If HostsAndClusters.Enabled is false, ignore all validation errors except the Enabled field itself
				if !vSphereConfig.HostsAndClusters.Enabled {
					return nil
				}
				// If HostsAndClusters.Enabled is true, return all validation errors
				validationErr := utils.NewValidationError(validationErrors)
				return fmt.Errorf(
					"💥 The 'hosts_and_clusters' section under the 'vsphere' cloud provider in the configuration file is invalid:\n%s",
					validationErr.Error(),
				)
			}
			return err
		}

		// Section: Storage
		if err := validate.Struct(vSphereConfig.Storage); err != nil {
			utils.LogInfo(ctx, "✨ Validating 'vsphere.storage' configuration...")
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				// If Storage.Enabled is false, ignore all validation errors except the Enabled field itself
				if !vSphereConfig.Storage.Enabled {
					return nil
				}
				// If Storage.Enabled is true, return all validation errors
				validationErr := utils.NewValidationError(validationErrors)
				return fmt.Errorf(
					"💥 The 'storage' section under the 'vsphere' cloud provider in the configuration file is invalid:\n%s",
					validationErr.Error(),
				)
			}
			return err
		}

		// Section: Network
		if err := validate.Struct(vSphereConfig.Network); err != nil {
			utils.LogInfo(ctx, "✨ Validating 'vsphere.network' configuration...")
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				// If Network.Enabled is false, ignore all validation errors except the Enabled field itself
				if !vSphereConfig.Network.Enabled {
					return nil
				}
				// If Network.Enabled is true, return all validation errors
				validationErr := utils.NewValidationError(validationErrors)
				return fmt.Errorf(
					"💥 The 'network' section under the 'vsphere' cloud provider in the configuration file is invalid:\n%s",
					validationErr.Error(),
				)
			}
			return err
		}

		// Section: VirtualMachines
		if err := validate.Struct(vSphereConfig.VirtualMachines); err != nil {
			utils.LogInfo(ctx, "✨ Validating 'vsphere.virtual_machines' configuration...")
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				// If VirtualMachines.Enabled is false, ignore all validation errors except the Enabled field itself
				if !vSphereConfig.VirtualMachines.Enabled {
					return nil
				}
				// If VirtualMachines.Enabled is true, return all validation errors
				validationErr := utils.NewValidationError(validationErrors)
				return fmt.Errorf(
					"💥 The 'virtual_machines' section under the 'vsphere' cloud provider in the configuration file is invalid:\n%s",
					validationErr.Error(),
				)
			}
			return err
		}
	}

	return nil

}

// validateVSphereName is a custom validation function that
// ensures a given name contains only the allowed characters.
// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
func validateVSphereName(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	// Only allow alphanumeric, dashes, underscores and dots
	matched, err := regexp.MatchString(`^[a-zA-Z0-9\-_.]+$`, name)
	if err != nil {
		return false
	}
	return matched
}
