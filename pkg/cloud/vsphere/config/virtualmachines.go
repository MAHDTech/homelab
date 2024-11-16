// Package vsphereconfig contains the functions for verifying the VMware configuration.
package vsphereconfig

import (
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	vsphereschema "homelab/pkg/cloud/vsphere/schema"

	utils "homelab/pkg/utils"
)

// verifyVirtualMachinesConfig ensures the Virtual Machines configuration is valid.
// If the validation fails, an error is returned.
func verifyVirtualMachinesConfig(
	ctx *pulumi.Context,
	configRaw interface{},
) (vsphereschema.VirtualMachines, error) {

	var configFinal vsphereschema.VirtualMachines

	// Take the raw configuration and unmarshal it into the schema.
	err := utils.TryObject(configRaw, "virtual_machines", &configFinal)
	if err != nil {
		return configFinal, err
	}

	// nolint:errcheck
	ctx.Log.Info("TODO: Validate the configuration for 'VirtualMachines' here.", nil)

	return configFinal, nil
}
