// Package awsconfig contains the functions for verifying the AWS configuration.
package awsconfig

import (
	"errors"
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	awsschema "homelab/pkg/cloud/aws/schema"
	utils "homelab/pkg/utils"
)

// VerifyConfig verifies the AWS configuration.
func VerifyConfig(ctx *pulumi.Context, configRaw awsschema.ConfigRaw) (awsschema.Config, error) {

	// #########################
	// Enabled
	// #########################

	// Validate that 'Enabled' is a boolean.
	_, ok := configRaw.Enabled.(bool)
	if !ok {
		message := "💥 'enabled' in the 'aws' section must be a boolean. You have provided: " + fmt.Sprint(
			configRaw.Enabled,
		)
		return awsschema.Config{}, errors.New(message)
	}

	// If 'Enabled' is false, skip the rest of the configuration verification.
	if !configRaw.Enabled.(bool) {
		utils.LogInfo(ctx, "🛈 AWS is disabled, skipping configuration verification.")
		return awsschema.Config{Enabled: false}, nil
	}

	/*
		##########################
		TODO: Finish AWS config verification here.
		##########################
	*/

	// #########################
	// Configuration
	// #########################

	// Build the final configuration to be returned.

	configFinal := awsschema.Config{
		Enabled: configRaw.Enabled.(bool),
	}

	utils.LogInfo(ctx, "✅ AWS configuration verified")

	return configFinal, nil
}
