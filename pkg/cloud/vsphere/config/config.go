package vsphereconfig

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func VerifyConfig(ctx *pulumi.Context, cfg ConfigRaw) (Config, error) {

	var vcenter Vcenter
	var infrastructure Infrastructure

	// #########################
	// vCenter
	// #########################

	// vCenter is required.
	if cfg.Vcenter == nil {
		message := "ERROR: A vSphere vCenter configuration is required in 'vcenter'"
		return Config{}, fmt.Errorf(message)
	}

	// Unmarshal the vCenter configuration.
	if err := mapstructure.Decode(cfg.Vcenter, &vcenter); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal vCenter configuration: %w", err)
	}

	// vCenter Cluster Name
	// The vCenter Cluster name is required and must adhere to these rules:
	// 		- Must start with a letter
	// 		- Can only contain letters, numbers, and hyphens
	// 		- Must be less than 80 characters
	if cfg.Vcenter.Cluster == "" {
		message := "ERROR: A vSphere Cluster name is required in 'vcenter.cluster'"
		return Config{}, fmt.Errorf(message)
	}
	for _, char := range cfg.Vcenter.Cluster {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') || char == '-') {
			message := "ERROR: The vSphere Cluster name can only contain letters, numbers, and hyphens"
			return Config{}, fmt.Errorf(message)
		}
	}
	if len(cfg.Vcenter.Cluster) > 80 {
		message := "ERROR: The vSphere Cluster name must be less than 80 characters"
		return Config{}, fmt.Errorf(message)
	}
	vcenter.Cluster = cfg.Vcenter.Cluster

	// vCenter Datastore Name
	// The vCenter Datastore name is required and must adhere to these rules:
	// 		- Must start with a letter
	// 		- Can only contain letters, numbers, and hyphens
	// 		- Must be less than 80 characters
	if cfg.Vcenter.Datastore == "" {
		message := "ERROR: A vSphere Datastore name is required in 'vcenter.datastore'"
		return Config{}, fmt.Errorf(message)
	}
	for _, char := range cfg.Vcenter.Datastore {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') || char == '-') {
			message := "ERROR: The vSphere Datastore name can only contain letters, numbers, and hyphens"
			return Config{}, fmt.Errorf(message)
		}
	}
	if len(cfg.Vcenter.Datastore) > 80 {
		message := "ERROR: The vSphere Datastore name must be less than 80 characters"
		return Config{}, fmt.Errorf(message)
	}
	vcenter.Datastore = cfg.Vcenter.Datastore

	// #########################
	// Configuration
	// #########################

	// Build the final configuration to be returned.
	configFinal := Config{
		Enabled: cfg.Enabled.(bool),
		Vcenter: vcenter,
	}

	return configFinal, nil
}
