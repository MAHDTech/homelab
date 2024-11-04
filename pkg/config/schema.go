package config

import (
	awsConfig "homelab/pkg/cloud/aws/config"
	azureConfig "homelab/pkg/cloud/azure/config"
	gcpConfig "homelab/pkg/cloud/gcp/config"
	globalConfig "homelab/pkg/cloud/global/config"
	nutanixConfig "homelab/pkg/cloud/nutanix/config"
	vsphereConfig "homelab/pkg/cloud/vsphere/config"
)

// ConfigRaw is used for the initial YAML parsing and validation.
type ConfigRaw struct {
	Global  globalConfig.ConfigRaw  `yaml:"global"`
	AWS     awsConfig.ConfigRaw     `yaml:"aws"`
	Azure   azureConfig.ConfigRaw   `yaml:"azure"`
	GCP     gcpConfig.ConfigRaw     `yaml:"gcp"`
	VSphere vsphereConfig.ConfigRaw `yaml:"vsphere"`
	Nutanix nutanixConfig.ConfigRaw `yaml:"nutanix"`
}

// Config represents the entire configuration object.
// Due to the number of configuration items, the config
// is broken into multiple sections based on the cloud provider.
type Config struct {
	Global  globalConfig.Config  `yaml:"global"`
	AWS     awsConfig.Config     `yaml:"aws"`
	Azure   azureConfig.Config   `yaml:"azure"`
	GCP     gcpConfig.Config     `yaml:"gcp"`
	VSphere vsphereConfig.Config `yaml:"vsphere"`
	Nutanix nutanixConfig.Config `yaml:"nutanix"`
}
