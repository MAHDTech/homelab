// Package main is the entry point for the homelab script.
package main

import (
	godotenv "github.com/joho/godotenv"

	pulumi "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiConfig "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"

	cloud "homelab/pkg/cloud"
	config "homelab/pkg/config"
	utils "homelab/pkg/utils"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Load the optional .env file or show a warning if it fails.
		err := godotenv.Load()
		if err != nil {
			utils.LogWarn(
				ctx,
				"Failed to load dotenv file '.env'. Please check that the file exists and is readable: %s",
				err.Error(),
			)
		}

		// Read the pulumi stack variables into a config object.
		configRaw := pulumiConfig.New(ctx, "homelab")

		// Verify the configuration.
		configVerified, err := config.VerifyConfig(ctx, *configRaw)
		if err != nil {
			utils.LogError(
				ctx,
				"❗ A fatal error occurred during verification of the configuration file: %s",
				err.Error(),
			)
			return err
		}

		// Create the cloud resources.
		err = cloud.CreateResources(ctx, &configVerified)
		if err != nil {
			utils.LogError(
				ctx,
				"❗ A fatal error occurred during creation of the cloud resources: %s",
				err.Error(),
			)
			return err
		}

		return nil

	})

}
