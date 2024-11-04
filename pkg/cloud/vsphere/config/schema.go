package vsphereconfig

import "strings"

// ConfigRaw is used for the initial YAML parsing and validation.
type ConfigRaw struct {
	Enabled        interface{} `yaml:"enabled"`
	Vcenter        interface{} `yaml:"vcenter"`
	Infrastructure interface{} `yaml:"infrastructure"`
}

// Config is the final struct for validated configuration.
type Config struct {
	Enabled        bool           `yaml:"enabled"`
	Vcenter        Vcenter        `yaml:"vcenter"`
	Infrastructure Infrastructure `yaml:"infrastructure"`
}

// Vcenter holds connection information for VMware vCenter.
type Vcenter struct {
	Cluster    string `yaml:"cluster"`
	Datacenter string `yaml:"datacenter"`
	Datastore  string `yaml:"datastore"`
}

// Infrastructure holds the configuration for the vSphere infrastructure.
type Infrastructure struct {
	VMs []VM `yaml:"vms"`
}

// VM holds the configuration for a vSphere VM.
type VM struct {
	Name         string `yaml:"name"`
	Description  string `yaml:"description"`
	Size         string `yaml:"size"`
	ResourcePool string `yaml:"resourcePool"`
	Folder       string `yaml:"folder"`
}

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
