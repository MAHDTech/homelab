// Package vsphereschema contains the structs for the VMware configuration.
package vsphereschema

// Storage holds the configuration for vSphere storage.
type Storage struct {
	// Enabled is whether the storage configuration is enabled.
	Enabled bool `yaml:"enabled" json:"enabled" default:"false" validate:"required"`

	// DatastoreClusters are the configuration for the datastore clusters.
	// This is optional.
	DatastoreClusters []DatastoreCluster `yaml:"datastoreClusters" json:"datastoreClusters" validate:"omitempty"`

	// Datastores are the configuration for the datastores.
	// This is optional.
	Datastores []Datastore `yaml:"datastores" json:"datastores" validate:"omitempty"`

	// StoragePolicies are the configuration for the storage policies.
	// This is optional.
	StoragePolicies []StoragePolicy `yaml:"storagePolicies" json:"storagePolicies" validate:"omitempty"`
}

// DatastoreCluster is a single vSphere Datastore Cluster.
type DatastoreCluster struct {
	// Name is the name of the Datastore Cluster.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Name string `yaml:"name" json:"name" validate:"required,min=1,max=80,vSphereName"`

	// Datacenter is the name of the Datacenter that the Datastore Cluster belongs to.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Datacenter string `yaml:"datacenter" json:"datacenter" validate:"required,min=1,max=80,vSphereName"`

	// Folder is the name of the Folder that the Datastore Cluster belongs to.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Folder string `yaml:"folder" json:"folder" validate:"required,min=1,max=80,vSphereName"`
}

// Datastore is a single vSphere Datastore.
type Datastore struct {
	// Name is the name of the Datastore.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Name string `yaml:"name" json:"name" validate:"required,min=1,max=80,vSphereName"`

	// Datacenter is the name of the Datacenter that the Datastore belongs to.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Datacenter string `yaml:"datacenter" json:"datacenter" validate:"required,min=1,max=80,vSphereName"`

	// Folder is the name of the Folder that the Datastore belongs to.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Folder string `yaml:"folder" json:"folder" validate:"required,min=1,max=80,vSphereName"`

	// Type is the type of the Datastore.
	// Valid values are: nfs, vmfs
	Type string `yaml:"type" json:"type" validate:"required,oneof=nfs vmfs"`

	// NFS is the configuration for an NFS Datastore.
	NFS NFS `yaml:"nfs" json:"nfs" validate:"required_if=Type nfs"`

	// VMFS is the configuration for a VMFS Datastore.
	VMFS VMFS `yaml:"vmfs" json:"vmfs" validate:"required_if=Type vmfs"`
}

// NFS is the configuration for an NFS Datastore.
type NFS struct {
	// Servers is a list of NFS servers.
	// These can be IP addresses or hostnames.
	Servers []string `yaml:"servers" json:"servers" validate:"required"`

	// Path is the path to the NFS share.
	Path string `yaml:"path" json:"path" validate:"required"`

	// Mode is the mode of the NFS share.
	// Valid values are: readOnly, readWrite
	Mode string `yaml:"mode" json:"mode" validate:"required,oneof=readOnly readWrite"`

	// Type is the type of the NFS share.
	// Valid values are: NFS, NFS41
	Type string `yaml:"type" json:"type" validate:"required,oneof=NFS NFS41"`
}

// VMFS is the configuration for a VMFS Datastore.
type VMFS struct {
	// TODO: Add VMFS configuration
}

// StoragePolicy is a single vSphere Storage Policy.
type StoragePolicy struct {
	// Name is the name of the Storage Policy.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Name string `yaml:"name" json:"name" validate:"required,min=1,max=80,vSphereName"`
	// TODO: Add Storage Policy configuration here.
}
