// Package awsconfig contains the functions for
// verifying the AWS configuration.
package awsconfig

import (
	"errors"
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	utils "homelab/pkg/utils"

	awsschema "homelab/pkg/cloud/aws/schema"
)

// VerifyConfig verifies the AWS configuration.
func VerifyConfig(
	ctx *pulumi.Context,
	configRaw awsschema.ConfigRaw,
) (awsschema.Config, error) {

	var configFinal awsschema.Config

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

	// Take the raw configuration and unmarshal it into the schema for final validation.
	if err := utils.TryObject("", configRaw, &configFinal); err != nil {
		return configFinal, errors.New("failed to parse configuration: " + err.Error())
	}

	// Validate the configuration using the validator tags.
	if err := configFinal.ValidateAWS(); err != nil {
		return awsschema.Config{}, err
	}

	utils.LogInfo(ctx, "✅ AWS configuration verified successfully")

	return configFinal, nil
}
