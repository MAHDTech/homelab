package main

import (
	"fmt"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiConfig "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"

	homelabConfig "homelab/pkg/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Read the pulumi stack variables.
		stackConfig := pulumiConfig.New(ctx, "homelab")

		// Validate the common configuration.
		configCommon, err := homelabConfig.Validate(ctx, stackConfig)
		if err != nil {
			message := "Failed to verify common configuration!"
			ctx.Log.Error(message, nil)
			return err
		}

		// For each supported cloud provider, validate the configuration.
		// If the cloud provider is not enabled, the validation will be skipped.
		for _, cloud := range homelabConfig.CloudProviders {
			message := fmt.Sprintf("Validating %s cloud provider configuration...", cloud)
			ctx.Log.Info(message, nil)
			_, err := homelabConfig.ValidateCloud(ctx, stackConfig, cloud)

			if err != nil {
				message := fmt.Sprintf("Failed to verify %s cloud provider configuration!", cloud)
				ctx.Log.Error(message, nil)
				return err
			}
		}

		return nil

	})

}
