// Package gcpconfig contains the functions for verifying the GCP configuration.
package gcpconfig

import (
	"errors"
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	gcpschema "homelab/pkg/cloud/gcp/schema"

	utils "homelab/pkg/utils"
)

// VerifyConfig verifies the GCP configuration.
func VerifyConfig(ctx *pulumi.Context, configRaw gcpschema.ConfigRaw) (gcpschema.Config, error) {

	// #########################
	// Enabled
	// #########################

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

	/*
		##########################
		TODO: Finish GCP config verification here.
		##########################
	*/

	// #########################
	// Configuration
	// #########################

	// Build the final configuration to be returned.

	configFinal := gcpschema.Config{
		Enabled: configRaw.Enabled.(bool),
	}

	utils.LogInfo(ctx, "✅ GCP configuration verified")
	return configFinal, nil
}
