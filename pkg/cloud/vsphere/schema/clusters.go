// Package vsphereschema contains the structs for the VMware configuration.
package vsphereschema

// HostsAndClusters represents the pane in the UI that
// has configuration for Hosts and Clusters.
type HostsAndClusters struct {
	Enabled       bool          `yaml:"enabled"`
	Datacenters   Datacenters   `yaml:"datacenters"`
	Clusters      Clusters      `yaml:"clusters"`
	ResourcePools ResourcePools `yaml:"resourcePools"`
}

// Datacenters is a list of vSphere Datacenters.
type Datacenters struct {
	Datacenter []Datacenter `yaml:"datacenters"`
}

// Datacenter is a single vSphere Datacenter.
type Datacenter struct {
	// The Name of the Datacenter.
	Name string `yaml:"name"`
	// Optional parent path to the Datacenter.
	Folder string `yaml:"folder"`
}

// Clusters is a list of vSphere Clusters.
type Clusters struct {
	Cluster []Cluster `yaml:"clusters"`
}

// Cluster is a single vSphere Cluster.
type Cluster struct {
	// The Name of the Cluster.
	Name string `yaml:"name"`
	// The Datacenter that the Cluster belongs to.
	Datacenter string `yaml:"datacenter"`
}

// ResourcePools is a list of vSphere Resource Pools.
type ResourcePools struct {
	ResourcePool []ResourcePool `yaml:"resourcePools"`
}

// ResourcePool struct defines the configuration for a vSphere Resource Pool.
type ResourcePool struct {
	// The Name of the Resource Pool.
	Name string `yaml:"name"`
	// The CPU configuration for the Resource Pool.
	CPU struct {
		Shares struct {
			Level string `yaml:"level" default:"normal"`
			Count int    `yaml:"count" default:"0"`
		}
		Reservation int `yaml:"reservation" default:"0"`
		Limit       int `yaml:"limit" default:"-1"`
	} `yaml:"cpu"`
	// The Memory configuration for the Resource Pool.
	Memory struct {
		Shares struct {
			Level string `yaml:"level" default:"normal"`
			Count int    `yaml:"count" default:"0"`
		}
		Reservation int `yaml:"reservation" default:"0"`
		Limit       int `yaml:"limit" default:"-1"`
	} `yaml:"memory"`
}
