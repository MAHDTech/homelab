// Package vsphereschema contains the structs for the VMware configuration.
package vsphereschema

// HostsAndClusters represents the pane in the UI that
// has configuration for Hosts and Clusters.
type HostsAndClusters struct {
	// Enabled is whether the hosts and clusters management is enabled.
	Enabled bool `yaml:"enabled" json:"enabled" default:"false" validate:"required"`

	// Datacenters is a list of vSphere Datacenters.
	// This is optional.
	Datacenters []Datacenter `yaml:"datacenters" json:"datacenters" validate:"omitempty"`

	// Clusters is a list of vSphere Clusters.
	// This is optional.
	Clusters []Cluster `yaml:"clusters" json:"clusters" validate:"omitempty"`

	// ResourcePools is a list of vSphere Resource Pools.
	// This is optional.
	ResourcePools []ResourcePool `yaml:"resourcePools" json:"resourcePools" validate:"omitempty"`
}

// Datacenter is a single vSphere Datacenter.
type Datacenter struct {
	// The Name of the Datacenter.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Name string `yaml:"name" json:"name" validate:"required,min=1,max=80,vSphereName"`

	// Optional parent path to the Datacenter.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Folder string `yaml:"folder" json:"folder" validate:"min=1,max=80,vSphereName"`
}

// Cluster is a single vSphere Cluster.
type Cluster struct {
	// The Name of the Cluster is required.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Name string `yaml:"name" json:"name" validate:"required,min=1,max=80,vSphereName"`

	// The Datacenter that the Cluster belongs to is required.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Datacenter string `yaml:"datacenter" json:"datacenter" validate:"required,min=1,max=80,vSphereName"`
}

// ResourcePool struct defines the configuration for a vSphere Resource Pool.
type ResourcePool struct {
	// The Name of the Resource Pool.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Name string `yaml:"name" json:"name" validate:"required,min=1,max=80,vSphereName"`

	// The CPU configuration for the Resource Pool.
	CPU struct {

		// The Shares configuration for the Resource Pool.
		Shares struct {
			// The Level of the CPU shares.
			// Valid values are: low, normal, high or custom.
			Level string `yaml:"level" json:"level" default:"normal" validate:"required,oneof=low normal high custom"`

			// The Count of the CPU shares.
			// If the Level is custom, the Count is required.
			// The value is a positive integer min of 0 with no max.
			Count int `yaml:"count" json:"count" default:"0" validate:"required_if=Level custom,min=0,numeric"`
		}

		// The Reservation configuration for the Resource Pool.
		Reservation struct {
			// MHz is the reservation in MHz with a min of 0 and max of the total CPU capacity.
			MHz int `yaml:"mhz" json:"mhz" default:"0" validate:"required,min=0"`

			// Expandable is whether the reservation is allowed to expand.
			Expandable bool `yaml:"expandable" json:"expandable" default:"false"`
		}

		// The Limit of the CPU in MHz with a default of -1 which means no limit.
		// Valid values are integers between -1 and infinity.
		Limit int `yaml:"limit" json:"limit" default:"-1" validate:"regexp=^(-1|[1-9][0-9]*)$"`
	} `yaml:"cpu" json:"cpu"`

	// The Memory configuration for the Resource Pool.
	Memory struct {

		// The Shares configuration for the Resource Pool.
		Shares struct {
			// The Level of the memory shares.
			// Valid values are: low, normal, high or custom.
			Level string `yaml:"level" json:"level" default:"normal" validate:"required,oneof=low normal high custom"`

			// The Count of the memory shares.
			// If the Level is custom, the Count is required.
			// The value is a positive integer min of 0 with no max.
			Count int `yaml:"count" json:"count" default:"0" validate:"required_if=Level custom,min=0,numeric"`
		}
		// The Reservation configuration for the Resource Pool.
		Reservation struct {
			// MB is the reservation in MB with a min of 0 and max of the total memory capacity.
			MB int `yaml:"mb" json:"mb" default:"0" validate:"required,min=0"`

			// Expandable is whether the reservation is allowed to expand.
			Expandable bool `yaml:"expandable" json:"expandable" default:"false"`
		}

		// The Limit of the memory in MB with a default of -1 which means no limit.
		// Valid values are integers between -1 and infinity.
		Limit int `yaml:"limit" json:"limit" default:"-1" validate:"regexp=^(-1|[1-9][0-9]*)$"`
	} `yaml:"memory" json:"memory"`
}
