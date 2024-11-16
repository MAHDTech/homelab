// Package vsphereconfig contains the functions for verifying the VMware configuration.
package vsphereconfig

import (
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	vsphereschema "homelab/pkg/cloud/vsphere/schema"

	utils "homelab/pkg/utils"
)

// verifySupervisorConfig ensures the Supervisor configuration is valid.
// If the validation fails, an error is returned.
func verifySupervisorConfig(
	ctx *pulumi.Context,
	configRaw interface{},
) (vsphereschema.Supervisor, error) {

	var configFinal vsphereschema.Supervisor

	// Take the raw configuration and unmarshal it into the schema.
	err := utils.TryObject(configRaw, "supervisor", &configFinal)
	if err != nil {
		return configFinal, err
	}

	// nolint:errcheck
	ctx.Log.Info("TODO: Validate the configuration for 'Supervisor' here.", nil)

	return configFinal, nil
}
