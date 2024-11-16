// Package schema contains the structs for the configuration.
package schema

import (
	globalschema "homelab/pkg/cloud/global/schema"

	awsschema "homelab/pkg/cloud/aws/schema"
	azureschema "homelab/pkg/cloud/azure/schema"
	gcpschema "homelab/pkg/cloud/gcp/schema"
	nutanixschema "homelab/pkg/cloud/nutanix/schema"
	vsphereschema "homelab/pkg/cloud/vsphere/schema"
)

// ConfigRaw is used for the initial YAML parsing and validation.
type ConfigRaw struct {
	Global globalschema.ConfigRaw `yaml:"global"`

	AWS     awsschema.ConfigRaw     `yaml:"aws"`
	Azure   azureschema.ConfigRaw   `yaml:"azure"`
	GCP     gcpschema.ConfigRaw     `yaml:"gcp"`
	Nutanix nutanixschema.ConfigRaw `yaml:"nutanix"`
	VSphere vsphereschema.ConfigRaw `yaml:"vsphere"`
}

// Config represents the entire configuration object.
// Due to the number of configuration items, the config
// is broken into multiple sections based on the cloud provider.
type Config struct {
	Global globalschema.Config `yaml:"global"`

	AWS     awsschema.Config     `yaml:"aws"`
	Azure   azureschema.Config   `yaml:"azure"`
	GCP     gcpschema.Config     `yaml:"gcp"`
	VSphere vsphereschema.Config `yaml:"vsphere"`
	Nutanix nutanixschema.Config `yaml:"nutanix"`
}
