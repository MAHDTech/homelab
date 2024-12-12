// Package vsphereschema contains the structs for the VMware configuration.
package vsphereschema

// Common holds the configuration for vSphere common.
type Common struct {
	// Enabled is whether the common configuration is enabled.
	Enabled bool `yaml:"enabled" json:"enabled" default:"false" validate:"required"`

	// Folders is a list of optional folders to create.
	// This is optional.
	Folders []Folder `yaml:"folders" json:"folders" validate:"omitempty"`

	// ContentLibraries is a list of optional content libraries to create.
	// This is optional.
	ContentLibraries []ContentLibrary `yaml:"contentLibraries" json:"contentLibraries" validate:"omitempty"`

	// Tags contains categories and tags.
	// This is optional.
	Tags Tags `yaml:"tags" json:"tags" validate:"omitempty"`
}

// Folder is a single folder type.
type Folder struct {
	// Name is the name of the folder and is required to be given.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Name string `yaml:"name" json:"name" validate:"required,min=1,max=80,vSphereName"`

	// Type is the Type of the folder and is required.
	// Valid values are: datacenter, host, vm, datastore, network
	Type string `yaml:"type" json:"type" default:"vm" validate:"required,oneof=datacenter host vm datastore network"`

	// Path is an optional parent path to the folder.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Path string `yaml:"path" json:"path" validate:"name"`
}

// ContentLibrary struct defines a content library which can be one of two types.
// 1. Published. Images are uploaded to the content library.
// 2. Subscribed. Images are downloaded from remote URL.
type ContentLibrary struct {
	// Name is the name of the content library and is required.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Name string `yaml:"name" json:"name" validate:"required,min=1,max=80,vSphereName"`

	// Description is an optional description of the content library.
	Description string `yaml:"description" json:"description"`

	// Type is the type of the content library and is required.
	// Valid values are: published, subscribed
	Type string `yaml:"type" json:"type" default:"published" validate:"required,oneof=published subscribed"`

	// Publish is the configuration for publishing which is required if the type is published.
	Publish *Publish `yaml:"publish" json:"publish" validate:"required_if=Type published"`

	// Subscription is the configuration for subscribing which is required if the type is subscribed.
	Subscription *Subscription `yaml:"subscription" json:"subscription" validate:"required_if=Type subscribed"`
}

// Publish struct defines the configuration for publishing
// which is required if the type is published.
type Publish struct {
	// Authentication is the authentication method for the content library.
	// Valid values are: none, basic
	Authentication string `yaml:"authentication" json:"authentication" default:"none" validate:"required,oneof=none basic"`

	// Username is the username for the content library.
	// This is required if the authentication method is basic.
	Username string `yaml:"username" json:"username" validate:"required_if=Authentication basic"`

	// Password is the password for the content library.
	// This is required if the authentication method is basic.
	Password string `yaml:"password" json:"password" validate:"required_if=Authentication basic"`
}

// Subscription struct defines the configuration for subscribing.
type Subscription struct {
	// Authentication is the authentication method for the content library.
	// Valid values are: none, basic
	Authentication string `yaml:"authentication" json:"authentication" default:"none" validate:"required,oneof=none basic"`

	// Username is the username for the content library.
	// This is required if the authentication method is basic.
	Username string `yaml:"username" json:"username" validate:"required_if=Authentication basic"`

	// Password is the password for the content library.
	// This is required if the authentication method is basic.
	Password string `yaml:"password" json:"password" validate:"required_if=Authentication basic"`

	// URL is the URL of the remote content library.
	// A URL must start with http:// or https://.
	URL string `yaml:"url" json:"url" validate:"required,url"`

	// Method is the method of the remote content library.
	// Valid values are: automatic, manual
	Method string `yaml:"method" json:"method" default:"manual" validate:"required,oneof=automatic manual"`
}

// Tags is a list of categories and tags.
type Tags struct {
	// Categories is a list of categories.
	Categories []Category `yaml:"categories" json:"categories" validate:"omitempty"`

	// Tags is a list of tags.
	Tags []Tag `yaml:"tags" json:"tags" validate:"omitempty"`
}

// Category is a single category type.
type Category struct {
	// Name is the name of the category and is required.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Name string `yaml:"name" json:"name" validate:"required,min=1,max=80,vSphereName"`

	// Description is an optional description of the category.
	Description string `yaml:"description" json:"description"`

	// Cardinality is the cardinality of the category.
	// Valid values are: single, multiple
	Cardinality string `yaml:"cardinality" json:"cardinality" validate:"required,oneof=single multiple"`
}

// Tag is a single tag type.
type Tag struct {
	// Name is the name of the tag and is required.
	// The length must be between 1 and 80 characters.
	// Only alphanumeric characters, dashes, underscores, periods and spaces are allowed.
	Name string `yaml:"name" json:"name" validate:"required,min=1,max=80,vSphereName"`

	// Category is the category of the tag and is required.
	Category string `yaml:"category" json:"category" validate:"required"`

	// Description is an optional description of the tag.
	Description string `yaml:"description" json:"description"`
}
