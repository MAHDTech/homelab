// Package azureconfig contains the functions for
// verifying the Azure configuration.
package azureconfig

import (
	"errors"
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	utils "homelab/pkg/utils"

	azureschema "homelab/pkg/cloud/azure/schema"
)

// VerifyConfig verifies the Azure configuration.
func VerifyConfig(
	ctx *pulumi.Context,
	configRaw azureschema.ConfigRaw,
) (azureschema.Config, error) {

	var configFinal azureschema.Config

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

	// Take the raw configuration and unmarshal it into the schema for final validation.
	if err := utils.TryObject("", configRaw, &configFinal); err != nil {
		return configFinal, errors.New("failed to parse configuration: " + err.Error())
	}

	// Validate the configuration using the validator tags.
	if err := configFinal.ValidateAzure(); err != nil {
		return azureschema.Config{}, err
	}

	utils.LogInfo(ctx, "✅ Azure configuration verified successfully")

	return configFinal, nil
}
