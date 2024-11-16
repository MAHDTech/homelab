// Package vsphereconfig contains the functions for verifying the VMware configuration.
package vsphereconfig

import (
	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	vsphereschema "homelab/pkg/cloud/vsphere/schema"

	utils "homelab/pkg/utils"
)

// verifyCommonConfig ensures the Common configuration is valid.
// If the validation fails, an error is returned.
func verifyCommonConfig(ctx *pulumi.Context, configRaw interface{}) (vsphereschema.Common, error) {

	var configFinal vsphereschema.Common

	// Take the raw configuration and unmarshal it into the schema.
	err := utils.TryObject(configRaw, "common", &configFinal)
	if err != nil {
		return configFinal, err
	}

	// nolint:errcheck
	ctx.Log.Info("TODO: Validate the configuration for 'Common' here.", nil)

	return configFinal, nil
}
