// Package vsphereschema contains the structs for the VMware configuration.
package vsphereschema

// ConfigRaw is used for the initial YAML parsing and validation.
type ConfigRaw struct {
	Enabled          interface{} `yaml:"enabled"`
	Common           interface{} `yaml:"common"`
	HostsAndClusters interface{} `yaml:"hosts_and_clusters"`
	VirtualMachines  interface{} `yaml:"virtual_machines"`
	Storage          interface{} `yaml:"storage"`
	Network          interface{} `yaml:"network"`
	Supervisor       interface{} `yaml:"supervisor"`
}

// Config is the final struct for validated configuration.
type Config struct {
	Enabled          bool             `yaml:"enabled"`
	Common           Common           `yaml:"common"`
	HostsAndClusters HostsAndClusters `yaml:"hosts_and_clusters"`
	VirtualMachines  VirtualMachines  `yaml:"virtual_machines"`
	Storage          Storage          `yaml:"storage"`
	Network          Network          `yaml:"network"`
	Supervisor       Supervisor       `yaml:"supervisor"`
}
