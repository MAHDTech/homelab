package azureconfig

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func VerifyConfig(ctx *pulumi.Context, cfg ConfigRaw) (Config, error) {

	// TODO: Verify the configuration. For now, return the raw configuration.
	configFinal := Config{
		Enabled: cfg.Enabled.(bool),
	}

	return configFinal, nil
}
