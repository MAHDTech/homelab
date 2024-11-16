// Package azureconfig contains the functions for verifying the Azure configuration.
package azureconfig

import (
	"errors"
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	azureschema "homelab/pkg/cloud/azure/schema"

	utils "homelab/pkg/utils"
)

// VerifyConfig verifies the Azure configuration.
func VerifyConfig(
	ctx *pulumi.Context,
	configRaw azureschema.ConfigRaw,
) (azureschema.Config, error) {

	// #########################
	// Enabled
	// #########################

	// Validate that 'Enabled' is a boolean.
	_, ok := configRaw.Enabled.(bool)
	if !ok {
		message := "💥 'enabled' in the 'azure' section must be a boolean. You have provided: " + fmt.Sprint(
			configRaw.Enabled,
		)
		return azureschema.Config{}, errors.New(message)
	}

	// If 'Enabled' is false, skip the rest of the configuration verification.
	if !configRaw.Enabled.(bool) {
		utils.LogInfo(ctx, "🛈 Azure is disabled, skipping configuration verification.")
		return azureschema.Config{Enabled: false}, nil
	}

	/*
		##########################
		TODO: Finish Azure config verification here.
		##########################
	*/

	// #########################
	// Configuration
	// #########################

	// Build the final configuration to be returned.

	configFinal := azureschema.Config{
		Enabled: configRaw.Enabled.(bool),
	}

	utils.LogInfo(ctx, "✅ Azure configuration verified")
	return configFinal, nil
}
