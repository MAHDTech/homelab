package config

import (
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiConfig "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"

	awsconfig "homelab/pkg/cloud/aws/config"
	azureconfig "homelab/pkg/cloud/azure/config"
	gcpconfig "homelab/pkg/cloud/gcp/config"
	globalconfig "homelab/pkg/cloud/global/config"
	nutanixconfig "homelab/pkg/cloud/nutanix/config"
	vsphereconfig "homelab/pkg/cloud/vsphere/config"
)

// VerifyConfig checks for required configuration items
// in the Global and Cloud Provider sections.
// If a required item is missing, an err is returned.
// Unlike the built-in pulumi "Require" this doesn't panic
// and display a stack trace, rather just the specific err
// message which is more user friendly.
func VerifyConfig(ctx *pulumi.Context, configPulumi pulumiConfig.Config) (Config, error) {

	var configRaw ConfigRaw
	var configFinal Config
	var err error

	ctx.Log.Info("TASK: Reading configuration...", nil)

	// Try and get the 'Global' section of the config.
	err = configPulumi.TryObject("global", &configRaw.Global)
	if err != nil {
		return configFinal, err
	}

	// Try and get the 'AWS' section of the config.
	err = configPulumi.TryObject("aws", &configRaw.AWS)
	if err != nil {
		return configFinal, err
	}

	// Try and get the 'Azure' section of the config.
	err = configPulumi.TryObject("azure", &configRaw.Azure)
	if err != nil {
		return configFinal, err
	}

	// Try and get the 'GCP' section of the config.
	err = configPulumi.TryObject("gcp", &configRaw.GCP)
	if err != nil {
		return configFinal, err
	}

	// Try and get the 'Nutanix' section of the config.
	err = configPulumi.TryObject("nutanix", &configRaw.Nutanix)
	if err != nil {
		return configFinal, err
	}

	// Try and get the 'VSphere' section of the config.
	err = configPulumi.TryObject("vsphere", &configRaw.VSphere)
	if err != nil {
		return configFinal, err
	}

	ctx.Log.Info("TASK: Verifying configuration...", nil)

	// Validate the 'Global' configuration.
	configGlobal, err := globalconfig.VerifyConfig(ctx, configRaw.Global)
	if err != nil {
		message := "ERROR: Global configuration is invalid: " + err.Error()
		ctx.Log.Error(message, nil)
		return configFinal, fmt.Errorf(message)
	}

	// Validate the 'AWS' configuration.
	configAWS, err := awsconfig.VerifyConfig(ctx, configRaw.AWS)
	if err != nil {
		message := "ERROR: AWS configuration is invalid: " + err.Error()
		ctx.Log.Error(message, nil)
		return configFinal, fmt.Errorf(message)
	}

	// Validate the 'Azure' configuration.
	configAzure, err := azureconfig.VerifyConfig(ctx, configRaw.Azure)
	if err != nil {
		message := "ERROR: Azure configuration is invalid: " + err.Error()
		ctx.Log.Error(message, nil)
		return configFinal, fmt.Errorf(message)
	}

	// Validate the 'GCP' configuration.
	configGCP, err := gcpconfig.VerifyConfig(ctx, configRaw.GCP)
	if err != nil {
		message := "ERROR: GCP configuration is invalid: " + err.Error()
		ctx.Log.Error(message, nil)
		return configFinal, fmt.Errorf(message)
	}

	// Validate the 'Nutanix' configuration.
	configNutanix, err := nutanixconfig.VerifyConfig(ctx, configRaw.Nutanix)
	if err != nil {
		message := "ERROR: Nutanix configuration is invalid: " + err.Error()
		ctx.Log.Error(message, nil)
		return configFinal, fmt.Errorf(message)
	}

	// Validate the 'VSphere' configuration.
	configVSphere, err := vsphereconfig.VerifyConfig(ctx, configRaw.VSphere)
	if err != nil {
		message := "ERROR: VSphere configuration is invalid: " + err.Error()
		ctx.Log.Error(message, nil)
		return configFinal, fmt.Errorf(message)
	}

	// Combine the validated configurations.
	configFinal = Config{
		Global:  configGlobal,
		AWS:     configAWS,
		Azure:   configAzure,
		GCP:     configGCP,
		Nutanix: configNutanix,
		VSphere: configVSphere,
	}

	return configFinal, nil
}
