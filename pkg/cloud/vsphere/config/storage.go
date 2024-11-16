// Package vsphereconfig contains the functions for verifying the VMware configuration.
package vsphereconfig

import (
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	vsphereschema "homelab/pkg/cloud/vsphere/schema"

	utils "homelab/pkg/utils"
)

// verifyStorageConfig ensures the Storage configuration is valid.
// If the validation fails, an error is returned.
func verifyStorageConfig(
	ctx *pulumi.Context,
	configRaw interface{},
) (vsphereschema.Storage, error) {

	var configFinal vsphereschema.Storage

	// Take the raw configuration and unmarshal it into the schema.
	err := utils.TryObject(configRaw, "storage", &configFinal)
	if err != nil {
		return configFinal, err
	}

	// nolint:errcheck
	ctx.Log.Info("TODO: Validate the configuration for 'Storage' here.", nil)

	return configFinal, nil
}
