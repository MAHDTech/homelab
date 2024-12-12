// Package vsphereschema contains the structs for the VMware configuration.
package vsphereschema

// Network is the configuration for the vSphere network.
type Network struct {
	// Enabled is whether the network configuration is enabled.
	Enabled bool `yaml:"enabled" json:"enabled" default:"false" validate:"required"`

	// DVSwitches is a list of Distributed Virtual Switches.
	// This is optional.
	DVSwitches []DVSwitch `yaml:"dvswitches" json:"dvswitches" validate:"omitempty"`

	// DVPortGroups is a list of Distributed Virtual Port Groups.
	// This is optional.
	DVPortGroups []DVPortGroup `yaml:"dvportgroups" json:"dvportgroups" validate:"omitempty"`

	// LoadBalancer is the configuration for a vSphere Load Balancer.
	// This is optional.
	LoadBalancer LoadBalancer `yaml:"loadbalancer" json:"loadbalancer" validate:"omitempty"`
}

// DVSwitch is the configuration for a vSphere Distributed Virtual Switch.
type DVSwitch struct {
	// Name is the name of the Distributed Virtual Switch.
	// Name is required and needs to be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Name string `yaml:"name" json:"name" validate:"required,min=1,max=80,vSphereName"`
}

// DVPortGroup is the configuration for a vSphere Distributed Virtual Port Group.
type DVPortGroup struct {
	// Name is the name of the Distributed Virtual Port Group.
	// Name is required and needs to be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Name string `yaml:"name" json:"name" validate:"required,min=1,max=80,vSphereName"`

	// DVSwitch is the Distributed Virtual Switch Name that the port group belongs to.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	DVSwitch string `yaml:"dvswitch" json:"dvswitch" validate:"required,min=1,max=80,vSphereName"`

	// VLAN is the VLAN ID of the port group.
	// Valid values are integers between 1 and 4094.
	VLAN int `yaml:"vlan" json:"vlan" validate:"required,min=1,max=4094"`
}

// LoadBalancer is the configuration for a vSphere Load Balancer.
type LoadBalancer struct {
	// TODO: Add Load Balancer configuration here.
}
