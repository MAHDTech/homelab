// Package vsphereschema contains the structs for the VMware configuration.
package vsphereschema

// Network is the configuration for the vSphere network.
type Network struct {
	Enabled      bool          `yaml:"enabled"`
	DVSwitches   []DVSwitch    `yaml:"dvswitches"`
	DVPortGroups []DVPortGroup `yaml:"dvportgroups"`
}

// DVSwitch is the configuration for a vSphere Distributed Virtual Switch.
type DVSwitch struct {
	Name string `yaml:"name"`
}

// DVPortGroup is the configuration for a vSphere Distributed Virtual Port Group.
type DVPortGroup struct {
	Name     string   `yaml:"name"`
	DVSwitch DVSwitch `yaml:"dvswitch"`
	VLAN     int      `yaml:"vlan"`
}
