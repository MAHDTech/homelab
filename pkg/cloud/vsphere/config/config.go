// Package vsphereconfig contains the functions for
// working with the VMware vSphere specific configuration.
package vsphereconfig

import (
	"errors"
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	utils "homelab/pkg/utils"

	vsphereschema "homelab/pkg/cloud/vsphere/schema"
)

// VerifyConfig validates the vSphere configuration by
// taking a raw configuration and returning a fully
// validated configuration or an error.
// We use a raw config to handle the unmarshalling
// into an interface before performing validation.
func VerifyConfig(
	ctx *pulumi.Context,
	configRaw vsphereschema.ConfigRaw,
) (vsphereschema.Config, error) {

	var configFinal vsphereschema.Config

	// Validate that 'Enabled' is a boolean.
	_, ok := configRaw.Enabled.(bool)
	if !ok {
		message := "💥 'enabled' in the 'vsphere' section must be a boolean. You have provided: " + fmt.Sprint(
			configRaw.Enabled,
		)
		return vsphereschema.Config{}, errors.New(message)
	}

	// If 'Enabled' is false, skip the rest of the configuration verification.
	if !configRaw.Enabled.(bool) {
		utils.LogInfo(ctx, "🛈 vSphere is disabled, skipping configuration verification.")
		return vsphereschema.Config{Enabled: false}, nil
	}

	// Take the raw configuration and unmarshal it into the schema for final validation.
	if err := utils.TryObject("", configRaw, &configFinal); err != nil {
		return configFinal, errors.New("failed to parse configuration: " + err.Error())
	}

	// Validate the configuration using the validator tags.
	if err := configFinal.ValidateVSphere(ctx); err != nil {
		return vsphereschema.Config{}, err
	}

	utils.LogInfo(ctx, "✅ vSphere configuration verified successfully")
	return configFinal, nil
}
