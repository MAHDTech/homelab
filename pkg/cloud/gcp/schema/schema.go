// Package gcpschema contains the structs for the GCP configuration.
package gcpschema

// ConfigRaw is used for the initial YAML parsing and validation.
type ConfigRaw struct {
	Enabled interface{} `yaml:"enabled"`
}

// Config is the final struct for validated configuration.
type Config struct {
	Enabled bool `yaml:"enabled"`
}
