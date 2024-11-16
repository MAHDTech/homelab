// Package config is used to verify the raw configuration file.
// Each cloud provider has its own configuration schema and
// validation logic which this package coordinates.
package config

import (
	"errors"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiConfig "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"

	schema "homelab/pkg/schema"
	utils "homelab/pkg/utils"

	globalconfig "homelab/pkg/cloud/global/config"

	awsconfig "homelab/pkg/cloud/aws/config"
	azureconfig "homelab/pkg/cloud/azure/config"
	gcpconfig "homelab/pkg/cloud/gcp/config"
	nutanixconfig "homelab/pkg/cloud/nutanix/config"
	vsphereconfig "homelab/pkg/cloud/vsphere/config"
)

// VerifyConfig takes a raw configuration and
// checks for required configuration items
// in the Global and Cloud Provider sections.
// If a required item is missing, an err is returned.
// Unlike the built-in pulumi "Require" this doesn't panic
// and display a stack trace, rather just the specific err
// message which is more user friendly and can be customised
// for the specific cloud provider.
func VerifyConfig(ctx *pulumi.Context, configPulumi pulumiConfig.Config) (schema.Config, error) {

	var configRaw schema.ConfigRaw
	var configFinal schema.Config
	var err error

	utils.LogInfo(ctx, "✨ Validating 'global' configuration...")

	// Read the 'Global' section of the configuration.
	err = configPulumi.TryObject("global", &configRaw.Global)
	if err != nil {
		message := "💥 The 'global' section in the configuration file is invalid: " + err.Error()
		return configFinal, errors.New(message)
	}

	// Validate the 'Global' section of the configuration.
	configGlobal, err := globalconfig.VerifyConfig(ctx, configRaw.Global)
	if err != nil {
		message := "💥 The 'global' section in the configuration file is invalid: " + err.Error()
		return configFinal, errors.New(message)
	}

	utils.LogInfo(ctx, "✨ Validating 'aws' configuration...")

	// Read the 'AWS' section of the configuration.
	err = configPulumi.TryObject("aws", &configRaw.AWS)
	if err != nil {
		message := "💥 The 'aws' section in the configuration file is invalid: " + err.Error()
		return configFinal, errors.New(message)
	}

	// Validate the 'AWS' section of the configuration.
	configAWS, err := awsconfig.VerifyConfig(ctx, configRaw.AWS)
	if err != nil {
		message := "💥 The 'aws' section in the configuration file is invalid: " + err.Error()
		return configFinal, errors.New(message)
	}

	utils.LogInfo(ctx, "✨ Validating 'azure' configuration...")

	// Read the 'Azure' section of the configuration.
	err = configPulumi.TryObject("azure", &configRaw.Azure)
	if err != nil {
		message := "💥 The 'azure' section in the configuration file is invalid: " + err.Error()
		return configFinal, errors.New(message)
	}

	// Validate the 'Azure' section of the configuration.
	configAzure, err := azureconfig.VerifyConfig(ctx, configRaw.Azure)
	if err != nil {
		message := "💥 The 'azure' section in the configuration file is invalid: " + err.Error()
		return configFinal, errors.New(message)
	}

	utils.LogInfo(ctx, "✨ Validating 'gcp' configuration...")

	// Read the 'GCP' section of the configuration.
	err = configPulumi.TryObject("gcp", &configRaw.GCP)
	if err != nil {
		message := "💥 The 'gcp' section in the configuration file is invalid: " + err.Error()
		return configFinal, errors.New(message)
	}

	// Validate the 'GCP' configuration.
	configGCP, err := gcpconfig.VerifyConfig(ctx, configRaw.GCP)
	if err != nil {
		message := "💥 The 'gcp' section in the configuration file is invalid: " + err.Error()
		return configFinal, errors.New(message)
	}

	utils.LogInfo(ctx, "✨ Validating 'nutanix' configuration...")

	// Read the 'Nutanix' section of the configuration.
	err = configPulumi.TryObject("nutanix", &configRaw.Nutanix)
	if err != nil {
		message := "💥 The 'nutanix' section in the configuration file is invalid: " + err.Error()
		return configFinal, errors.New(message)
	}

	// Validate the 'Nutanix' section of the configuration.
	configNutanix, err := nutanixconfig.VerifyConfig(ctx, configRaw.Nutanix)
	if err != nil {
		message := "💥 The 'nutanix' section in the configuration file is invalid: " + err.Error()
		return configFinal, errors.New(message)
	}

	utils.LogInfo(ctx, "✨ Validating 'vsphere' configuration...")

	// Read the 'vSphere' section of the configuration.
	err = configPulumi.TryObject("vsphere", &configRaw.VSphere)
	if err != nil {
		message := "💥 The 'vsphere' section in the configuration file is invalid: " + err.Error()
		return configFinal, errors.New(message)
	}

	// Validate the 'vSphere' section of the configuration.
	configVSphere, err := vsphereconfig.VerifyConfig(ctx, configRaw.VSphere)
	if err != nil {
		message := "💥 The 'vsphere' section in the configuration file is invalid: " + err.Error()
		return configFinal, errors.New(message)
	}

	// Combine the validated configurations into a single struct.
	configFinal = schema.Config{
		Global:  configGlobal,
		AWS:     configAWS,
		Azure:   configAzure,
		GCP:     configGCP,
		Nutanix: configNutanix,
		VSphere: configVSphere,
	}

	utils.LogInfo(ctx, "✅ All configuration verified")

	return configFinal, nil
}
