// Package vsphereschema contains the structs for the VMware configuration.
package vsphereschema

import "strings"

// VirtualMachines holds the configuration for vSphere virtual machines.
type VirtualMachines struct {
	// Enabled is whether the virtual machines configuration is enabled.
	Enabled bool `yaml:"enabled" json:"enabled" default:"false" validate:"required"`

	// VMs is a list of vSphere VMs.
	VMs []VM `yaml:"vms" json:"vms" validate:"required"`
}

// VM is a single vSphere VM.
type VM struct {
	// Name is the name of the VM.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Name string `yaml:"name" json:"name" validate:"required,min=1,max=80,vSphereName"`

	// Description is an optional description of the VM.
	Description string `yaml:"description" json:"description"`

	// Size is the size of the VM.
	// Valid values are: tiny, small, medium, large, xlarge, xxlarge
	Size VMSKU `yaml:"size" json:"size" validate:"required,oneof=tiny small medium large xlarge"`

	// ResourcePool is the resource pool of the VM.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	ResourcePool string `yaml:"resourcePool" json:"resourcePool" validate:"required,min=1,max=80,vSphereName"`

	// Folder is the folder of the VM.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Folder string `yaml:"folder" json:"folder" validate:"required,min=1,max=80,vSphereName"`

	// Template is the template of the VM.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Template string `yaml:"template" json:"template" validate:"required,min=1,max=80,vSphereName"`
}

// VMSKU is a valid vSphere VM size.
// Valid values are: tiny, small, medium, large, xlarge
type VMSKU string

// VMSKUSpec defines the CPU and Memory configuration
type VMSKUSpec struct {
	// Number of vCPUs
	// An integer between 1 and 32.
	CPU int `yaml:"cpu" json:"cpu" validate:"required,min=1,max=32"`

	// Memory in MB
	// An integer between 1024 and 131072.
	Memory int64 `yaml:"memory" json:"memory" validate:"required,min=1024,max=131072"`
}

// GetVMSKUSpec returns the CPU and Memory specification for a given VM SKU
func GetVMSKUSpec(sku string) VMSKUSpec {

	switch strings.ToLower(string(sku)) {

	// Tiny has 1 vCPU and 1GB of RAM.
	case "tiny":
		return VMSKUSpec{
			CPU:    1,
			Memory: 1 * 1024, // 1GB
		}

	// Small has 1 vCPUs and 4GB of RAM.
	case "small":
		return VMSKUSpec{
			CPU:    1,
			Memory: 4 * 1024, // 4GB
		}

	// Medium has 2 vCPUs and 8GB of RAM.
	case "medium":
		return VMSKUSpec{
			CPU:    2,
			Memory: 8 * 1024, // 8GB
		}

	// Large has 4 vCPUs and 16GB of RAM.
	case "large":
		return VMSKUSpec{
			CPU:    4,
			Memory: 16 * 1024, // 16GB
		}

	// XLarge has 4 vCPUs and 32GB of RAM.
	case "xlarge":
		return VMSKUSpec{
			CPU:    4,
			Memory: 32 * 1024, // 32GB
		}

	// XXLarge has 8 vCPUs and 64GB of RAM.
	case "xxlarge":
		return VMSKUSpec{
			CPU:    8,
			Memory: 64 * 1024, // 64GB
		}

	// Default to Tiny if an unknown SKU is provided.
	default:
		return VMSKUSpec{
			CPU:    1,
			Memory: 1 * 1024, // Default to Tiny
		}
	}
}
