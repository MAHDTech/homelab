// Package vsphereschema contains the structs for the VMware configuration.
package vsphereschema

// Supervisor is the configuration for the vSphere supervisor.
type Supervisor struct {
	Enabled        bool              `yaml:"enabled"`
	Name           string            `yaml:"name"`
	Cluster        string            `yaml:"cluster"`
	ContentLibrary string            `yaml:"content_library"`
	Network        SupervisorNetwork `yaml:"network"`
}

// SupervisorNetwork is the configuration for the vSphere supervisor network.
type SupervisorNetwork struct {
	Type string `yaml:"type"`
}

// TODO: Complete Supervisor configuration schema.
