// Package vsphereschema contains the structs for the VMware configuration.
package vsphereschema

// Storage holds the configuration for vSphere storage.
type Storage struct {
	Enabled           bool              `yaml:"enabled"`
	DatastoreClusters DatastoreClusters `yaml:"datastoreClusters"`
	Datastores        Datastores        `yaml:"datastores"`
	StoragePolicies   StoragePolicies   `yaml:"storagePolicies"`
}

// DatastoreClusters is a list of vSphere Datastore Clusters.
type DatastoreClusters struct {
	DatastoreCluster []DatastoreCluster `yaml:"datastoreClusters"`
}

// DatastoreCluster is a single vSphere Datastore Cluster.
type DatastoreCluster struct {
	Name       string `yaml:"name"`
	Datacenter string `yaml:"datacenter"`
	Folder     string `yaml:"folder"`
}

// Datastores is a list of vSphere Datastores.
type Datastores struct {
	Datastore []Datastore `yaml:"datastores"`
}

// Datastore is a single vSphere Datastore.
type Datastore struct {
	Name       string `yaml:"name"`
	Datacenter string `yaml:"datacenter"`
	Folder     string `yaml:"folder"`
	Type       string `yaml:"type"`
	NFS        NFS    `yaml:"nfs"`
	VMFS       VMFS   `yaml:"vmfs"`
}

// NFS is the configuration for an NFS Datastore.
type NFS struct {
	Servers []string `yaml:"servers"`
	Path    string   `yaml:"path"`
	Mode    string   `yaml:"mode"`
	Type    string   `yaml:"type"`
}

// VMFS is the configuration for a VMFS Datastore.
type VMFS struct {
	// TODO: Add VMFS configuration
}

// StoragePolicies is a list of vSphere Storage Policies.
type StoragePolicies struct {
	StoragePolicy []StoragePolicy `yaml:"storagePolicies"`
}

// StoragePolicy is a single vSphere Storage Policy.
type StoragePolicy struct {
	Name string `yaml:"name"`
}
