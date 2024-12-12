// Package vsphereschema contains the structs for the VMware configuration.
package vsphereschema

// Supervisor is the configuration for the vSphere supervisor.
type Supervisor struct {
	// Enabled is whether the supervisor configuration is enabled.
	Enabled bool `yaml:"enabled" json:"enabled" default:"false" validate:"required"`

	// Name is the name of the supervisor.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Name string `yaml:"name" json:"name" validate:"required,min=1,max=80,vSphereName"`

	// Cluster is the name of the cluster to deploy the supervisor to.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Cluster string `yaml:"cluster" json:"cluster" validate:"required,min=1,max=80,vSphereName"`

	// ContentLibrary is the name of the content library to use for the supervisor.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	ContentLibrary string `yaml:"contentLibrary" json:"contentLibrary" validate:"required,min=1,max=80,vSphereName"`

	// Network is the network configuration for the supervisor.
	Network SupervisorNetwork `yaml:"network" json:"network" validate:"required"`

	// Storage is the storage configuration for the supervisor.
	Storage SupervisorStorage `yaml:"storage" json:"storage" validate:"required"`

	// Management is the management configuration for the supervisor.
	Management SupervisorManagement `yaml:"management" json:"management" validate:"required"`

	// ControlPlane is the control plane configuration for the supervisor.
	ControlPlane SupervisorControlPlane `yaml:"controlPlane" json:"controlPlane" validate:"required"`

	// Namespaces is the a list of namespaces to be created on the supervisor.
	// This is optional.
	Namespaces *[]SupervisorNamespace `yaml:"namespaces" json:"namespaces" validate:"omitempty"`
}

// SupervisorNetwork is the configuration for the vSphere supervisor network.
type SupervisorNetwork struct {
	// Type is the type of the network.
	Type string `yaml:"type" json:"type"`

	// VSwitch is the name of the Virtual Switch to use for the supervisor.
	// This is only used if the Type is vswitch.
	VSwitch string `yaml:"vswitch" json:"vswitch" validate:"required_if=Type vswitch"`

	// EdgeCluster is the name of the Edge Cluster to use for the supervisor.
	// This is only used if the Type is nsx.
	EdgeCluster string `yaml:"edgecluster" json:"edgecluster" validate:"required_if=Type nsx"`

	// DNS is the DNS configuration for the supervisor.
	DNS SupervisorNetworkDNS `yaml:"dns" json:"dns" validate:"required"`
}

// SupervisorNetworkDNS is the configuration for the vSphere supervisor network DNS.
type SupervisorNetworkDNS struct {
	// Servers is a list of DNS servers.
	Servers []string `yaml:"servers" json:"servers" validate:"required"`

	// Suffix is the DNS suffix.
	Suffix string `yaml:"suffix" json:"suffix" validate:"required"`
}

// SupervisorStorage is the configuration for the vSphere supervisor storage.
type SupervisorStorage struct {
	// Policy is the name of the storage policy to use for the supervisor.
	Policy string `yaml:"policy" json:"policy"`
}

// SupervisorManagement is the configuration for the vSphere supervisor management.
type SupervisorManagement struct {
	// Network is the name of the network to use for the supervisor management.
	Network string `yaml:"network" json:"network" validate:"required"`

	// StartingAddress is the starting IP address of the management network.
	StartingAddress string `yaml:"startingAddress" json:"startingAddress" validate:"required"`

	// SubnetMask is the subnet mask of the management network.
	SubnetMask string `yaml:"subnetMask" json:"subnetMask" validate:"required"`

	// Gateway is the gateway of the management network.
	Gateway string `yaml:"gateway" json:"gateway" validate:"required"`

	// AddressCount is the number of addresses to allocate for the management network.
	AddressCount int `yaml:"addressCount" json:"addressCount" validate:"required"`
}

// SupervisorControlPlane is the configuration for the vSphere supervisor control plane.
type SupervisorControlPlane struct {
	// Size is the size of the control plane.
	Size string `yaml:"size" json:"size" validate:"required,oneof=small medium large"`
}

// SupervisorNamespace is the configuration for a vSphere supervisor namespace.
type SupervisorNamespace struct {
	// Name is the name of the namespace.
	Name string `yaml:"name" json:"name" validate:"required"`

	// Description is the description of the namespace.
	Description string `yaml:"description" json:"description"`
}
