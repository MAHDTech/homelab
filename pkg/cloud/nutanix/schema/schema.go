// Package nutanixschema contains the structs for the Nutanix configuration.
package nutanixschema

// ConfigRaw is used for the initial YAML parsing and validation.
type ConfigRaw struct {
	Enabled interface{} `yaml:"enabled"`
}

// Config is the final struct for validated configuration.
type Config struct {
	Enabled bool `yaml:"enabled"`
}
