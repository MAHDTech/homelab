// Package vsphereschema contains the structs for the VMware configuration.
package vsphereschema

import "strings"

// VirtualMachines holds the configuration for vSphere virtual machines.
type VirtualMachines struct {
	Enabled bool `yaml:"enabled"`
	VMs     VMs  `yaml:"vms"`
}

// VMs is a list of vSphere VMs.
type VMs struct {
	VM []VM `yaml:"vms"`
}

// VM is a single vSphere VM.
type VM struct {
	Name         string `yaml:"name"`
	Description  string `yaml:"description"`
	Size         VMSKU  `yaml:"size"`
	ResourcePool string `yaml:"resourcePool"`
	Folder       string `yaml:"folder"`
	Template     string `yaml:"template"`
}

// VMSKU is a valid vSphere VM size.
// Valid values are: tiny, small, medium, large, xlarge
type VMSKU string

// VMSKUSpec defines the CPU and Memory configuration
type VMSKUSpec struct {
	// Number of vCPUs
	CPU int
	// Memory in MB
	Memory int64
}

// GetVMSKUSpec returns the CPU and Memory specification for a given VM SKU
func GetVMSKUSpec(sku string) VMSKUSpec {

	switch strings.ToLower(string(sku)) {

	// Small has 2 vCPUs and 8GB of RAM.
	case "small":
		return VMSKUSpec{
			CPU:    2,
			Memory: 8 * 1024, // 8GB
		}

	// Medium has 4 vCPUs and 16GB of RAM.
	case "medium":
		return VMSKUSpec{
			CPU:    4,
			Memory: 16 * 1024, // 16GB
		}

	// Large has 8 vCPUs and 32GB of RAM.
	case "large":
		return VMSKUSpec{
			CPU:    8,
			Memory: 32 * 1024, // 32GB
		}

	// XLarge has 8 vCPUs and 64GB of RAM.
	case "xlarge":
		return VMSKUSpec{
			CPU:    8,
			Memory: 64 * 1024, // 64GB
		}

	// Default to Tiny if an unknown SKU is provided.
	default:
		return VMSKUSpec{
			CPU:    1,
			Memory: 4 * 1024, // Default to Tiny
		}
	}
}
