// Package vsphereschema contains the structs for the VMware configuration.
package vsphereschema

// Common holds the configuration for vSphere common.
type Common struct {
	Folders          Folders          `yaml:"folders"`
	ContentLibraries ContentLibraries `yaml:"contentLibraries"`
}

// Folders is a list of folder types.
type Folders struct {
	Folder []Folder `yaml:"folders"`
}

// Folder is a single folder type.
type Folder struct {
	// The Name of the folder.
	Name string `yaml:"name"`
	// The Type of the folder.
	// Valid values are: datacenter, host, vm, datastore, network
	Type string `yaml:"type" default:"vm"`
	// Optional parent path to the folder.
	Path string `yaml:"path"`
}

// ContentLibraries is a list of content libraries.
type ContentLibraries struct {
	ContentLibrary []ContentLibrary `yaml:"contentLibraries"`
}

// ContentLibrary struct defines a content library which can be one of two types.
// 1. Published. Images are uploaded to the content library.
// 2. Subscribed. Images are downloaded from remote URL.
type ContentLibrary struct {
	Name         string       `yaml:"name"`
	Description  string       `yaml:"description"`
	Publish      Publish      `yaml:"publish"`
	Subscription Subscription `yaml:"subscription"`
}

// Publish struct defines the configuration for publishing.
type Publish struct {
	Enabled        bool   `yaml:"enabled"        default:"false"`
	Authentication string `yaml:"authentication" default:"none"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
}

// Subscription struct defines the configuration for subscribing.
type Subscription struct {
	Enabled        bool   `yaml:"enabled"        default:"false"`
	Authentication string `yaml:"authentication" default:"none"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	URL            string `yaml:"url"`
	Method         string `yaml:"method"         default:"manual"`
}
