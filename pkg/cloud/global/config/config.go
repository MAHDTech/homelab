package globalconfig

import (
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func VerifyConfig(ctx *pulumi.Context, configGlobal ConfigRaw) (Config, error) {

	// TODO: Remove debug message.
	message := fmt.Sprintf("DEBUG: configGlobal START: %+v", configGlobal)
	ctx.Log.Info(message, nil)

	// Validate that 'Enabled' is a boolean.
	_, ok := configGlobal.Enabled.(bool)
	if !ok {
		message := "ERROR: Global configuration 'enabled' must be a boolean. You have provided: " + fmt.Sprintf("%v", configGlobal.Enabled)
		return Config{}, fmt.Errorf(message)
	}

	// Validate that 'Debug' is a boolean.
	_, ok = configGlobal.Debug.(bool)
	if !ok {
		message := "ERROR: Global configuration 'debug' must be a boolean. You have provided: " + fmt.Sprintf("%v", configGlobal.Debug)
		return Config{}, fmt.Errorf(message)
	}

	// Create the final struct to be returned.
	configGlobalFinal := Config{
		Enabled: configGlobal.Enabled.(bool),
		Debug:   configGlobal.Debug.(bool),
	}

	// Ensure that the global configuration is enabled.
	if !configGlobalFinal.Enabled {
		message := "NOTICE: Global configuration is disabled, skipping further processing"
		return Config{}, fmt.Errorf(message)
	}

	return configGlobalFinal, nil
}
