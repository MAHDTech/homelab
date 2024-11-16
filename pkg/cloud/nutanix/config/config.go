// Package nutanixconfig contains the functions for verifying the Nutanix configuration.
package nutanixconfig

import (
	"errors"
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	nutanixschema "homelab/pkg/cloud/nutanix/schema"

	utils "homelab/pkg/utils"
)

// VerifyConfig verifies the Nutanix configuration.
func VerifyConfig(
	ctx *pulumi.Context,
	configRaw nutanixschema.ConfigRaw,
) (nutanixschema.Config, error) {

	// #########################
	// Enabled
	// #########################

	// Validate that 'Enabled' is a boolean.
	_, ok := configRaw.Enabled.(bool)
	if !ok {
		message := "💥 'enabled' in the 'nutanix' section must be a boolean. You have provided: " + fmt.Sprint(
			configRaw.Enabled,
		)
		return nutanixschema.Config{}, errors.New(message)
	}

	// If 'Enabled' is false, skip the rest of the configuration verification.
	if !configRaw.Enabled.(bool) {
		utils.LogInfo(ctx, "🛈 Nutanix is disabled, skipping configuration verification.")
		return nutanixschema.Config{Enabled: false}, nil
	}

	/*
		##########################
		TODO: Finish Nutanix config verification here.
		##########################
	*/

	// #########################
	// Configuration
	// #########################

	// Build the final configuration to be returned.

	configFinal := nutanixschema.Config{
		Enabled: configRaw.Enabled.(bool),
	}

	utils.LogInfo(ctx, "✅ Nutanix configuration verified")

	return configFinal, nil
}
