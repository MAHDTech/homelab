// Package globalconfig is used to verify the raw global
// configuration section of the stack configuration.
package globalconfig

import (
	"errors"
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	globalschema "homelab/pkg/cloud/global/schema"

	utils "homelab/pkg/utils"
)

// VerifyConfig takes a raw global configuration and
// validates the required items.
// If a required item is missing, an err is returned.
func VerifyConfig(
	ctx *pulumi.Context,
	configGlobal globalschema.ConfigRaw,
) (globalschema.Config, error) {

	// Validate that 'Enabled' is a boolean.
	_, ok := configGlobal.Enabled.(bool)
	if !ok {
		message := "💥 'enabled' in the 'global' section must be a boolean. You have provided: " + fmt.Sprint(
			configGlobal.Enabled,
		)
		return globalschema.Config{}, errors.New(message)
	}

	// Validate that 'Debug' is a boolean.
	_, ok = configGlobal.Debug.(bool)
	if !ok {
		message := "💥 'debug' in the 'global' section must be a boolean. You have provided: " + fmt.Sprint(
			configGlobal.Debug,
		)
		return globalschema.Config{}, errors.New(message)
	}

	// Create the final struct to be returned.
	configGlobalFinal := globalschema.Config{
		Enabled: configGlobal.Enabled.(bool),
		Debug:   configGlobal.Debug.(bool),
	}

	utils.LogInfo(ctx, "✅ Global configuration verified")

	return configGlobalFinal, nil
}
