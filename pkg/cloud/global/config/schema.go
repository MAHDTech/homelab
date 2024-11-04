package globalconfig

// ConfigRaw is used for the initial YAML parsing and validation.
type ConfigRaw struct {
	Enabled interface{} `yaml:"enabled"`
	Debug   interface{} `yaml:"debug"`
}

// Config is the final struct for validated configuration.
type Config struct {
	Enabled bool `yaml:"enabled"`
	Debug   bool `yaml:"debug"`
}
