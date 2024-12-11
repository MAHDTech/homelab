// Package gcpconfig contains the functions for
// verifying the GCP configuration.
package gcpconfig

import (
	"errors"
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	utils "homelab/pkg/utils"

	gcpschema "homelab/pkg/cloud/gcp/schema"
)

// VerifyConfig verifies the GCP configuration.
func VerifyConfig(
	ctx *pulumi.Context,
	configRaw gcpschema.ConfigRaw,
) (gcpschema.Config, error) {

	var configFinal gcpschema.Config

	// Validate that 'Enabled' is a boolean.
	_, ok := configRaw.Enabled.(bool)
	if !ok {
		message := "💥 'enabled' in the 'gcp' section must be a boolean. You have provided: " + fmt.Sprint(
			configRaw.Enabled,
		)
		return gcpschema.Config{}, errors.New(message)
	}

	// If 'Enabled' is false, skip the rest of the configuration verification.
	if !configRaw.Enabled.(bool) {
		utils.LogInfo(ctx, "🛈 GCP is disabled, skipping configuration verification.")
		return gcpschema.Config{Enabled: false}, nil
	}

	// Take the raw configuration and unmarshal it into the schema for final validation.
	if err := utils.TryObject("", configRaw, &configFinal); err != nil {
		return configFinal, errors.New("failed to parse configuration: " + err.Error())
	}

	// Validate the configuration using the validator tags.
	if err := configFinal.ValidateGCP(); err != nil {
		return gcpschema.Config{}, err
	}

	utils.LogInfo(ctx, "✅ GCP configuration verified successfully")

	return configFinal, nil
}
