// Package vsphereconfig contains the functions for verifying the VMware configuration.
package vsphereconfig

import (
	"errors"
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	vsphereschema "homelab/pkg/cloud/vsphere/schema"

	utils "homelab/pkg/utils"
)

// VerifyConfig validates the vSphere configuration by
// taking a raw configuration and returning a fully
// validated configuration or an error.
// We use a raw config to handle the unmarshalling
// into an interface before performing validation.
func VerifyConfig(
	ctx *pulumi.Context,
	configRaw vsphereschema.ConfigRaw,
) (vsphereschema.Config, error) {

	// #########################
	// Enabled
	// #########################

	// Validate that 'Enabled' is a boolean.
	_, ok := configRaw.Enabled.(bool)
	if !ok {
		message := "💥 'enabled' in the 'vsphere' section must be a boolean. You have provided: " + fmt.Sprint(
			configRaw.Enabled,
		)
		return vsphereschema.Config{}, errors.New(message)
	}

	// If 'Enabled' is false, skip the rest of the configuration verification.
	if !configRaw.Enabled.(bool) {
		utils.LogInfo(ctx, "🛈 vSphere is disabled, skipping configuration verification.")
		return vsphereschema.Config{Enabled: false}, nil
	}

	// #########################
	// Common
	// #########################

	configCommon, err := verifyCommonConfig(ctx, configRaw.Common)
	if err != nil {
		message := "💥 The 'common' section in the 'vsphere' configuration is invalid: %s"
		return vsphereschema.Config{}, fmt.Errorf(message, err)
	}

	// ########################
	// Hosts and Clusters
	// ########################

	configHostsAndClusters, err := verifyHostsAndClustersConfig(ctx, configRaw.HostsAndClusters)
	if err != nil {
		message := "💥 The 'hostsandclusters' section in the 'vsphere' configuration is invalid: %s"
		return vsphereschema.Config{}, fmt.Errorf(message, err)
	}

	// #########################
	// Virtual Machines
	// #########################

	configVirtualMachines, err := verifyVirtualMachinesConfig(ctx, configRaw.VirtualMachines)
	if err != nil {
		message := "💥 The 'virtualmachines' section in the 'vsphere' configuration is invalid: %s"
		return vsphereschema.Config{}, fmt.Errorf(message, err)
	}

	// #########################
	// Storage
	// #########################

	configStorage, err := verifyStorageConfig(ctx, configRaw.Storage)
	if err != nil {
		message := "💥 The 'storage' section in the 'vsphere' configuration is invalid: %s"
		return vsphereschema.Config{}, fmt.Errorf(message, err)
	}

	// #########################
	// Network
	// #########################

	configNetwork, err := verifyNetworkConfig(ctx, configRaw.Network)
	if err != nil {
		message := "💥 The 'network' section in the 'vsphere' configuration is invalid: %s"
		return vsphereschema.Config{}, fmt.Errorf(message, err)
	}

	// #########################
	// Supervisor
	// #########################

	// Validate the Supervisor configuration here...
	configSupervisor, err := verifySupervisorConfig(ctx, configRaw.Supervisor)
	if err != nil {
		message := "💥 The 'supervisor' section in the 'vsphere' configuration is invalid: %s"
		return vsphereschema.Config{}, fmt.Errorf(message, err)
	}

	// #########################
	// Configuration
	// #########################

	// Build the final configuration to be returned.

	configFinal := vsphereschema.Config{
		Enabled:          configRaw.Enabled.(bool),
		Common:           configCommon,
		HostsAndClusters: configHostsAndClusters,
		VirtualMachines:  configVirtualMachines,
		Storage:          configStorage,
		Network:          configNetwork,
		Supervisor:       configSupervisor,
	}

	utils.LogInfo(ctx, "✅ vSphere configuration verified")

	return configFinal, nil
}

/*
	// #########################
	// vCenter
	// #########################

	// vCenter is required.
	if cfg.Vcenter == nil {
		message := "💥 A vSphere vCenter configuration is required in 'vcenter'"
		return vsphereschema.Config{}, fmt.Errorf(message)
	}

	// Unmarshal the vCenter configuration into an interface.
	if err := mapstructure.Decode(cfg.Vcenter, &vcenter); err != nil {
		return vsphereschema.Config{}, fmt.Errorf("failed to unmarshal vCenter configuration: %w", err)
	}

	// vCenter Datacenter Name
	// The vCenter Datacenter name is required and must adhere to these rules:
	// 		- Must start with a letter
	// 		- Can only contain letters, numbers, spaces and hyphens
	// 		- Must be less than 80 characters
	if vcenter.Datacenter == "" {
		message := "💥 A vSphere Datacenter name is required in 'vcenter.datacenter'"
		return vsphereschema.Config{}, fmt.Errorf(message)
	}
	for _, char := range vcenter.Datacenter {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') || char == '-' || char == ' ') {
			message := "💥 The vSphere Datacenter name can only contain letters, numbers, spaces and hyphens. You provided: %s"
			return vsphereschema.Config{}, fmt.Errorf(message, string(char))
		}
	}
	if len(vcenter.Datacenter) > 80 {
		message := "💥 The vSphere Datacenter name must be less than 80 characters"
		return vsphereschema.Config{}, fmt.Errorf(message)
	}

	// The vCenter Cluster name is required and must adhere to these rules:
	// 		- Must start with a letter
	// 		- Can only contain letters, numbers, and hyphens
	// 		- Must be less than 80 characters
	if vcenter.Cluster == "" {
		message := "💥 A vSphere Cluster name is required in 'vcenter.cluster'"
		return vsphereschema.Config{}, fmt.Errorf(message)
	}
	for _, char := range vcenter.Cluster {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') || char == '-') {
			message := "💥 The vSphere Cluster name can only contain letters, numbers, and hyphens. You provided: %s"
			return vsphereschema.Config{}, fmt.Errorf(message, string(char))
		}
	}
	if len(vcenter.Cluster) > 80 {
		message := "💥 The vSphere Cluster name must be less than 80 characters"
		return vsphereschema.Config{}, fmt.Errorf(message)
	}

	// vCenter Datastore Name
	// The vCenter Datastore name is required and must adhere to these rules:
	// 		- Must start with a letter
	// 		- Can only contain letters, numbers, and hyphens
	// 		- Must be less than 80 characters
	if vcenter.Datastore == "" {
		message := "💥 A vSphere Datastore name is required in 'vcenter.datastore'"
		return vsphereschema.Config{}, fmt.Errorf(message)
	}
	for _, char := range vcenter.Datastore {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') || char == '-') {
			message := "💥 The vSphere Datastore name can only contain letters, numbers, and hyphens. You provided: %s"
			return vsphereschema.Config{}, fmt.Errorf(message, string(char))
		}
	}
	if len(vcenter.Datastore) > 80 {
		message := "💥 The vSphere Datastore name must be less than 80 characters"
		return vsphereschema.Config{}, fmt.Errorf(message)
	}

	// #########################
	// Infrastructure
	// #########################

	// Verify each of the Virtual Machine configurations.

	// Verify each of the Resource Pool configurations.

	// #########################
	// Configuration
	// #########################

	// Build the final configuration to be returned.
	configFinal := vsphereschema.Config{
		Enabled:        cfg.Enabled.(bool),
		Vcenter:        vcenter,
		Infrastructure: infrastructure,
	}

	return configFinal, nil
}
*/
