// Package nutanixconfig contains the functions for
// verifying the Nutanix configuration.
package nutanixconfig

import (
	"errors"
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	utils "homelab/pkg/utils"

	nutanixschema "homelab/pkg/cloud/nutanix/schema"
)

// VerifyConfig verifies the Nutanix configuration.
func VerifyConfig(
	ctx *pulumi.Context,
	configRaw nutanixschema.ConfigRaw,
) (nutanixschema.Config, error) {

	var configFinal nutanixschema.Config

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

	// Take the raw configuration and unmarshal it into the schema for final validation.
	if err := utils.TryObject("", configRaw, &configFinal); err != nil {
		return configFinal, errors.New("failed to parse configuration: " + err.Error())
	}

	// Validate the configuration using the validator tags.
	if err := configFinal.ValidateNutanix(); err != nil {
		return nutanixschema.Config{}, err
	}

	utils.LogInfo(ctx, "✅ Nutanix configuration verified successfully")

	return configFinal, nil
}
