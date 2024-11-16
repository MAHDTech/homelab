// Package vsphereconfig contains the functions for verifying the VMware configuration.
package vsphereconfig

import (
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	vsphereschema "homelab/pkg/cloud/vsphere/schema"

	utils "homelab/pkg/utils"
)

// verifyNetworkConfig ensures the Network configuration is valid.
// If the validation fails, an error is returned.
func verifyNetworkConfig(
	ctx *pulumi.Context,
	configRaw interface{},
) (vsphereschema.Network, error) {

	var configFinal vsphereschema.Network

	// Take the raw configuration and unmarshal it into the schema.
	err := utils.TryObject(configRaw, "network", &configFinal)
	if err != nil {
		return configFinal, err
	}

	// nolint:errcheck
	ctx.Log.Info("TODO: Validate the configuration for 'Network' here.", nil)

	return configFinal, nil
}
